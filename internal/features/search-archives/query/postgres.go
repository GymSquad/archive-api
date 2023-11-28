package query

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"regexp"

	"github.com/GymSquad/archive-api/api"
	"github.com/GymSquad/archive-api/internal/db"
	searcharchives "github.com/GymSquad/archive-api/internal/features/search-archives"

	_ "embed"
)

//go:embed search-archive.sql
var searchArchivesQuery string

type postgresQuery struct {
	db db.DBTX
}

// NewPostgresQuery creates a new postgres query
func NewPostgresQuery(db db.DBTX) *postgresQuery {
	return &postgresQuery{
		db: db,
	}
}

type postgresSearchParams struct {
	campusID     string
	departmentID string
	officeID     string
	websiteID    string
	query        string
	limit        int
}

type postgresSearchResult struct {
	campusID       string
	campusName     string
	departmentID   string
	departmentName string
	officeID       string
	officeName     string
	websiteID      string
	websiteName    string
	websiteURL     string
	totalResults   int
}

// SearchArchives implements SearchArchivesQuery.
func (p *postgresQuery) SearchArchives(ctx context.Context, query string, cursor *string, limit int) ([]api.SearchResultEntry, api.Pagination, error) {
	c, err := decodeCursor(cursor)
	if err != nil {
		slog.Error("failed to decode cursor", "err", err)
		return nil, api.Pagination{}, errors.Join(searcharchives.ErrorInvalidCursor, err)
	}

	results, err := p.searchArchives(ctx, postgresSearchParams{
		campusID:     c.campusID,
		departmentID: c.departmentID,
		officeID:     c.officeID,
		websiteID:    c.websiteID,
		query:        query,
		limit:        limit,
	})
	if err != nil {
		slog.Error("failed to search archives", "err", err)
		return nil, api.Pagination{}, errors.Join(searcharchives.ErrorInternal, err)
	}

	var pagination api.Pagination
	if len(results) > 0 {
		lastResult := results[len(results)-1]
		c.campusID = lastResult.campusID
		c.departmentID = lastResult.departmentID
		c.officeID = lastResult.officeID
		c.websiteID = lastResult.websiteID
		pagination.NextCursor = c.encode()
		pagination.TotalResults = lastResult.totalResults
	} else {
		pagination.TotalResults = -1
	}
	pagination.NumResults = len(results)

	var entries []api.SearchResultEntry
	for _, r := range results {
		compositeId := fmt.Sprintf("%s$%s$%s", r.campusID, r.departmentID, r.officeID)
		if len(entries) == 0 || entries[len(entries)-1].Id != compositeId {
			entries = append(entries, api.SearchResultEntry{
				Id:         compositeId,
				Campus:     r.campusName,
				Department: r.departmentName,
				Office:     r.officeName,
				Websites:   []api.SearchResultWebsiteEntry{},
			})
		}
		lastEntry := entries[len(entries)-1]
		lastEntry.Websites = append(lastEntry.Websites, api.SearchResultWebsiteEntry{
			Id:   r.websiteID,
			Name: r.websiteName,
			Url:  r.websiteURL,
		})
		entries[len(entries)-1] = lastEntry
	}

	return entries, pagination, nil
}

type postgresCursor struct {
	campusID     string
	departmentID string
	officeID     string
	websiteID    string
}

func defaultCursor() *postgresCursor {
	return &postgresCursor{
		campusID:     "",
		departmentID: "",
		officeID:     "",
		websiteID:    "",
	}
}

func (c *postgresCursor) encode() *string {
	if c == nil {
		return nil
	}
	if c.campusID == "" && c.departmentID == "" && c.officeID == "" && c.websiteID == "" {
		return nil
	}
	s := fmt.Sprintf("(cid=%s,did=%s,oid=%s,wid=%s)", c.campusID, c.departmentID, c.officeID, c.websiteID)
	b64 := base64.StdEncoding.EncodeToString([]byte(s))
	return &b64
}

var cursorRe = regexp.MustCompile(`cid=([^,]+),did=([^,]+),oid=([^,]+),wid=([^,]+)`)

func decodeCursor(cursor *string) (*postgresCursor, error) {
	c := defaultCursor()
	if cursor == nil {
		return c, nil
	}

	b64, err := base64.StdEncoding.DecodeString(*cursor)
	if err != nil {
		return nil, err
	}

	matches := cursorRe.FindStringSubmatch(string(b64))
	if len(matches) != 5 {
		return nil, fmt.Errorf("invalid cursor")
	}

	c.campusID = matches[1]
	c.departmentID = matches[2]
	c.officeID = matches[3]
	c.websiteID = matches[4]
	return c, nil
}

func (p *postgresQuery) searchArchives(ctx context.Context, params postgresSearchParams) ([]postgresSearchResult, error) {
	rows, err := p.db.QueryContext(
		ctx,
		searchArchivesQuery,
		params.campusID,
		params.departmentID,
		params.officeID,
		params.websiteID,
		params.query,
		params.limit,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}

	var results []postgresSearchResult
	for rows.Next() {
		var r postgresSearchResult
		err := rows.Scan(
			&r.campusID,
			&r.campusName,
			&r.departmentID,
			&r.departmentName,
			&r.officeID,
			&r.officeName,
			&r.websiteID,
			&r.websiteName,
			&r.websiteURL,
			&r.totalResults,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, r)
	}

	return results, nil
}
