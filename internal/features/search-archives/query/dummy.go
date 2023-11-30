package query

import (
	"context"
	"strconv"

	"github.com/GymSquad/archive-api/api"
)

type dummyQuery struct{}

// NewDummyQuery creates a new dummy query
func NewDummyQuery() *dummyQuery {
	return &dummyQuery{}
}

// SearchArchives implements SearchArchivesQuery.
// This is a dummy query that returns a single result with 10 websites.
func (d *dummyQuery) SearchArchives(ctx context.Context, query string, cursor *string, limit int) ([]api.SearchResultEntry, api.Pagination, error) {
	var result []api.SearchResultEntry
	var pagination api.Pagination

	result = append(result, api.SearchResultEntry{
		Id:         "nctu-administration-library",
		Campus:     "交大相關",
		Department: "行政單位",
		Office:     "圖書館",
		Websites:   []api.Website{},
	})

	for i := 0; i < 10; i++ {
		result[0].Websites = append(result[0].Websites, api.Website{
			Id:   strconv.Itoa(i + 1),
			Name: "交大圖書館",
			Url:  "https://lib.nctu.edu.tw/",
		})
	}

	pagination = api.Pagination{
		NextCursor:   nil,
		NumResults:   10,
		TotalResults: 10,
	}

	return result, pagination, nil
}
