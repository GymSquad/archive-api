package searcharchives

import (
	"context"
	"errors"

	"github.com/GymSquad/archive-api/api"
	"github.com/GymSquad/archive-api/internal/server"
)

var (
	// ErrorInvalidCursor is returned when the cursor is invalid
	ErrorInvalidCursor = errors.New("invalid cursor")
	// ErrorInternal is returned when an internal error occurs
	ErrorInternal = errors.New("internal error")
)

type SearchArchivesQuery interface {
	SearchArchives(ctx context.Context, query string, cursor *string, limit int) ([]api.SearchResultEntry, api.Pagination, error)
}

// SearchWebsitesHandler is the handler for the search websites endpoint
type SearchWebsitesHandler struct {
	query SearchArchivesQuery
}

// Compile time check to ensure that SearchArchivesHandler implements server.SearchWebsitesHandler.
var _ server.SearchWebsitesHandler = (*SearchWebsitesHandler)(nil)

// NewHTTPHandler creates a new SearchArchivesHandler
func NewHTTPHandler(query SearchArchivesQuery) *SearchWebsitesHandler {
	return &SearchWebsitesHandler{
		query: query,
	}
}

// HandleRequest implements server.SearchWebsitesHandler.
func (s *SearchWebsitesHandler) HandleRequest(ctx context.Context, request api.GetApiWebsiteSearchRequestObject) (api.GetApiWebsiteSearchResponseObject, error) {
	query := ""
	if request.Params.Q != nil {
		query = *request.Params.Q
	}

	limit := 10
	if request.Params.Limit != nil {
		limit = *request.Params.Limit
	}

	result, pagination, err := s.query.SearchArchives(ctx, query, request.Params.Cursor, limit)
	if err != nil {
		if errors.Is(err, ErrorInvalidCursor) {
			apiError := "invalid cursor"
			return api.GetApiWebsiteSearch400JSONResponse{
				Error: &apiError,
			}, nil
		}
		apiError := "internal server error"
		return api.GetApiWebsiteSearch500JSONResponse{
			Error: &apiError,
		}, nil
	}

	return api.GetApiWebsiteSearch200JSONResponse(api.WebsiteSearchResult{
		Result:     result,
		Pagination: pagination,
	}), nil
}
