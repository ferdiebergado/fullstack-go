package db

import (
	"context"
	"database/sql"
	"fmt"
	"slices"
	"time"

	"github.com/ferdiebergado/fullstack-go/pkg/http/request"
)

type FieldType int

type ActiveActivityDetailWithCount struct {
	ActiveActivityDetail
	TotalItems int64
}

const (
	TextField FieldType = iota + 1
	DateField

	sqlSelect = `-- name: ListActiveActivities :many
SELECT id, title, start_date, end_date, venue_id, host_id, metadata, created_at, updated_at, deleted_at, venue, region, host, COUNT(*) OVER () AS total_items
FROM active_activity_details `
	sqlWhereStr  = "WHERE %s ILIKE $1 "
	sqlWhereDate = "WHERE %s = $1 "
	sqlOrder     = "ORDER BY %s %s "
)

func (q *Queries) ListActiveActivities(ctx context.Context, args request.QueryParams, searchFieldType FieldType) ([]ActiveActivityDetailWithCount, error) {

	var rows *sql.Rows
	var err error
	var search any = args.Search
	var queryParams []any = []any{args.Limit, args.Offset}

	sqlStr := sqlSelect + sqlOrder + "LIMIT $1 OFFSET $2"

	query := fmt.Sprintf(sqlStr, args.SortCol, args.SortDir)

	if search != "" {
		var sqlWhere string

		if searchFieldType == TextField {
			sqlWhere = sqlWhereStr
			search = fmt.Sprintf("%%%s%%", search)

		} else if searchFieldType == DateField {
			sqlWhere = sqlWhereDate

			search, err = time.Parse(time.DateOnly, search.(string))

			if err != nil {
				return nil, fmt.Errorf("time parse search: %w", err)
			}
		}

		sqlStr = sqlSelect + sqlWhere + sqlOrder + "LIMIT $2 OFFSET $3"
		query = fmt.Sprintf(sqlStr, args.SearchCol, args.SortCol, args.SortDir)
		queryParams = slices.Insert(queryParams, 0, search)
	}

	rows, err = q.db.QueryContext(ctx, query,
		queryParams...,
	)

	if err != nil {
		return nil, fmt.Errorf("query listactivities: %w", err)
	}

	defer rows.Close()

	var items []ActiveActivityDetailWithCount

	for rows.Next() {
		var i ActiveActivityDetailWithCount
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.StartDate,
			&i.EndDate,
			&i.VenueID,
			&i.HostID,
			&i.Metadata,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Venue,
			&i.Region,
			&i.Host,
			&i.TotalItems,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("rows close: %w", err)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows scan: %w", err)
	}

	return items, nil
}
