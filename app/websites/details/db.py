from dataclasses import dataclass

import structlog
from pydantic import BaseModel, ConfigDict

from app import db

logger = structlog.get_logger("websites.details.db")


@dataclass
class AffiliationId:
    campus_id: str
    department_id: str
    office_id: str


class WebsiteInfo(BaseModel):
    website_name: str
    website_url: str
    affiliationIds: list[AffiliationId]

    model_config: ConfigDict = {
        "json_schema_extra": {
            "example": {
                "website_name": "圖書館",
                "website_url": "https://lib.nycu.edu.tw/",
                "affiliationIds": [
                    {
                        "campus_id": "clhjj1v9k0002gbgnbfju6luj",
                        "department_id": "clhjj1vgl000pgbgnuso925sb",
                        "office_id": "clhjj1vqs0038gbgn9cq37eyc",
                    },
                    {
                        "campus_id": "clhjj1v9k0002gbgnbfju6luj",
                        "department_id": "clhjj1vgl000pgbgnuso925sb",
                        "office_id": "clhjj1vra0043gbgnh8wnz32k",
                    },
                ],
            }
        }
    }


async def get_website_info_by_id(conn: db.Connection, website_id: str) -> WebsiteInfo:
    rows: list[db.Record] = await conn.fetch(
        """
        SELECT
            web.name AS website_name,
            web.url AS website_url,
            off.id AS office_id,
            dep.id AS department_id,
            cat.id AS campus_id
        FROM
            "Website" web
        JOIN
            "_OfficeToWebsite" otw ON web.id = otw."B"
        JOIN
            "Office" off ON otw."A" = off.id
        JOIN
            "Department" dep ON off."departmentId" = dep.id
        JOIN
            "Category" cat ON dep."categoryId" = cat.id
        WHERE
            web.id = $1;
        """,
        website_id,
    )

    if len(rows) == 0:
        logger.error("website not found", website_id=website_id)
        raise ValueError(f"website_id {website_id} not found")

    return WebsiteInfo(
        website_name=rows[0]["website_name"],
        website_url=rows[0]["website_url"],
        affiliationIds=(
            [
                AffiliationId(
                    campus_id=row["campus_id"],
                    department_id=row["department_id"],
                    office_id=row["office_id"],
                )
                for row in rows
            ]
        ),
    )
