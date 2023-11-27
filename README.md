# NYCU Library Web Archive API

A minimal API for the NYCU Library Web Archive project.

## Getting Started

### Prerequisites

- Golang 1.21

### Installing

1. Clone the repository

   ```bash
   git clone https://github.com/GymSquad/archive-api.git
   ```

2. Install dependencies

   ```bash
   make deps
   ```

3. Run the server

   ```bash
   make run
   ```

   Now you should be able to access the API at [`http://localhost:8080`](http://localhost:8080) and the API documentation at [`http://localhost:8080/docs/index.html`](http://localhost:8080/docs/index.html).

## Development

We use oapi-codegen to generate the API code from the OpenAPI specification. To regenerate the code, run:

```bash
make generate
```

### Seed the database

```bash
docker compose exec -T db psql -U ${DB_USER:-app} < /path/to/dump-file.sql
```
