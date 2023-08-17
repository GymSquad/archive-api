package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()
	r.GET("/api/website/:websiteId", getArchiveDates)
	r.Run()

}

func getArchiveDates(c *gin.Context) {
	rootPath := os.Getenv("ROOT_PATH")

	archiveId := c.Param("websiteId")
	entries, err := os.ReadDir(filepath.Join(rootPath, archiveId))

	if err != nil {
		c.String(http.StatusNotFound, "Not Found")
	}

	dates := []string{}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		dates = append(dates, entry.Name())
	}

	c.JSON(http.StatusOK, dates)
}
