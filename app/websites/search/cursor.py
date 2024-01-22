import base64
import binascii
import string
from typing import Literal, Self

from pydantic import BaseModel, ConfigDict, ValidationError
from result import Err, Ok, Result

type ParseCursorErrorType = Literal[
    "base64_decode_error",
    "unicode_decode_error",
    "invalid_cursor_format",
    "validation_failure",
]


class ParseCursorError(BaseModel):
    type: ParseCursorErrorType
    msg: str

    model_config: ConfigDict = {"extra": "allow"}


class SearchWebsiteCursor(BaseModel):
    campus_id: str = ""
    department_id: str = ""
    office_id: str = ""
    website_id: str = ""

    def to_base64(self) -> str:
        str_pairs = ",".join(
            f"{key}={value}" for key, value in self.model_dump().items()
        )
        return base64.b64encode(f"({str_pairs})".encode()).decode()

    @classmethod
    def from_base64(cls, base64_str: str) -> Result[Self, ParseCursorError]:
        try:
            decoded = base64.b64decode(base64_str, validate=True)
        except binascii.Error as e:
            return Err(ParseCursorError(type="base64_decode_error", msg=str(e)))

        try:
            decoded = decoded.decode()
        except UnicodeDecodeError as e:
            return Err(ParseCursorError(type="unicode_decode_error", msg=str(e)))

        str_pairs = decoded.strip(string.whitespace + "()").split(",")
        if not all(pair.count("=") == 1 for pair in str_pairs):
            return Err(
                ParseCursorError(
                    type="invalid_cursor_format",
                    msg="invalid key-value pair format",
                )
            )

        if (n := len(str_pairs)) != 4:
            return Err(
                ParseCursorError.model_validate(
                    {
                        "type": "invalid_cursor_format",
                        "msg": "invalid number of pairs",
                        "got": n,
                        "expected": 4,
                    }
                )
            )

        try:
            return Ok(cls.model_validate(dict(pair.split("=") for pair in str_pairs)))
        except ValidationError as e:
            return Err(ParseCursorError(type="validation_failure", msg=str(e)))
