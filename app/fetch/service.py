import asyncio
from re import L
import time
import pandas as pd
from typing import Dict, Union, List
from aiohttp import ClientSession


class Service():
    def __init__(self, currency_api_key: str, resource_addr: str, query_interval: int) -> None:
        self.currency_addr: str = f'http://data.fixer.io/api/latest?access_key={currency_api_key}'
        self.resource_addr: str = resource_addr
        self.query_interval: int = query_interval
        self.client: Union[ClientSession, None] = None
        self.rate: float = 0
        self.last_request: float = 0
        self.comodity_cache: Union[dict, None] = None
        self.aggregate_cache: Union[dict, None] = None
        self.query_task: Union[asyncio.Task, None] = None

    def filtering(self, data: List[Dict]) -> List[Dict]:
        new_resources = []
        for item in data:
            if None in item.values():
                continue
            item['price'] = int(item['price'])
            item['size'] = int(item['size'])
            item['tgl_parsed'] = item['tgl_parsed'][0:10]
            new_resources.append(item)

        return new_resources

    def preprocessing(self, data: List[Dict]):
        for item in data:
            item.pop('uuid', None)
            item.pop('komoditas', None)
            item.pop('area_kota', None)
            item.pop('timestamp', None)

    async def periodic_rate_query(self):
        async with self.client.get(self.currency_addr) as resp:
            body = await resp.json()
            if body.get('success', False):
                self.rate = body['rates']['IDR']
        await asyncio.sleep(self.query_interval)

    async def fetch(self) -> dict:
        async with self.client.get(self.resource_addr) as resp:
            body = await resp.json()
        post_filter = self.filtering(body)

        for data in post_filter:
            data['usd_price'] = self.rate * data['price']
            data['idr_usd_rate'] = self.rate
        result = {'result': post_filter}
        self.last_request = time.time()
        self.comodity_cache = result
        return result

    async def aggregate(self) -> dict:
        async with self.client.get(self.resource_addr) as resp:
            body = await resp.json()
        post_filter = self.filtering(body)
        self.preprocessing(post_filter)
        df = pd.DataFrame(post_filter)
        df['tgl_parsed'] = pd.to_datetime(df['tgl_parsed'])
        df = df.groupby([pd.Grouper(key='tgl_parsed', freq='W-MON'), 'area_provinsi']).agg(
            min_size=('size', 'min'),
            max_size=('size', 'max'),
            med_size=('size', 'median'),
            mean_size=('size', 'mean'),
            min_price=('price', 'min'),
            max_price=('price', 'max'),
            med_price=('price', 'median'),
            mean_price=('price', 'mean')
        )
        df_dict = df.to_dict(orient='index')
        df_dict_parsed = []
        for k, v in df_dict.items():
            tgl_parsed = k[0].strftime('%Y-%m-%d')
            area = k[1]
            v['start_week'] = tgl_parsed
            v['area_provinsi'] = area
            df_dict_parsed.append(v)
        result = {'result': df_dict_parsed}
        self.aggregate_cache = result
        return result