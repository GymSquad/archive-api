from datetime import date
from itertools import groupby
from typing import Annotated

import structlog
from fastapi import APIRouter, Depends, HTTPException, Query, status
from pydantic_core import Url

from app import db
from app.websites import search
from app.websites.schemas import (
    Affiliation,
    Pagination,
    SearchResponse,
    SearchResultEntry,
    SearchResultWebsite,
    UpdateResponse,
    UpdateWebsitePayload,
)
from app.websites.search.cursor import SearchWebsiteCursor

router = APIRouter(
    prefix="/api/website",
    tags=["websites"],
)


@router.patch("/{website_id}")
async def update_website(
    website_id: str, payload: UpdateWebsitePayload
) -> UpdateResponse:
    return UpdateResponse(
        id="clrnc14dr000008l235gx4c6c",
        affiliations=[
            Affiliation(
                campus="交大相關",
                department="行政單位",
                office="圖書館",
            )
        ],
        name="交通大學圖書館",
        url=Url("https://www.lib.nctu.edu.tw/"),
    )


@router.get(
    "/search",
    responses={
        400: {
            "description": "Invalid parameters",
            "content": {
                "application/json": {
                    "example": {
                        "detail": [
                            {
                                "loc": ["query", "cursor"],
                                "msg": "invalid key-value pair format",
                                "type": "invalid_cursor_format",
                            },
                        ],
                    },
                },
            },
        },
    },
)
async def search_websites(
    db: Annotated[db.Connection, Depends(db.get_db)],
    q: str | None = None,
    cursor: Annotated[
        str | None, Query(description="Cursor for pagination (in base64)")
    ] = None,
    limit: Annotated[int, Query(ge=1, le=50)] = 10,
) -> SearchResponse:
    logger = structlog.get_logger("websites.search")

    match cursor:
        case None:
            cursor_obj = SearchWebsiteCursor()
        case c if (parsed := SearchWebsiteCursor.from_base64(c)).is_ok():
            cursor_obj = parsed.unwrap()
        case _:
            logger.warning("Invalid cursor received", cursor=cursor)
            raise HTTPException(
                status_code=status.HTTP_400_BAD_REQUEST,
                detail=[
                    {
                        "loc": ["query", "cursor"],
                        "msg": "invalid key-value pair format",
                        "type": "invalid_cursor_format",
                    },
                ],
            )

    searched = await search.search_websites(db, q or "", cursor_obj, limit)

    grouped = groupby(
        searched.entries,
        key=lambda e: f"{e.campus_id}${e.department_id}${e.office_id}",
    )

    result: list[SearchResultEntry] = []
    for group_key, search_entries in grouped:
        search_entries = list(search_entries)
        if len(search_entries) == 0:
            continue

        first = search_entries[0]
        websites = [
            SearchResultWebsite(
                id=entry.website_id,
                name=entry.website_name,
                url=Url(entry.website_url),
            )
            for entry in search_entries
        ]
        result.append(
            SearchResultEntry(
                id=group_key,
                campus=first.campus_name,
                department=first.department_name,
                office=first.office_name,
                websites=websites,
            )
        )

    return SearchResponse(
        result=result,
        pagination=Pagination(
            next_cursor=searched.next_cursor.to_base64()
            if searched.next_cursor is not None
            else None,
            num_results=len(searched.entries),
            num_left=searched.num_left,
        ),
    )


# This needs to be after `search_websites` because otherwise FastAPI will
# treat `search` as a website_id.
@router.get("/{website_id}")
async def get_archived_dates(website_id: str) -> list[date]:
    return []
