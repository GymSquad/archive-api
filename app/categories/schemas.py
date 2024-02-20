from pydantic import BaseModel, ConfigDict


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


class DepartmentResponse(BaseModel):
    id: str
    name: str


class GetDepartmentsByCategoryIdResponse(BaseModel):
    departments: list[DepartmentResponse]

    model_config = ConfigDict(
        json_schema_extra={
            "example": {
                "departments": [
                    {"id": "clhjj1vfa000jgbgn22c3lvyh", "name": "學術單位"},
                    {"id": "clhjj1vfp000mgbgnrxv5wvpp", "name": "行政單位"},
                    {"id": "clhjj1vg2000ngbgnz7hf002d", "name": "研究單位"},
                    {"id": "clhjj1vgp000rgbgnxwkh39rg", "name": "學生活動"},
                ]
            }
        }
    )
