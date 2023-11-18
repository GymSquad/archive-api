package main

import (
	"log/slog"

	_ "github.com/GymSquad/archive-api/docs"
	getarchiveddates "github.com/GymSquad/archive-api/internal/features/get-archived-dates"
	searcharchives "github.com/GymSquad/archive-api/internal/features/search-archives"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	// RootPath is the path to the root of the archive``
	RootPath = "/archive"
)

// @title NYCU Library Web Archive API
// @description Minimal API for the NYCU Library Web Archive project
// @version 1.0.0

// @host localhost:8080
// @BasePath /api
func main() {
	r := gin.Default()

	getArchiveDatesHandler := getarchiveddates.NewHTTPHandler(RootPath)
	searcharchivesHandler := searcharchives.NewHTTPHandler(nil)

	g := r.Group("/api")
	RegisterRoutes(g, getArchiveDatesHandler, searcharchivesHandler)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	slog.Info("Starting server at", "url", "http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		slog.Error("Failed to start server", "err", err)
	}
}

type Routable interface {
	RegisterRoutes(router gin.IRouter)
}

func RegisterRoutes(router gin.IRouter, handlers ...Routable) {
	for _, handler := range handlers {
		handler.RegisterRoutes(router)
	}
}
