// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: travels.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

const createTravel = `-- name: CreateTravel :one
INSERT INTO
    travels (
        start_date,
        end_date,
        status,
        remarks,
        metadata,
        activity_id
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    id, start_date, end_date, status, remarks, metadata, activity_id, created_at, updated_at, deleted_at
`

type CreateTravelParams struct {
	StartDate  Date            `json:"start_date"`
	EndDate    Date            `json:"end_date"`
	Status     int16           `json:"status"`
	Remarks    *string         `json:"remarks"`
	Metadata   json.RawMessage `json:"metadata"`
	ActivityID int32           `json:"activity_id"`
}

func (q *Queries) CreateTravel(ctx context.Context, arg CreateTravelParams) (Travel, error) {
	row := q.db.QueryRowContext(ctx, createTravel,
		arg.StartDate,
		arg.EndDate,
		arg.Status,
		arg.Remarks,
		arg.Metadata,
		arg.ActivityID,
	)
	var i Travel
	err := row.Scan(
		&i.ID,
		&i.StartDate,
		&i.EndDate,
		&i.Status,
		&i.Remarks,
		&i.Metadata,
		&i.ActivityID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteTravel = `-- name: DeleteTravel :exec
UPDATE travels SET deleted_at = NOW() WHERE id = $1
`

func (q *Queries) DeleteTravel(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTravel, id)
	return err
}

const findTravel = `-- name: FindTravel :one
SELECT id, start_date, end_date, status, remarks, metadata, activity_id, created_at, updated_at, deleted_at FROM travels WHERE id = $1
`

func (q *Queries) FindTravel(ctx context.Context, id int32) (Travel, error) {
	row := q.db.QueryRowContext(ctx, findTravel, id)
	var i Travel
	err := row.Scan(
		&i.ID,
		&i.StartDate,
		&i.EndDate,
		&i.Status,
		&i.Remarks,
		&i.Metadata,
		&i.ActivityID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const findTravelByActivityId = `-- name: FindTravelByActivityId :many
SELECT id, start_date, end_date, status, remarks, metadata, activity_id, created_at, updated_at, deleted_at FROM travels WHERE activity_id = $1
`

func (q *Queries) FindTravelByActivityId(ctx context.Context, activityID int32) ([]Travel, error) {
	rows, err := q.db.QueryContext(ctx, findTravelByActivityId, activityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Travel
	for rows.Next() {
		var i Travel
		if err := rows.Scan(
			&i.ID,
			&i.StartDate,
			&i.EndDate,
			&i.Status,
			&i.Remarks,
			&i.Metadata,
			&i.ActivityID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
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

const findTravelByActivityTitle = `-- name: FindTravelByActivityTitle :many
SELECT t.id, t.start_date, t.end_date, status, remarks, t.metadata, activity_id, t.created_at, t.updated_at, t.deleted_at, a.id, title, a.start_date, a.end_date, venue, host, a.metadata, a.created_at, a.updated_at, a.deleted_at
FROM travels AS t
    INNER JOIN activities AS a ON t.activity_id = a.id
WHERE
    t.activity_id = $1
ORDER BY t.start_date DESC
`

type FindTravelByActivityTitleRow struct {
	ID          int32           `json:"id"`
	StartDate   Date            `json:"start_date"`
	EndDate     Date            `json:"end_date"`
	Status      int16           `json:"status"`
	Remarks     *string         `json:"remarks"`
	Metadata    json.RawMessage `json:"metadata"`
	ActivityID  int32           `json:"activity_id"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   sql.NullTime    `json:"deleted_at"`
	ID_2        int32           `json:"id_2"`
	Title       string          `json:"title"`
	StartDate_2 Date            `json:"start_date_2"`
	EndDate_2   Date            `json:"end_date_2"`
	Venue       *string         `json:"venue"`
	Host        *string         `json:"host"`
	Metadata_2  json.RawMessage `json:"metadata_2"`
	CreatedAt_2 time.Time       `json:"created_at_2"`
	UpdatedAt_2 time.Time       `json:"updated_at_2"`
	DeletedAt_2 sql.NullTime    `json:"deleted_at_2"`
}

func (q *Queries) FindTravelByActivityTitle(ctx context.Context, activityID int32) ([]FindTravelByActivityTitleRow, error) {
	rows, err := q.db.QueryContext(ctx, findTravelByActivityTitle, activityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindTravelByActivityTitleRow
	for rows.Next() {
		var i FindTravelByActivityTitleRow
		if err := rows.Scan(
			&i.ID,
			&i.StartDate,
			&i.EndDate,
			&i.Status,
			&i.Remarks,
			&i.Metadata,
			&i.ActivityID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.ID_2,
			&i.Title,
			&i.StartDate_2,
			&i.EndDate_2,
			&i.Venue,
			&i.Host,
			&i.Metadata_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.DeletedAt_2,
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

const findTravelByStartDate = `-- name: FindTravelByStartDate :many
SELECT id, start_date, end_date, status, remarks, metadata, activity_id, created_at, updated_at, deleted_at FROM travels WHERE start_date = $1
`

func (q *Queries) FindTravelByStartDate(ctx context.Context, startDate Date) ([]Travel, error) {
	rows, err := q.db.QueryContext(ctx, findTravelByStartDate, startDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Travel
	for rows.Next() {
		var i Travel
		if err := rows.Scan(
			&i.ID,
			&i.StartDate,
			&i.EndDate,
			&i.Status,
			&i.Remarks,
			&i.Metadata,
			&i.ActivityID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
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

const listTravels = `-- name: ListTravels :many
SELECT id, start_date, end_date, status, remarks, metadata, activity_id, created_at, updated_at, deleted_at FROM travels ORDER BY start_date DESC
`

func (q *Queries) ListTravels(ctx context.Context) ([]Travel, error) {
	rows, err := q.db.QueryContext(ctx, listTravels)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Travel
	for rows.Next() {
		var i Travel
		if err := rows.Scan(
			&i.ID,
			&i.StartDate,
			&i.EndDate,
			&i.Status,
			&i.Remarks,
			&i.Metadata,
			&i.ActivityID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
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

const restoreTravel = `-- name: RestoreTravel :exec
UPDATE travels SET deleted_at = NULL WHERE id = $1
`

func (q *Queries) RestoreTravel(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, restoreTravel, id)
	return err
}

const updateTravel = `-- name: UpdateTravel :exec
UPDATE travels
SET
    start_date = $1,
    end_date = $2,
    status = $3,
    remarks = $4,
    activity_id = $5,
    metadata = $6,
    updated_at = NOW()
WHERE
    id = $7
`

type UpdateTravelParams struct {
	StartDate  Date            `json:"start_date"`
	EndDate    Date            `json:"end_date"`
	Status     int16           `json:"status"`
	Remarks    *string         `json:"remarks"`
	ActivityID int32           `json:"activity_id"`
	Metadata   json.RawMessage `json:"metadata"`
	ID         int32           `json:"id"`
}

func (q *Queries) UpdateTravel(ctx context.Context, arg UpdateTravelParams) error {
	_, err := q.db.ExecContext(ctx, updateTravel,
		arg.StartDate,
		arg.EndDate,
		arg.Status,
		arg.Remarks,
		arg.ActivityID,
		arg.Metadata,
		arg.ID,
	)
	return err
}
