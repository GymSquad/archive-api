package server

import (
	"context"

	"github.com/GymSquad/archive-api/api"
)

// GetArchivedDatesHandler is the interface that handles the GetApiWebsiteWebsiteId request.
type GetArchivedDatesHandler interface {
	HandleRequest(ctx context.Context, request api.GetApiWebsiteWebsiteIdRequestObject) (api.GetApiWebsiteWebsiteIdResponseObject, error)
}

// SearchWebsitesHandler is the interface that handles the GetApiWebsiteSearch request.
type SearchWebsitesHandler interface {
	HandleRequest(ctx context.Context, request api.GetApiWebsiteSearchRequestObject) (api.GetApiWebsiteSearchResponseObject, error)
}

// UpdateWebsiteHandler is the interface that handles the PatchApiWebsiteWebsiteId request.
type UpdateWebsiteHandler interface {
	HandleRequest(ctx context.Context, request api.PatchApiWebsiteWebsiteIdRequestObject) (api.PatchApiWebsiteWebsiteIdResponseObject, error)
}

type archiveAPI struct {
	dates  GetArchivedDatesHandler
	search SearchWebsitesHandler
	update UpdateWebsiteHandler
}

var _ api.StrictServerInterface = (*archiveAPI)(nil)

// NewArchiveAPI constructs a new StrictServerInterface implementation with the provided handlers.
func NewArchiveAPI(
	getArchivedDatesHandler GetArchivedDatesHandler,
	searchWebsitesHandler SearchWebsitesHandler,
	updateWebsiteHandler UpdateWebsiteHandler,
) *archiveAPI {
	return &archiveAPI{
		dates:  getArchivedDatesHandler,
		search: searchWebsitesHandler,
		update: updateWebsiteHandler,
	}
}

// GetApiWebsiteSearch implements api.StrictServerInterface.
func (s *archiveAPI) GetApiWebsiteSearch(ctx context.Context, request api.GetApiWebsiteSearchRequestObject) (api.GetApiWebsiteSearchResponseObject, error) {
	return s.search.HandleRequest(ctx, request)
}

// GetApiWebsiteWebsiteId implements api.StrictServerInterface.
func (s *archiveAPI) GetApiWebsiteWebsiteId(ctx context.Context, request api.GetApiWebsiteWebsiteIdRequestObject) (api.GetApiWebsiteWebsiteIdResponseObject, error) {
	return s.dates.HandleRequest(ctx, request)
}

// PatchApiWebsiteWebsiteId implements api.StrictServerInterface.
func (s *archiveAPI) PatchApiWebsiteWebsiteId(ctx context.Context, request api.PatchApiWebsiteWebsiteIdRequestObject) (api.PatchApiWebsiteWebsiteIdResponseObject, error) {
	return s.update.HandleRequest(ctx, request)
}
