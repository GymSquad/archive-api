import time
from typing import cast

import structlog
from asgi_correlation_id.context import correlation_id
from fastapi import Request, Response, status
from fastapi.datastructures import Address
from starlette.middleware.base import RequestResponseEndpoint
from uvicorn.protocols.utils import get_path_with_query_string

access_logger = structlog.stdlib.get_logger("api.access")


async def logging_middleware(
    request: Request, call_next: RequestResponseEndpoint
) -> Response:
    structlog.contextvars.clear_contextvars()

    # These context vars will be added to all log entries emitted during the request
    request_id = correlation_id.get()
    structlog.contextvars.bind_contextvars(request_id=request_id)

    start_time = time.perf_counter_ns()
    # If the call_next raises an error, we still want to return our own 500 response,
    # so we can add headers to it (process time, request ID...)
    response = Response(status_code=status.HTTP_500_INTERNAL_SERVER_ERROR)
    try:
        response = await call_next(request)
    except Exception:
        structlog.stdlib.get_logger("api.error").exception("Uncaught exception")
        raise
    finally:
        process_time = time.perf_counter_ns() - start_time
        status_code = response.status_code
        url = get_path_with_query_string(request.scope)  # type: ignore [reportGeneralTypeIssues]
        client_host = cast(Address, request.client).host
        client_port = cast(Address, request.client).port
        http_method = request.method
        http_version = request.scope["http_version"]
        # Recreate the Uvicorn access log format, but add all parameters as structured information
        access_logger.info(
            f"""{client_host}:{client_port} - "{http_method} {url} HTTP/{http_version}" {status_code}""",
            http={
                "url": str(request.url),
                "status_code": status_code,
                "method": http_method,
                "request_id": request_id,
                "version": http_version,
            },
            network={"client": {"ip": client_host, "port": client_port}},
            duration=process_time,
        )
        response.headers["X-Process-Time"] = str(process_time / 1e9)
        return response  # noqa: B012
