from typing import List
from pydantic import BaseModel

class Comodity(BaseModel):
    uuid: str
    komoditas: str
    area_provinsi: str
    area_kota: str
    size: int
    price: float
    tgl_parsed: str
    timestamp: str
    usd_price: float
    usd_idr_rate: float

class Aggregate(BaseModel):
    start_week: str
    area_provinsi: str
    min_size: int
    max_size: int
    med_size: float
    mean_size: float
    min_price: float
    max_price: float
    med_price: float
    mean_price: float

class RespComodity(BaseModel):
    result: List[Comodity]

class RespAggregate(BaseModel):
    result: List[Aggregate]

class RespJWT(BaseModel):
    phone: str
    name: str
    role: str
    timestamp: str