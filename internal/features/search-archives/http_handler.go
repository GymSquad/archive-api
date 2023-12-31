package searcharchives

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/GymSquad/archive-api/api"
	"github.com/GymSquad/archive-api/internal/server"
)

// SearchArchivesHandler is the handler for the search archives endpoint
type SearchArchivesHandler struct {
	db *sql.DB
}

// Compile time check to ensure that SearchArchivesHandler implements server.SearchWebsitesHandler.
var _ server.SearchWebsitesHandler = (*SearchArchivesHandler)(nil)

// NewHTTPHandler creates a new SearchArchivesHandler
func NewHTTPHandler(db *sql.DB) *SearchArchivesHandler {
	return &SearchArchivesHandler{
		db: db,
	}
}

// HandleRequest implements server.SearchWebsitesHandler.
func (*SearchArchivesHandler) HandleRequest(ctx context.Context, request api.GetApiWebsiteSearchRequestObject) (api.GetApiWebsiteSearchResponseObject, error) {
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
