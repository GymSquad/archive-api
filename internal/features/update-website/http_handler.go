package updatewebsite

import (
	"context"

	"github.com/GymSquad/archive-api/api"
	"github.com/GymSquad/archive-api/internal/db"
	"github.com/GymSquad/archive-api/internal/server"
)

// UpdateWebsiteHandler is the handler for the update website endpoint
type UpdateWebsiteHandler struct {
	db db.DBTX
}

// Compile time check to ensure that UpdateWebsiteHandler implements server.UpdateWebsiteHandler.
var _ server.UpdateWebsiteHandler = (*UpdateWebsiteHandler)(nil)

// NewHTTPHandler creates a new UpdateWebsiteHandler
func NewHTTPHandler(db db.DBTX) *UpdateWebsiteHandler {
	return &UpdateWebsiteHandler{
		db: db,
	}
}

// HandleRequest implements server.SearchWebsitesHandler.
func (*UpdateWebsiteHandler) HandleRequest(ctx context.Context, request api.PatchApiWebsiteWebsiteIdRequestObject) (api.PatchApiWebsiteWebsiteIdResponseObject, error) {
	return api.PatchApiWebsiteWebsiteId200JSONResponse(api.Website{
		Id:         request.WebsiteId,
		Campus:     "交大相關",
		Department: "行政單位",
		Office:     "圖書館",
		Name:       "交大圖書館",
		Url:        "https://lib.nctu.edu.tw/",
	}), nil
}
