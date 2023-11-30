package getarchiveddates

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/GymSquad/archive-api/api"
	"github.com/GymSquad/archive-api/internal/server"
)

// GetArchiveDatesHandler is the handler for the get archive dates endpoint
type GetArchiveDatesHandler struct {
	RootPath string
}

// Compile time check to ensure that GetArchiveDatesHandler implements server.GetArchivedDatesHandler.
var _ server.GetArchivedDatesHandler = (*GetArchiveDatesHandler)(nil)

// NewHTTPHandler creates a new GetArchiveDatesHandler
func NewHTTPHandler(rootPath string) *GetArchiveDatesHandler {
	return &GetArchiveDatesHandler{
		RootPath: rootPath,
	}
}

// HandleRequest implements server.GetArchivedDatesHandler.
func (h *GetArchiveDatesHandler) HandleRequest(
	ctx context.Context,
	request api.GetArchivedDatesApiWebsiteWebsiteIdGetRequestObject,
) (api.GetArchivedDatesApiWebsiteWebsiteIdGetResponseObject, error) {
	archiveId := request.WebsiteId
	entries, err := os.ReadDir(filepath.Join(h.RootPath, archiveId))
	if err != nil {
		return api.GetArchivedDatesApiWebsiteWebsiteIdGet404JSONResponse("archive for the specified id not found"), nil
	}

	dates := []string{}
	for _, entry := range entries {
		if !entry.IsDir() || !isDate(entry.Name()) {
			continue
		}
		dates = append(dates, entry.Name())
	}

	return api.GetArchivedDatesApiWebsiteWebsiteIdGet200JSONResponse(dates), nil
}

func isDate(date string) bool {
	date = date + "T00:00:00Z"
	_, err := time.Parse(time.RFC3339, date)
	return err == nil
}
