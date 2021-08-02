from typing import List
from pydantic import BaseModel

class Comodity(BaseModel):
    uuid: str
    komoditas: str
    area_provinsi: str
    area_kota: str
    size: int
    price: int
    tgl_parsed: str
    timestamp: str
    usd_price: int
    idr_usd_rate: float

class Aggregate(BaseModel):
    start_week: str
    area_provinsi: str
    min_size: int
    max_size: int
    med_size: float
    mean_size: float
    min_price: int
    max_size: int
    med_size: float
    mean_size: float

class RespComodity(BaseModel):
    result: List[Comodity]

class RespAggregate(BaseModel):
    result: List[Aggregate]

class RespJWT(BaseModel):
    phone: str
    name: str
    role: str
    timestamp: str