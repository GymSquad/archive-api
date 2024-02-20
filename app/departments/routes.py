from typing import Annotated

from fastapi import APIRouter, Depends

from app import db
from app.db import Record
from app.departments.schemas import GetOfficesByDepartmentIdResponse, OfficeResponse

router = APIRouter(prefix="/departments", tags=["departments"])


@router.get("/{department_id}/offices")
async def get_offices_by_department_id(
    department_id: str, db: Annotated[db.Connection, Depends(db.get_db)]
) -> GetOfficesByDepartmentIdResponse:
    rows: list[Record] = await db.fetch(
        'SELECT id, name FROM "Office" WHERE "departmentId" = $1;', department_id
    )

    return GetOfficesByDepartmentIdResponse(
        offices=[OfficeResponse(id=row["id"], name=row["name"]) for row in rows]
    )
