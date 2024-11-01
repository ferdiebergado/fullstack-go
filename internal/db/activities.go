package db

import (
	"context"
	"fmt"

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
)

const sqlSelect = `-- name: ListActiveActivities :many
SELECT id, title, start_date, end_date, venue_id, host_id, metadata, created_at, updated_at, deleted_at, venue, region, host, COUNT(*) OVER () AS total_items
FROM active_activity_details WHERE COALESCE($1, '') = '' `
const sqlOrder = "ORDER BY %s %s LIMIT $2 OFFSET $3"
const listActiveActivities = sqlSelect + "OR %s ILIKE $1 " + sqlOrder
const listActiveActivitiesByDate = sqlSelect + "OR %s = '%s'::date " + sqlOrder

func (q *Queries) ListActiveActivities(ctx context.Context, arg request.QueryParams, searchFieldType FieldType) ([]ActiveActivityDetailWithCount, error) {

	query := fmt.Sprintf(listActiveActivities, arg.SearchCol, arg.SortCol, arg.SortDir)

	if searchFieldType == DateField {
		query = fmt.Sprintf(listActiveActivitiesByDate, arg.SearchCol, arg.Search, arg.SortCol, arg.SortDir)
	}

	rows, err := q.db.QueryContext(ctx, query,
		arg.Search,
		arg.Limit,
		arg.Offset,
	)

	if err != nil {
		return nil, err
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
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
