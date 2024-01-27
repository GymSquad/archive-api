# Web Archive API

### Installing

1. Clone the repo

2. Install dependencies

```bash
pdm install
```

3. Configure environment variables in `.env` (see [`.env.example`](.env.example))

4. Start database

```bash
docker compose up
```

5. If your volume is empty, you need to seed the database 

```bash
docker compose exec -T db psql -U ${DB_USER:-app} ${DB_NAME:-db} < /path/to/dump-file.sql
```
