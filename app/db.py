from typing import Any, AsyncGenerator, Protocol

import asyncpg

from app.core import settings

type Pool = asyncpg.pool.Pool
type Connection = asyncpg.connection.Connection


class Record(Protocol):
    def __getitem__(self, key: str) -> Any:
        ...


pool: Pool | None = None


async def init_pool():
    global pool
    pool = await asyncpg.create_pool(dsn=str(settings.DATABASE_URL))


async def close_pool():
    global pool
    if pool is not None:
        await pool.close()
        pool = None


async def get_db() -> AsyncGenerator[Connection, None]:
    global pool
    if pool is None:
        await init_pool()

    assert pool is not None
    async with pool.acquire() as connection:
        yield connection
