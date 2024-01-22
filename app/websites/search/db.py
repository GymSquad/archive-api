from dataclasses import dataclass

import structlog
from pydantic import BaseModel

from app import db

from .cursor import SearchWebsiteCursor


class SearchResultEntry(BaseModel):
    campus_id: str
    campus_name: str
    department_id: str
    department_name: str
    office_id: str
    office_name: str
    website_id: str
    website_name: str
    website_url: str


@dataclass
class SearchWebsitesResult:
    entries: list[SearchResultEntry]
    num_left: int
    next_cursor: SearchWebsiteCursor | None


logger = structlog.get_logger("websites.search.db")


async def search_websites(
    conn: db.Connection,
    query: str,
    cursor: SearchWebsiteCursor,
    limit: int,
) -> SearchWebsitesResult:
    records: list[db.Record] = await conn.fetch(
        """
        SELECT
            c.id AS campus_id, c.name AS campus_name,
            d.id AS department_id, d.name AS department_name,
            o.id AS office_id, o.name AS office_name,
            w.id AS website_id, w.name AS website_name, w.url AS website_url,
            COUNT(*) OVER() AS num_left
        FROM
            "Category" c
        JOIN
            "Department" d ON c.id = d."categoryId"
        JOIN
            "Office" o ON d.id = o."departmentId"
        JOIN
            "_OfficeToWebsite" otw ON o.id = otw."A"
        JOIN
            "Website" w ON otw."B" = w.id
        WHERE
            (
                c.id > $1
                OR (c.id = $1 AND d.id > $2)
                OR (c.id = $1 AND d.id = $2 AND o.id > $3)
                OR (c.id = $1 AND d.id = $2 AND o.id = $3 AND w.id > $4)
            )
            AND
            (
                $5::TEXT IS NULL
                OR c.name ILIKE '%' || $5 || '%'
                OR d.name ILIKE '%' || $5 || '%'
                OR o.name ILIKE '%' || $5 || '%'
                OR w.name ILIKE '%' || $5 || '%'
                OR w.url ILIKE '%' || $5 || '%'
            )
        ORDER BY
            c.id ASC,
            d.id ASC,
            o.id ASC,
            w.id ASC
        LIMIT $6;
        """,
        cursor.campus_id,
        cursor.department_id,
        cursor.office_id,
        cursor.website_id,
        query,
        # intentionally overfetching by 1 to determine if there are more results
        limit + 1,
    )

    logger.info("search results", num_records=len(records), limit=limit)

    num_left = records[0]["num_left"] if len(records) > 0 else 0

    next_cursor: SearchWebsiteCursor | None = None
    if len(records) > limit:
        logger.info("more results available", discarded=records[-1])
        records = records[:limit]
        next_cursor = SearchWebsiteCursor(
            campus_id=records[-1]["campus_id"],
            department_id=records[-1]["department_id"],
            office_id=records[-1]["office_id"],
            website_id=records[-1]["website_id"],
        )

    entries = [
        SearchResultEntry(
            campus_id=r["campus_id"],
            campus_name=r["campus_name"],
            department_id=r["department_id"],
            department_name=r["department_name"],
            office_id=r["office_id"],
            office_name=r["office_name"],
            website_id=r["website_id"],
            website_name=r["website_name"],
            website_url=r["website_url"],
        )
        for r in records
    ]

    return SearchWebsitesResult(
        entries=entries,
        num_left=num_left,
        next_cursor=next_cursor,
    )
