package main

import (
	"github.com/GymSquad/archive-api/api"
	swaggerui "github.com/GymSquad/archive-api/internal/features/swagger-ui"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	// RootPath is the path to the root of the archive``
	RootPath = "/archive"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())

	s := &svc{}
	api.RegisterHandlers(e, s)

	// Register swagger handlers
	swaggerHandler, err := swaggerui.DefaultHandler()
	if err != nil {
		e.Logger.Fatalf("Failed to create swagger handler: %v", err)
	}
	swaggerHandler.Register(e)

	e.Logger.Fatal(e.Start(":8080"))
}

type svc struct{}

// GetApiWebsiteSearch implements archiveapi.ServerInterface.
func (*svc) GetApiWebsiteSearch(c echo.Context, params api.GetApiWebsiteSearchParams) error {
	panic("unimplemented")
}

// GetApiWebsiteWebsiteId implements archiveapi.ServerInterface.
func (*svc) GetApiWebsiteWebsiteId(c echo.Context, websiteId string) error {
	panic("unimplemented")
}

// PatchApiWebsiteWebsiteId implements archiveapi.ServerInterface.
func (*svc) PatchApiWebsiteWebsiteId(c echo.Context, websiteId string) error {
	panic("unimplemented")
}

var _ api.ServerInterface = (*svc)(nil)
