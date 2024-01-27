from typing import Optional

from pydantic import AnyUrl, BaseModel, ConfigDict, Field


class Affiliation(BaseModel):
    campus: str = Field(..., description="Campus name")
    department: str = Field(..., description="Department name")
    office: str = Field(..., description="Office name")


class UpdateWebsitePayload(BaseModel):
    affiliations: list[Affiliation] = Field(
        [], description="The affiliations of the website"
    )
    name: Optional[str] = Field(None, description="The name of the website")
    url: Optional[AnyUrl] = Field(None, description="The URL of the website")

    model_config: ConfigDict = {
        "json_schema_extra": {
            "example": {
                "affiliations": [
                    {
                        "campus": "交大相關",
                        "department": "行政單位",
                        "office": "圖書館",
                    }
                ],
                "name": "交通大學圖書館",
                "url": "https://www.lib.nctu.edu.tw/",
            }
        }
    }


class UpdateResponse(BaseModel):
    id: str = Field(..., description="The ID of the website")
    name: str = Field(..., description="The name of the website")
    url: AnyUrl = Field(..., description="The URL of the website")
    affiliations: list[Affiliation] = Field(
        ..., description="The affiliations of the website"
    )

    model_config: ConfigDict = {
        "json_schema_extra": {
            "example": {
                "id": "clrnc14dr000008l235gx4c6c",
                "name": "交通大學圖書館",
                "url": "https://www.lib.nctu.edu.tw/",
                "affiliations": [
                    {
                        "campus": "交大相關",
                        "department": "行政單位",
                        "office": "圖書館",
                    }
                ],
            }
        }
    }


class SearchResultWebsite(BaseModel):
    id: str = Field(..., description="The ID of the website")
    name: str = Field(..., description="The name of the website")
    url: AnyUrl = Field(..., description="The URL of the website")


class SearchResultEntry(BaseModel):
    id: str = Field(..., description="Compound ID of the website")
    campus: str = Field(..., description="Campus name")
    department: str = Field(..., description="Department name")
    office: str = Field(..., description="Office name")
    websites: list[SearchResultWebsite] = Field(
        ..., description="The websites that match the query"
    )


class Pagination(BaseModel):
    next_cursor: Optional[str] = Field(None, description="Cursor for the next page")
    num_results: int = Field(..., description="Number of results in this page")
    num_left: int = Field(..., description="Number of results left")


class SearchResponse(BaseModel):
    result: list[SearchResultEntry] = Field(..., description="The search result")
    pagination: Pagination = Field(..., description="Pagination information")

    model_config: ConfigDict = {
        "json_schema_extra": {
            "example": {
                "result": [
                    {
                        "id": "clrnc1mt7000108l21whi9y4r$clrnc1tdu000208l2dm2hcimi$clrnc1wx0000308l2h9st7aac",
                        "campus": "交大相關",
                        "department": "行政單位",
                        "office": "圖書館",
                        "websites": [
                            {
                                "id": "clrnc14dr000008l235gx4c6c",
                                "name": "交通大學圖書館",
                                "url": "https://www.lib.nctu.edu.tw/",
                            }
                        ],
                    }
                ],
                "pagination": {
                    "next_cursor": "KGNhbXB1c19pZD1jbHJuYzFtdDcwMDAxMDhsMjF3aGk5eTRyLGRlcGFydG1lbnRfaWQ9Y2xybmMxdGR1MDAwMjA4bDJkbTJoY2ltaSxvZmZpY2VfaWQ9Y2xybmMxd3gwMDAwMzA4bDJoOXN0N2FhYyx3ZWJzaXRlX2lkPWNscm5jMTRkcjAwMDAwOGwyMzVneDRjNmMp",
                    "num_results": 1,
                    "total_results": 10,
                },
            }
        }
    }
