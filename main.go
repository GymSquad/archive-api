package main

import (
	"database/sql"
	"fmt"
	"os"

	"log/slog"

	"github.com/GymSquad/archive-api/api"
	getarchiveddates "github.com/GymSquad/archive-api/internal/features/get-archived-dates"
	searcharchives "github.com/GymSquad/archive-api/internal/features/search-archives"
	searcharchivesquery "github.com/GymSquad/archive-api/internal/features/search-archives/query"
	swaggerui "github.com/GymSquad/archive-api/internal/features/swagger-ui"
	updatewebsite "github.com/GymSquad/archive-api/internal/features/update-website"
	"github.com/GymSquad/archive-api/internal/server"
	"github.com/caarlos0/env/v10"
	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type config struct {
	RootPath string `env:"ROOT_PATH" envDefault:"/archive"`

	DbHost     string `env:"DB_HOST" envDefault:"localhost"`
	DbPort     string `env:"DB_PORT" envDefault:"5432"`
	DbUser     string `env:"DB_USER" envDefault:"app"`
	DbPassword string `env:"DB_PASSWORD" envDefault:"app"`
	DbName     string `env:"DB_NAME" envDefault:"db"`
}

const searchArchivesQuery = `
SELECT
	c.id AS campus_id,
	c.name AS campus_name,
	d.id AS department_id,
	d.name AS department_name,
	o.id AS office_id,
	o.name AS office_name,
	w.id AS website_id,
	w.name AS website_name,
	w.url AS website_url,
	COUNT(*) OVER() AS total_results
FROM 
	"Category" c
JOIN
	"Department" d ON c.id = d."categoryId"
JOIN
	"Office" o ON d.id = o."departmentId"
JOIN
	"_OfficeToWebsite" otw ON o.id = otw."A"
JOIN
	"Website" w ON otw."B" = w.id
WHERE 
	(
		c.id > $1
		OR (c.id = $1 AND d.id > $2)
		OR (c.id = $1 AND d.id = $2 AND o.id > $3)
		OR (c.id = $1 AND d.id = $2 AND o.id = $3 AND w.id > $4)
	)
	AND 
	(
		$5::TEXT IS NULL
		OR c.name ILIKE '%' || $5 || '%'
		OR d.name ILIKE '%' || $5 || '%'
		OR o.name ILIKE '%' || $5 || '%'
		OR w.name ILIKE '%' || $5 || '%'
		OR w.url ILIKE '%' || $5 || '%'
	)
ORDER BY
	c.id, d.id, o.id, w.id
LIMIT
	$6;`

func main() {
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		slog.Error("failed to parse env", "err", err)
		os.Exit(1)
	}

	slog.Info("parsed env", "cfg", cfg)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		slog.Error("failed to open db", "err", err)
		os.Exit(1)
	}
	defer db.Close()

	datesHandler := getarchiveddates.NewHTTPHandler(cfg.RootPath)

	// dummyQuery := searcharchivesquery.NewDummyQuery()
	postgresQuery := searcharchivesquery.NewPostgresQuery(db)
	searchHandler := searcharchives.NewHTTPHandler(postgresQuery)
	updateHandler := updatewebsite.NewHTTPHandler(db)

	aApi := server.NewArchiveAPI(datesHandler, searchHandler, updateHandler)
	strictApiHandler := api.NewStrictHandler(aApi, nil)

	r := gin.Default()
	api.RegisterHandlers(r, strictApiHandler)

	// Register swagger handlers
	swaggerHandler, err := swaggerui.DefaultHandler()
	if err != nil {
		slog.Error("failed to create swagger handler", "err", err)
		os.Exit(1)
	}
	swaggerHandler.Register(r)

	if err := r.Run(":8080"); err != nil {
		slog.Error("failed to run gin server", "err", err)
		os.Exit(1)
	}
}
