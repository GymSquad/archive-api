package updatewebsite

import (
	"context"

	"github.com/GymSquad/archive-api/api"
	"github.com/GymSquad/archive-api/internal/server"
)

type Affiliation struct {
	CampusID     string
	DepartmentID string
	OfficeID     string
}

type UpdateWebsitePayload struct {
	// WebsiteID is the ID of the website to update
	WebsiteID string

	// Offices is the new list of offices if it should be updated
	Offices *[]string
	// Name is the new name of the website if it should be updated
	Name *string
	// Url is the new URL of the website if it should be updated
	Url *string
}

type UpdatedAffiliation struct {
	CampusID       string
	CampusName     string
	DepartmentID   string
	DepartmentName string
	OfficeID       string
	OfficeName     string
}

// ToTransport converts the UpdatedAffiliation to the transport representation
func (a *UpdatedAffiliation) ToTransport() api.Affiliation {
	return api.Affiliation{
		CampusId:       a.CampusID,
		CampusName:     a.CampusName,
		DepartmentId:   a.DepartmentID,
		DepartmentName: a.DepartmentName,
		OfficeId:       a.OfficeID,
		OfficeName:     a.OfficeName,
	}
}

type UpdatedWebsite struct {
	ID           string
	Affiliations []UpdatedAffiliation
	Name         string
	Url          string
}

type UpdateWebisteCommand interface {
	Execute(context.Context, UpdateWebsitePayload) (UpdatedWebsite, error)
}

// UpdateWebsiteHandler is the handler for the update website endpoint
type UpdateWebsiteHandler struct {
	update UpdateWebisteCommand
}

// Compile time check to ensure that UpdateWebsiteHandler implements server.UpdateWebsiteHandler.
var _ server.UpdateWebsiteHandler = (*UpdateWebsiteHandler)(nil)

// NewHTTPHandler creates a new UpdateWebsiteHandler
func NewHTTPHandler(command UpdateWebisteCommand) *UpdateWebsiteHandler {
	return &UpdateWebsiteHandler{
		update: command,
	}
}

// HandleRequest implements server.SearchWebsitesHandler.
func (h *UpdateWebsiteHandler) HandleRequest(ctx context.Context, request api.UpdateWebsiteApiWebsiteWebsiteIdPatchRequestObject) (api.UpdateWebsiteApiWebsiteWebsiteIdPatchResponseObject, error) {
	website, err := h.update.Execute(ctx, UpdateWebsitePayload{
		WebsiteID: request.WebsiteId,
		Offices:   request.Body.OfficeIds,
		Name:      request.Body.Name,
		Url:       request.Body.Url,
	})
	if err != nil {
		return nil, err
	}

	affiliations := []api.Affiliation{}
	for _, affiliation := range website.Affiliations {
		affiliations = append(affiliations, affiliation.ToTransport())
	}

	return api.UpdateWebsiteApiWebsiteWebsiteIdPatch200JSONResponse(api.UpdatedWebsite{
		Id:           website.ID,
		Name:         website.Name,
		Url:          website.Url,
		Affiliations: affiliations,
	}), nil
}
