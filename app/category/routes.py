from typing import Annotated

from fastapi import APIRouter, Depends

from app import db
from app.db import Record

router = APIRouter(prefix="/api/category", tags=["category"])


@router.get("/id")
async def get_all_categories_id(
    db: Annotated[db.Connection, Depends(db.get_db)],
) -> list[str]:
    campuses: list[Record] = await db.fetch('SELECT id FROM "Category";')

    return [campus["id"] for campus in campuses]
