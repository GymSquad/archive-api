from pydantic import BaseModel, ConfigDict


class OfficeResponse(BaseModel):
    id: str
    name: str


class GetOfficesByDepartmentIdResponse(BaseModel):
    offices: list[OfficeResponse]

    model_config = ConfigDict(
        json_schema_extra={
            "example": {
                "offices": [
                    {"id": "clhjj1voh000tgbgng77i1wp5", "name": "教務處"},
                    {"id": "clhjj1voi000ugbgnqma6bphq", "name": "校長室"},
                ]
            }
        }
    )
