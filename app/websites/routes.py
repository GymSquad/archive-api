from datetime import date
from typing import Annotated

from fastapi import APIRouter, Query
from pydantic_core import Url

from app.websites.schemas import (
    Pagination,
    SearchResponse,
    UpdateResponse,
    UpdateWebsitePayload,
)

router = APIRouter(
    prefix="/api/website",
    tags=["websites"],
)


@router.get("/{website_id}")
async def get_archived_dates(website_id: str) -> list[date]:
    return []


@router.patch("/{website_id}")
async def update_website(
    website_id: str, payload: UpdateWebsitePayload
) -> UpdateResponse:
    return UpdateResponse(
        id="clrnc14dr000008l235gx4c6c",
        campus="交大相關",
        department="行政單位",
        office="圖書館",
        name="交通大學圖書館",
        url=Url("https://www.lib.nctu.edu.tw/"),
    )


@router.get(
    "/{website_id}/search",
    responses={400: {"description": "Invalid parameters"}},
)
async def search(
    website_id: str,
    q: str = "",
    cursor: str = "",
    limit: Annotated[int, Query(ge=1, le=50)] = 10,
) -> SearchResponse:
    return SearchResponse(
        result=[],
        pagination=Pagination(
            next_cursor="",
            has_next=False,
            num_results=0,
            total_results=0,
        ),
    )
