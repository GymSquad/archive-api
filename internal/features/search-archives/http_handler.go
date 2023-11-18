package searcharchives

import (
	"database/sql"
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// SearchArchivesHandler is the handler for the search archives endpoint
type SearchArchivesHandler struct {
	db *sql.DB
}

// NewHTTPHandler creates a new SearchArchivesHandler
func NewHTTPHandler(db *sql.DB) *SearchArchivesHandler {
	return &SearchArchivesHandler{
		db: db,
	}
}

func (h *SearchArchivesHandler) RegisterRoutes(router gin.IRouter) {
	router.GET("/search", h.Handle)
}

type SearchArchivesResult struct {
	ArchiveID string `json:"archive_id"`
	WebsiteID string `json:"website_id"`
	Date      string `json:"date"`
}

type SearchArchivesPagination struct {
	NextCursor string `json:"next_cursor"`
}

type SearchArchivesResponse struct {
	Results    []SearchArchivesResult   `json:"results"`
	Pagination SearchArchivesPagination `json:"pagination"`
}

// Handle godoc
//
//		@Summary Search archives
//		@Description Search the archives
//		@Tags archive
//		@Produce json
//		@Param q query string false "Search query"
//	 	@Param cursor query string false "Cursor"
//	 	@Param limit query int false "Limit"
//		@Success 200 {object} SearchArchivesResponse
//		@Failure 400 {string} string "Invalid input"
//		@Failure 500 {string} string "Internal Server Error"
//		@Router /search [get]
func (h *SearchArchivesHandler) Handle(c *gin.Context) {
	q := c.Query("q")
	cursor := c.Query("cursor")
	limit := c.Query("limit")

	if limit == "" {
		limit = "10"
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid limit")
		return
	}

	b, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid cursor")
		return
	}
	cursor = string(b)
	cursor, found := strings.CutPrefix(cursor, "archive:")
	if !found {
		c.String(http.StatusBadRequest, "Invalid cursor")
		return
	}

	results, err := searchArchives(h.db, q, cursor, intLimit)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, SearchArchivesResponse{
		Results: results,
		Pagination: SearchArchivesPagination{
			NextCursor: "2",
		},
	})
}

func searchArchives(*sql.DB, string, string, int) ([]SearchArchivesResult, error) {
	return []SearchArchivesResult{
		{ArchiveID: "1", WebsiteID: "1", Date: "2021-01-01"},
	}, nil
}
