package main

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	// RootPath is the path to the root of the archive``
	RootPath = "/archive"
)

func main() {
	r := gin.Default()
	r.GET("/api/website/:websiteId", getArchiveDates(RootPath))

	slog.Info("Starting server at", "url", "http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		slog.Error("Failed to start server", "err", err)
	}
}

func isDate(date string) bool {
	date = date + "T00:00:00Z"
	_, err := time.Parse(time.RFC3339, date)
	return err == nil
}

func getArchiveDates(rootPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		archiveId := c.Param("websiteId")
		entries, err := os.ReadDir(filepath.Join(rootPath, archiveId))

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
}
