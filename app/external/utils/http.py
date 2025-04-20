import json
from typing import TypeVar

import httpx

T = TypeVar("T")


def parse_jsonl_response(response: httpx.Response, cls: type[T]) -> list[T]:
    """
    Parse a JSON Lines (jsonl) response into a list of objects.
    Each line in the response should be a valid JSON object.
    """
    data: list[T] = []
    for line in response.iter_lines():
        if line.strip():  # Skip empty lines
            # Parse the JSON object from the line
            data.append(cls(**json.loads(line)))

    return data