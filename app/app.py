from contextlib import asynccontextmanager

from asgi_correlation_id import CorrelationIdMiddleware
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
from starlette.middleware.base import BaseHTTPMiddleware

from app import db, logger, websites
from app.core import settings

logger.setup_logging(
    development=settings.ENVIRONMENT == "development",
    log_level=settings.LOG_LEVEL,
)


@asynccontextmanager
async def lifespan(_app: FastAPI):
    await db.init_pool()
    yield
    await db.close_pool()


app = FastAPI(
    title="Archive API",
    description="API for the WebArchive project",
    version="0.1.0",
    lifespan=lifespan,
)

origins = ["http://localhost:5173"]
app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

app.add_middleware(BaseHTTPMiddleware, dispatch=logger.logging_middleware)

# This middleware must be placed after the logging, to populate the context with the request ID
# NOTE: Why last??
# Answer: middlewares are applied in the reverse order of when they are added (you can verify this
# by debugging `app.middleware_stack` and recursively drilling down the `app` property).
app.add_middleware(CorrelationIdMiddleware)


class PingResponse(BaseModel):
    message: str = "Server is up and running 🚀"


@app.get("/", tags=["ping"])
async def ping() -> PingResponse:
    """Ping the API to check if it's alive."""

    return PingResponse()


app.include_router(websites.router)
