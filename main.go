package main

import (
	"context"
	"os"
	"strconv"
	"time"

	"log/slog"

	"github.com/GymSquad/archive-api/api"
	swaggerui "github.com/GymSquad/archive-api/internal/features/swagger-ui"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime/types"
)

const (
	// RootPath is the path to the root of the archive``
	RootPath = "/archive"
)

func main() {
	r := gin.Default()

	s := &svc{}
	strictApiHandler := api.NewStrictHandler(s, nil)
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

type svc struct{}

// GetApiWebsiteSearch implements api.StrictServerInterface.
func (*svc) GetApiWebsiteSearch(ctx context.Context, request api.GetApiWebsiteSearchRequestObject) (api.GetApiWebsiteSearchResponseObject, error) {
	var result []api.SearchResultEntry
	var pagination api.Pagination

	result = append(result, api.SearchResultEntry{
		Id:         "nctu-administration-library",
		Campus:     "交大相關",
		Department: "行政單位",
		Office:     "圖書館",
		Websites:   []api.SearchResultWebsiteEntry{},
	})

	for i := 0; i < 10; i++ {
		result[0].Websites = append(result[0].Websites, api.SearchResultWebsiteEntry{
			Id:   strconv.Itoa(i + 1),
			Name: "交大圖書館",
			Url:  "https://lib.nctu.edu.tw/",
		})
	}

	pagination = api.Pagination{
		NextCursor:   nil,
		NumResults:   10,
		TotalResults: 10,
	}

	return api.GetApiWebsiteSearch200JSONResponse(api.WebsiteSearchResult{
		Result:     result,
		Pagination: pagination,
	}), nil
}

// GetApiWebsiteWebsiteId implements api.StrictServerInterface.
func (*svc) GetApiWebsiteWebsiteId(ctx context.Context, request api.GetApiWebsiteWebsiteIdRequestObject) (api.GetApiWebsiteWebsiteIdResponseObject, error) {
	var dates []types.Date

	date, err := time.Parse(types.DateFormat, "2021-01-01")
	if err != nil {
		return nil, err
	}
	dates = append(dates, types.Date{Time: date})

	return api.GetApiWebsiteWebsiteId200JSONResponse(api.ArchivedDates(dates)), nil
}

// PatchApiWebsiteWebsiteId implements api.StrictServerInterface.
func (*svc) PatchApiWebsiteWebsiteId(ctx context.Context, request api.PatchApiWebsiteWebsiteIdRequestObject) (api.PatchApiWebsiteWebsiteIdResponseObject, error) {
	return api.PatchApiWebsiteWebsiteId200JSONResponse(api.Website{
		Id:         request.WebsiteId,
		Campus:     "交大相關",
		Department: "行政單位",
		Office:     "圖書館",
		Name:       "交大圖書館",
		Url:        "https://lib.nctu.edu.tw/",
	}), nil
}

var _ api.StrictServerInterface = (*svc)(nil)
