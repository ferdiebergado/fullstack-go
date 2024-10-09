// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: divisions.sql

package db

import (
	"context"
)

const findDivision = `-- name: FindDivision :one
SELECT id, name, region_id FROM divisions WHERE id = $1
`

func (q *Queries) FindDivision(ctx context.Context, id int32) (Division, error) {
	row := q.db.QueryRowContext(ctx, findDivision, id)
	var i Division
	err := row.Scan(&i.ID, &i.Name, &i.RegionID)
	return i, err
}

const findDivisionsByRegion = `-- name: FindDivisionsByRegion :many
SELECT id, name, region_id FROM divisions WHERE region_id = $1 ORDER BY name
`

func (q *Queries) FindDivisionsByRegion(ctx context.Context, regionID int16) ([]Division, error) {
	rows, err := q.db.QueryContext(ctx, findDivisionsByRegion, regionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Division
	for rows.Next() {
		var i Division
		if err := rows.Scan(&i.ID, &i.Name, &i.RegionID); err != nil {
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

const getDivisionOrderedByRegion = `-- name: GetDivisionOrderedByRegion :many
SELECT id, name, region_id FROM divisions ORDER BY region_id
`

func (q *Queries) GetDivisionOrderedByRegion(ctx context.Context) ([]Division, error) {
	rows, err := q.db.QueryContext(ctx, getDivisionOrderedByRegion)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Division
	for rows.Next() {
		var i Division
		if err := rows.Scan(&i.ID, &i.Name, &i.RegionID); err != nil {
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

const getDivisionWithRegion = `-- name: GetDivisionWithRegion :many
SELECT d.id, d.name, d.region_id, r.name as region
FROM divisions d
    JOIN regions r ON r.region_id = d.region_id
ORDER BY d.name
`

type GetDivisionWithRegionRow struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	RegionID int16  `json:"region_id"`
	Region   string `json:"region"`
}

func (q *Queries) GetDivisionWithRegion(ctx context.Context) ([]GetDivisionWithRegionRow, error) {
	rows, err := q.db.QueryContext(ctx, getDivisionWithRegion)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetDivisionWithRegionRow
	for rows.Next() {
		var i GetDivisionWithRegionRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.RegionID,
			&i.Region,
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

const getDivisions = `-- name: GetDivisions :many
SELECT id, name, region_id FROM divisions ORDER BY name
`

func (q *Queries) GetDivisions(ctx context.Context) ([]Division, error) {
	rows, err := q.db.QueryContext(ctx, getDivisions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Division
	for rows.Next() {
		var i Division
		if err := rows.Scan(&i.ID, &i.Name, &i.RegionID); err != nil {
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