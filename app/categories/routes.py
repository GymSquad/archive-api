from typing import Annotated

from fastapi import APIRouter, Depends
from pydantic import BaseModel, ConfigDict

from app import db
from app.db import Record

router = APIRouter(prefix="/api/categories", tags=["categories"])


class AllCategoriesIdResponse(BaseModel):
    ids: list[str]

    model_config = ConfigDict(
        json_schema_extra={
            "example": {
                "ids": [
                    "clhjj1v7c0000gbgn8a5z182n",
                    "clhjj1v9k0002gbgnbfju6luj",
                    "clhjj1v9m0004gbgn0gsst205",
                ]
            }
        }
    )


@router.get("/id")
async def get_all_categories_id(
    db: Annotated[db.Connection, Depends(db.get_db)],
) -> AllCategoriesIdResponse:
    campuses: list[Record] = await db.fetch('SELECT id FROM "Category";')

    return AllCategoriesIdResponse(ids=[campus["id"] for campus in campuses])
