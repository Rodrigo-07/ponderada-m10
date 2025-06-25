import os, asyncpg, asyncio, logging
from typing import Optional

DB_DSN = os.getenv('DB_DSN', 'postgresql://admin:123@db:5432/rinha')

_pool = None

# async def init_pool():
#     global _pool
#     if _pool is None:
        
#         # pool de conexões para não precisar criar uma nova conexão a cada vez
#         _pool = await asyncpg.create_pool(
#             dsn=DB_DSN,
#             min_size=10,
#             max_size=25,
#             statement_cache_size=0, # try to avoid to use more memory
#         )

_lock = asyncio.Lock() # one time init

async def get_pool() -> asyncpg.Pool:
    """Retorna o pool, criando-o se ainda não existir."""
    global _pool
    if _pool is None:
        async with _lock:
            if _pool is None:
                retries = 15
                while retries:
                    try:
                        _pool = await asyncpg.create_pool(
                            dsn=DB_DSN,
                            min_size=5,
                            max_size=15,
                            command_timeout=None
                        )
                        # _pool = await asyncpg.create_pool(
                        #     dsn=DB_DSN,
                        #     min_size=10,
                        #     max_size=25,
                        #     statement_cache_size=0, # try to avoid to use more memory
                        # )
                        break
                    except Exception as e:
                        logging.warning("BD indisponível, retry… (%s)", e)
                        retries -= 1
                        await asyncio.sleep(2)
                if _pool is None:
                    raise RuntimeError("Não conectou ao Postgres")
    return _pool
