import os, json, datetime
import falcon.asgi

from db import get_pool 
from models import (
    sql_upsert_saldo,
    sql_insert_tx,
    sql_select_saldo,
    sql_select_ultimas,
    sql_upsert_e_insert
)

class TransacaoResource:
    async def on_post(self, req, resp, id: int):
        body = await req.get_media()
        try:
            valor = int(body['valor'])
            if valor <= 0:
                raise ValueError
            tipo = body['tipo']
            if tipo not in ('c', 'd'):
                raise ValueError
            descricao = body['descricao']
            if not (1 <= len(descricao) <= 10):
                raise ValueError
        except (KeyError, ValueError):
            raise falcon.HTTPUnprocessableEntity()

        delta = valor if tipo == 'c' else -valor
        
        async with (await get_pool()).acquire() as conn:
            row = await conn.fetchrow(
                sql_upsert_e_insert,
                id,      
                delta,   
                valor,   
                tipo,    
                descricao
            )

        if row is None:
            raise falcon.HTTPUnprocessableEntity()

        resp.media  = {'limite': row['limite'], 'saldo': row['saldo']}
        resp.status = falcon.HTTP_200

class ExtratoResource:
    async def on_get(self, req, resp, id: int):
        async with (await get_pool()).acquire() as conn:
            saldo = await conn.fetchrow(sql_select_saldo, id)
            if saldo is None:
                raise falcon.HTTPNotFound()
            txs = await conn.fetch(sql_select_ultimas, id)

        resp.media = {
            'saldo': {
                'total': saldo['saldo'],
                'data_extrato': datetime.datetime.utcnow().isoformat(timespec='microseconds') + 'Z',
                'limite': saldo['limite'],
            },
            'ultimas_transacoes': [
                {
                    'valor': r['valor'],
                    'tipo': r['tipo'],
                    'descricao': r['descricao'],
                    'realizada_em': r['realizada_em'].isoformat(timespec='microseconds') + 'Z'
                } for r in txs
            ]
        }

app = falcon.asgi.App()

app.add_route('/clientes/{id:int}/transacoes', TransacaoResource())
app.add_route('/clientes/{id:int}/extrato',     ExtratoResource())
