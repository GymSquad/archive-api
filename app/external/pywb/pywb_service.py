from typing import final

import httpx

from app.core.config import settings
from app.external.pywb.types import PywbArchiveInfo
from app.external.utils.http import parse_jsonl_response


@final
class PywbService:
    def __init__(self, base_url: str):
        self.http_client = httpx.AsyncClient(base_url=base_url)

    async def get_archive_info(
        self, collection: str, url: str
    ) -> list[PywbArchiveInfo]:
        response = await self.http_client.get(
            f"/{collection}/cdx",
            params={
                "url": url,
                "output": "json",
                "source": f"{collection}/indexes/autoindex.cdxj",
            },
        )

        return parse_jsonl_response(response, PywbArchiveInfo)


pywb_service: PywbService | None = None


def get_pywb_service() -> PywbService:
    global pywb_service
    if pywb_service is None:
        pywb_service = PywbService(settings.PYWB_BASE_URL)

    return pywb_service
