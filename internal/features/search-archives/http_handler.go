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

// Compile time check to ensure that SearchWebsitesHandler implements server.SearchWebsitesHandler.
var _ server.SearchWebsitesHandler = (*SearchWebsitesHandler)(nil)

// NewHTTPHandler creates a new SearchArchivesHandler
func NewHTTPHandler(query SearchArchivesQuery) *SearchWebsitesHandler {
	return &SearchWebsitesHandler{
		query: query,
	}
}

// HandleRequest implements server.SearchWebsitesHandler.
func (s *SearchWebsitesHandler) HandleRequest(ctx context.Context, request api.SearchWebsitesApiWebsiteSearchGetRequestObject) (api.SearchWebsitesApiWebsiteSearchGetResponseObject, error) {
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
			return api.SearchWebsitesApiWebsiteSearchGet400JSONResponse(api.InvalidPayload{
				Error: "invalid cursor",
			}), nil
		}
		return nil, errors.New("internal server error")
	}

	return api.SearchWebsitesApiWebsiteSearchGet200JSONResponse(api.SearchResult{
		Result:     result,
		Pagination: pagination,
	}), nil
}
