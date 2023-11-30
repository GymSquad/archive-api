package server

import (
	"context"

	"github.com/GymSquad/archive-api/api"
)

// GetArchivedDatesHandler is the interface that handles the GetApiWebsiteWebsiteId request.
type GetArchivedDatesHandler interface {
	HandleRequest(ctx context.Context, request api.GetArchivedDatesApiWebsiteWebsiteIdGetRequestObject) (api.GetArchivedDatesApiWebsiteWebsiteIdGetResponseObject, error)
}

// SearchWebsitesHandler is the interface that handles the GetApiWebsiteSearch request.
type SearchWebsitesHandler interface {
	HandleRequest(ctx context.Context, request api.SearchWebsitesApiWebsiteSearchGetRequestObject) (api.SearchWebsitesApiWebsiteSearchGetResponseObject, error)
}

// UpdateWebsiteHandler is the interface that handles the PatchApiWebsiteWebsiteId request.
type UpdateWebsiteHandler interface {
	HandleRequest(ctx context.Context, request api.UpdateWebsiteApiWebsiteWebsiteIdPatchRequestObject) (api.UpdateWebsiteApiWebsiteWebsiteIdPatchResponseObject, error)
}

type GetAllCampusesHandler interface {
	HandleRequest(ctx context.Context, request api.GetAllCampusesApiCampusGetRequestObject) (api.GetAllCampusesApiCampusGetResponseObject, error)
}

type GetAllDepartmentsHandler interface {
	HandleRequest(ctx context.Context, request api.GetAllDepartmentsApiCampusCampusIdDepartmentGetRequestObject) (api.GetAllDepartmentsApiCampusCampusIdDepartmentGetResponseObject, error)
}

type GetAllOfficesHandler interface {
	HandleRequest(ctx context.Context, request api.GetAllOfficesApiDepartmentDepartmentIdOfficeGetRequestObject) (api.GetAllOfficesApiDepartmentDepartmentIdOfficeGetResponseObject, error)
}

type archiveAPI struct {
	dates       GetArchivedDatesHandler
	search      SearchWebsitesHandler
	update      UpdateWebsiteHandler
	campuses    GetAllCampusesHandler
	departments GetAllDepartmentsHandler
	offices     GetAllOfficesHandler
}

var _ api.StrictServerInterface = (*archiveAPI)(nil)

// NewArchiveAPI constructs a new StrictServerInterface implementation with the provided handlers.
func NewArchiveAPI(
	dates GetArchivedDatesHandler,
	search SearchWebsitesHandler,
	update UpdateWebsiteHandler,
	campuses GetAllCampusesHandler,
	departments GetAllDepartmentsHandler,
	offices GetAllOfficesHandler,
) *archiveAPI {
	return &archiveAPI{
		dates:       dates,
		search:      search,
		update:      update,
		campuses:    campuses,
		departments: departments,
		offices:     offices,
	}
}

// SearchWebsitesApiWebsiteSearchGet implements api.StrictServerInterface.
func (s *archiveAPI) SearchWebsitesApiWebsiteSearchGet(ctx context.Context, request api.SearchWebsitesApiWebsiteSearchGetRequestObject) (api.SearchWebsitesApiWebsiteSearchGetResponseObject, error) {
	return s.search.HandleRequest(ctx, request)
}

// GetArchivedDatesApiWebsiteWebsiteIdGet implements api.StrictServerInterface.
func (s *archiveAPI) GetArchivedDatesApiWebsiteWebsiteIdGet(ctx context.Context, request api.GetArchivedDatesApiWebsiteWebsiteIdGetRequestObject) (api.GetArchivedDatesApiWebsiteWebsiteIdGetResponseObject, error) {
	return s.dates.HandleRequest(ctx, request)
}

// PatchApiWebsiteWebsiteId implements api.StrictServerInterface.
func (s *archiveAPI) UpdateWebsiteApiWebsiteWebsiteIdPatch(ctx context.Context, request api.UpdateWebsiteApiWebsiteWebsiteIdPatchRequestObject) (api.UpdateWebsiteApiWebsiteWebsiteIdPatchResponseObject, error) {
	return s.update.HandleRequest(ctx, request)
}

// GetAllCampusesApiCampusGet implements api.StrictServerInterface.
func (s *archiveAPI) GetAllCampusesApiCampusGet(ctx context.Context, request api.GetAllCampusesApiCampusGetRequestObject) (api.GetAllCampusesApiCampusGetResponseObject, error) {
	return s.campuses.HandleRequest(ctx, request)
}

// GetAllDepartmentsApiCampusCampusIdDepartmentGet implements api.StrictServerInterface.
func (s *archiveAPI) GetAllDepartmentsApiCampusCampusIdDepartmentGet(ctx context.Context, request api.GetAllDepartmentsApiCampusCampusIdDepartmentGetRequestObject) (api.GetAllDepartmentsApiCampusCampusIdDepartmentGetResponseObject, error) {
	return s.departments.HandleRequest(ctx, request)
}

// GetAllOfficesApiDepartmentDepartmentIdOfficeGet implements api.StrictServerInterface.
func (s *archiveAPI) GetAllOfficesApiDepartmentDepartmentIdOfficeGet(ctx context.Context, request api.GetAllOfficesApiDepartmentDepartmentIdOfficeGetRequestObject) (api.GetAllOfficesApiDepartmentDepartmentIdOfficeGetResponseObject, error) {
	return s.offices.HandleRequest(ctx, request)
}
