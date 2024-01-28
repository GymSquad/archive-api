import json

import click

from app.app import app


@click.command(name="dump-schema", help="Dump OpenAPI schema to a file")
@click.option("--dump-path", "-D", type=click.Path(), default="openapi.json", help="Path to dump OpenAPI schema")
def dump_schema(dump_path: str):
    schema = app.openapi()

    with open(dump_path, "w") as schema_file:
        json.dump(schema, schema_file, indent=2)


if __name__ == "__main__":
    dump_schema()
