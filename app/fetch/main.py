import asyncio
import os
import time
from aiohttp.client import ClientSession
from fastapi import responses

import jwt
from fastapi import FastAPI, Request, Security
from fastapi.openapi.models import APIKey
from fastapi.param_functions import Header
from fastapi.responses import JSONResponse, Response
from fastapi.security import APIKeyHeader
from jwt.exceptions import InvalidSignatureError

import models
from service import Service

SECRET = os.getenv('SECRET', 'secret')
CURRENCY_API_KEY = os.getenv('CURRENCY_API_KEY', '5d2810f27649c3f5fc456d352effea11')
QUERY_INTERVAL = int(os.getenv('QUERY_INTERVAL', '10800'))
RESOURCE_ADDR = os.getenv('RESOURCE_ADDR', 'https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list')
API_KEY_NAME = 'Authorization'
api_key_header = APIKeyHeader(name=API_KEY_NAME)

app = FastAPI(
    title='Fetch Service API',
    description='This is a simple fetch service for querying resources and verifying JWT'
)
service = Service(CURRENCY_API_KEY, RESOURCE_ADDR, QUERY_INTERVAL)
app.state.service = service

@app.on_event('startup')
async def startup():
    service.client = ClientSession()
    service.query_task = asyncio.create_task(service.periodic_rate_query())  # Start background task for querying currency rate periodically

@app.middleware('http')
async def jwt_middleware(req: Request, call_next) -> Response:
    # Exclude API documentation path from JWT verification
    if req.url.path in ('/redoc', '/docs', '/openapi.json'):
        response = await call_next(req)
        return response

    # Verify JWT from Authorization header
    if 'authorization' not in req.headers:
        return JSONResponse(status_code=401, content={'remark': 'Missing Authorization header'})
    apikey = req.headers['authorization']
    try:
        claims = jwt.decode(apikey, SECRET, ['HS256'])
        req.state.claims = claims
    except InvalidSignatureError:  # Invalid signature JWT
        return JSONResponse(status_code=401, content={'remark': 'Invalid authorization token'})
    if req.url.path == '/aggregate' and claims['role'] != 'admin':  # Invalid role for accessing resources
        return JSONResponse(status_code=401, content={'remark': 'Only admin role can access aggregate'})
    response = await call_next(req)
    return response

@app.get('/fetch', tags=['Fetch'], response_model=models.RespComodity, summary='Fetch comodity data from resources')
async def fetch(Authorization: APIKey=Security(api_key_header)):
    # Check if there is valid cache in memory. Fetch to API if not exist
    if time.time() - service.last_request < QUERY_INTERVAL:
        result = service.comodity_cache
    else:
        result = await service.fetch()
    resp = models.RespComodity(**result)
    return resp

@app.get('/aggregate', tags=['Fetch'], response_model=models.RespAggregate, summary='Aggregate comodity data from resources')
async def aggregate(Authorization: APIKey=Security(api_key_header)):
    # Check if there is valid cache in memory. Fetch to API if not exist
    if service.aggregate_cache is not None:
        result = service.aggregate_cache
    else:
        result = await service.aggregate()
    resp = models.RespAggregate(**result)
    return resp

@app.post('/verify', tags=['Fetch'], response_model=models.RespJWT, summary='Parse JWT claims')
async def hello_world(req: Request, Authorization: APIKey=Security(api_key_header)):
    return req.state.claims
