import json
from pathlib import Path

from app.app import app


def main():
    dump_path = Path(__file__).parent.parent / "openapi.json"
    schema = app.openapi()

    with open(dump_path, "w") as schema_file:
        json.dump(schema, schema_file, indent=2)


if __name__ == "__main__":
    main()
