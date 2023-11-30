package updatewebsite

import (
	"context"
	"database/sql"

	"github.com/GymSquad/archive-api/api"
	"github.com/GymSquad/archive-api/internal/server"
)

// UpdateWebsiteHandler is the handler for the update website endpoint
type UpdateWebsiteHandler struct {
	db *sql.DB
}

// Compile time check to ensure that UpdateWebsiteHandler implements server.UpdateWebsiteHandler.
var _ server.UpdateWebsiteHandler = (*UpdateWebsiteHandler)(nil)

// NewHTTPHandler creates a new UpdateWebsiteHandler
func NewHTTPHandler(db *sql.DB) *UpdateWebsiteHandler {
	return &UpdateWebsiteHandler{
		db: db,
	}
}

// HandleRequest implements server.SearchWebsitesHandler.
func (h *UpdateWebsiteHandler) HandleRequest(ctx context.Context, request api.UpdateWebsiteApiWebsiteWebsiteIdPatchRequestObject) (api.UpdateWebsiteApiWebsiteWebsiteIdPatchResponseObject, error) {
	return api.UpdateWebsiteApiWebsiteWebsiteIdPatch200JSONResponse(api.UpdatedWebsite{
		Id:   request.WebsiteId,
		Name: "交大圖書館",
		Url:  "https://lib.nctu.edu.tw/",
		Affiliations: []api.Affiliation{
			{
				CampusId:       "clpl85gua0000356ribtcvv9d",
				CampusName:     "交大相關",
				DepartmentId:   "clpl862ag0000356qdt2gy68z",
				DepartmentName: "行政單位",
				OfficeId:       "clpl869zn0000356qjk6mc5zu",
				OfficeName:     "圖書館",
			},
		},
	}), nil
}
