from pydantic import BaseModel, Field


class PywbArchiveInfo(BaseModel):
    urlkey: str
    timestamp: str
    url: str
    mime: str
    status: str
    digest: str
    length: str
    offset: str
    filename: str
    source: str
    source_coll: str = Field(alias="source-coll")
    access: str
