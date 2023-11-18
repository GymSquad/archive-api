package getarchiveddates

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// GetArchiveDatesHandler is the handler for the get archive dates endpoint
type GetArchiveDatesHandler struct {
	RootPath string
}

// NewHTTPHandler creates a new GetArchiveDatesHandler
func NewHTTPHandler(rootPath string) *GetArchiveDatesHandler {
	return &GetArchiveDatesHandler{
		RootPath: rootPath,
	}
}

func (h *GetArchiveDatesHandler) RegisterRoutes(router gin.IRouter) {
	router.GET("/website/:websiteId", h.Handle)
}

// Handle godoc
//
//	@Summary Get archived dates
//	@Description Get the dates that a website was archived
//	@Tags archive
//	@Produce json
//	@Param websiteId path string true "Website ID"
//	@Success 200 {array} string "List of dates"
//	@Failure 404 {string} string "Not Found"
//	@Router /website/{websiteId} [get]
func (h *GetArchiveDatesHandler) Handle(c *gin.Context) {
	archiveId := c.Param("websiteId")
	entries, err := os.ReadDir(filepath.Join(h.RootPath, archiveId))
	if err != nil {
		c.String(http.StatusNotFound, "Not Found")
	}

	dates := []string{}
	for _, entry := range entries {
		if !entry.IsDir() || !isDate(entry.Name()) {
			continue
		}
		dates = append(dates, entry.Name())
	}

	c.JSON(http.StatusOK, dates)
}
func isDate(date string) bool {
	date = date + "T00:00:00Z"
	_, err := time.Parse(time.RFC3339, date)
	return err == nil
}
