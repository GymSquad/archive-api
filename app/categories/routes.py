from typing import Annotated

from fastapi import APIRouter, Depends

from app import db
from app.categories.schemas import (
    AllCategoriesIdResponse,
    DepartmentResponse,
    GetDepartmentsByCategoryIdResponse,
)
from app.db import Record

router = APIRouter(prefix="/api/categories", tags=["categories"])


@router.get("/id")
async def get_all_categories_id(
    db: Annotated[db.Connection, Depends(db.get_db)],
) -> AllCategoriesIdResponse:
    campuses: list[Record] = await db.fetch('SELECT id FROM "Category";')

    return AllCategoriesIdResponse(ids=[campus["id"] for campus in campuses])


@router.get("/{category_id}/departments")
async def get_departments_by_category_id(
    category_id: str,
    db: Annotated[db.Connection, Depends(db.get_db)],
) -> GetDepartmentsByCategoryIdResponse:
    rows: list[Record] = await db.fetch(
        'SELECT id, name FROM "Department" WHERE "categoryId" = $1;', category_id
    )

    return GetDepartmentsByCategoryIdResponse(
        departments=[DepartmentResponse(id=row["id"], name=row["name"]) for row in rows]
    )
