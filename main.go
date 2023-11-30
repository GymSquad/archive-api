package main

import (
	"os"

	"log/slog"

	"github.com/GymSquad/archive-api/api"
	getarchiveddates "github.com/GymSquad/archive-api/internal/features/get-archived-dates"
	searcharchives "github.com/GymSquad/archive-api/internal/features/search-archives"
	searcharchivesquery "github.com/GymSquad/archive-api/internal/features/search-archives/query"
	swaggerui "github.com/GymSquad/archive-api/internal/features/swagger-ui"
	updatewebsite "github.com/GymSquad/archive-api/internal/features/update-website"
	updatewebsitecommand "github.com/GymSquad/archive-api/internal/features/update-website/command"
	"github.com/GymSquad/archive-api/internal/server"
	"github.com/caarlos0/env/v10"
	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	RootPath string `env:"ROOT_PATH" envDefault:"/archive"`

	DbHost     string `env:"DB_HOST" envDefault:"localhost"`
	DbPort     string `env:"DB_PORT" envDefault:"5432"`
	DbUser     string `env:"DB_USER" envDefault:"app"`
	DbPassword string `env:"DB_PASSWORD" envDefault:"app"`
	DbName     string `env:"DB_NAME" envDefault:"db"`
}

func main() {
	cfg := NewConfig()

	datesHandler := getarchiveddates.NewHTTPHandler(cfg.RootPath)

	dummyQuery := searcharchivesquery.NewDummyQuery()
	searchHandler := searcharchives.NewHTTPHandler(dummyQuery)

	dummyCommand := updatewebsitecommand.NewDummyCommand()
	updateHandler := updatewebsite.NewHTTPHandler(dummyCommand)

	aApi := server.NewArchiveAPI(datesHandler, searchHandler, updateHandler, nil, nil, nil)
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

// NewConfig loads the environment variables into a config object and returns it.
func NewConfig() *config {
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		slog.Error("failed to parse env", "err", err)
		os.Exit(1)
	}
	return &cfg
}
