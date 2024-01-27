import base64
import binascii
import string
from datetime import date
from typing import Annotated, Literal, Optional, Self

import structlog
from fastapi import APIRouter, HTTPException, Query, status
from pydantic import BaseModel, ConfigDict, ValidationError
from pydantic_core import Url
from result import Err, Ok, Result

from app.websites.schemas import (
    Affiliation,
    Pagination,
    SearchResponse,
    UpdateResponse,
    UpdateWebsitePayload,
)

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


type ParseCursorErrorType = Literal[
    "base64_decode_error",
    "unicode_decode_error",
    "invalid_cursor_format",
    "validation_failure",
]


class ParseCursorError(BaseModel):
    type: ParseCursorErrorType
    msg: str

    model_config: ConfigDict = {"extra": "allow"}


class SearchWebsiteCursor(BaseModel):
    campus_id: str
    department_id: str
    office_id: str
    website_id: str

    def to_base64(self) -> str:
        str_pairs = ",".join(
            f"{key}={value}" for key, value in self.model_dump().items()
        )
        return base64.b64encode(f"({str_pairs})".encode()).decode()

    @classmethod
    def from_base64(cls, base64_str: str) -> Result[Self, ParseCursorError]:
        try:
            decoded = base64.b64decode(base64_str, validate=True)
        except binascii.Error as e:
            return Err(ParseCursorError(type="base64_decode_error", msg=str(e)))

        try:
            decoded = decoded.decode()
        except UnicodeDecodeError as e:
            return Err(ParseCursorError(type="unicode_decode_error", msg=str(e)))

        str_pairs = decoded.strip(string.whitespace + "()").split(",")
        if not all(pair.count("=") == 1 for pair in str_pairs):
            return Err(
                ParseCursorError(
                    type="invalid_cursor_format",
                    msg="invalid key-value pair format",
                )
            )

        if (n := len(str_pairs)) != 4:
            return Err(
                ParseCursorError.model_validate(
                    {
                        "type": "invalid_cursor_format",
                        "msg": "invalid number of pairs",
                        "got": n,
                        "expected": 4,
                    }
                )
            )

        try:
            return Ok(cls.model_validate(dict(pair.split("=") for pair in str_pairs)))
        except ValidationError as e:
            return Err(ParseCursorError(type="validation_failure", msg=str(e)))


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
    q: str | None = None,
    cursor: Annotated[
        str | None, Query(description="Cursor for pagination (in base64)")
    ] = None,
    limit: Annotated[int, Query(ge=1, le=50)] = 10,
) -> SearchResponse:
    logger = structlog.get_logger("websites.search")

    cursor_obj: Optional[SearchWebsiteCursor] = None
    if cursor is not None:
        match SearchWebsiteCursor.from_base64(cursor):
            case Ok(obj):
                cursor_obj = obj
            case Err(err):
                logger.warning(
                    "Invalid cursor received",
                    cursor=cursor,
                    error=err,
                )
                raise HTTPException(
                    status_code=status.HTTP_400_BAD_REQUEST,
                    detail=[{"loc": ["query", "cursor"], **err.model_dump()}],
                )

    logger.info("Cursor object", cursor_obj=cursor_obj)

    return SearchResponse(
        result=[],
        pagination=Pagination(
            next_cursor="",
            num_results=0,
            total_results=0,
        ),
    )


# This needs to be after `search_websites` because otherwise FastAPI will
# treat `search` as a website_id.
@router.get("/{website_id}")
async def get_archived_dates(website_id: str) -> list[date]:
    return []
