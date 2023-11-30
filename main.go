package main

import (
	"os"

	"log/slog"

	"github.com/GymSquad/archive-api/api"
	getarchiveddates "github.com/GymSquad/archive-api/internal/features/get-archived-dates"
	searcharchives "github.com/GymSquad/archive-api/internal/features/search-archives"
	swaggerui "github.com/GymSquad/archive-api/internal/features/swagger-ui"
	updatewebsite "github.com/GymSquad/archive-api/internal/features/update-website"
	"github.com/GymSquad/archive-api/internal/server"
	"github.com/gin-gonic/gin"
)

const (
	// DefaultRootPath is the default root path for the archive
	DefaultRootPath = "/archive"
)

func main() {
	var rootPath string
	if path, ok := os.LookupEnv("ROOT_PATH"); ok {
		rootPath = path
	} else {
		rootPath = DefaultRootPath
	}

	r := gin.Default()

	datesHandler := getarchiveddates.NewHTTPHandler(cfg.RootPath)

	dummyQuery := searcharchivesquery.NewDummyQuery()
	searchHandler := searcharchives.NewHTTPHandler(dummyQuery)

	dummyCommand := updatewebsitecommand.NewDummyCommand()
	updateHandler := updatewebsite.NewHTTPHandler(dummyCommand)

	aApi := server.NewArchiveAPI(datesHandler, searchHandler, updateHandler, nil, nil, nil)
	strictApiHandler := api.NewStrictHandler(aApi, nil)
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
