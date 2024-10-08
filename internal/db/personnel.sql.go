// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: personnel.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"
)

const createPersonnel = `-- name: CreatePersonnel :one
INSERT INTO
    personnel (
        lastname,
        firstname,
        mi,
        position_id,
        office_id,
        metadata
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    id, lastname, firstname, mi, position_id, office_id, metadata, created_at, updated_at, deleted_at
`

type CreatePersonnelParams struct {
	Lastname   string          `json:"lastname"`
	Firstname  string          `json:"firstname"`
	Mi         sql.NullString  `json:"mi"`
	PositionID int16           `json:"position_id"`
	OfficeID   sql.NullInt16   `json:"office_id"`
	Metadata   json.RawMessage `json:"metadata"`
}

func (q *Queries) CreatePersonnel(ctx context.Context, arg CreatePersonnelParams) (Personnel, error) {
	row := q.db.QueryRowContext(ctx, createPersonnel,
		arg.Lastname,
		arg.Firstname,
		arg.Mi,
		arg.PositionID,
		arg.OfficeID,
		arg.Metadata,
	)
	var i Personnel
	err := row.Scan(
		&i.ID,
		&i.Lastname,
		&i.Firstname,
		&i.Mi,
		&i.PositionID,
		&i.OfficeID,
		&i.Metadata,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deletePersonnel = `-- name: DeletePersonnel :exec
UPDATE personnel SET deleted_at = NOW() WHERE id = $1
`

func (q *Queries) DeletePersonnel(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deletePersonnel, id)
	return err
}

const findPersonnel = `-- name: FindPersonnel :one
SELECT id, lastname, firstname, mi, position_id, office_id, metadata, created_at, updated_at, deleted_at FROM personnel WHERE id = $1
`

func (q *Queries) FindPersonnel(ctx context.Context, id int32) (Personnel, error) {
	row := q.db.QueryRowContext(ctx, findPersonnel, id)
	var i Personnel
	err := row.Scan(
		&i.ID,
		&i.Lastname,
		&i.Firstname,
		&i.Mi,
		&i.PositionID,
		&i.OfficeID,
		&i.Metadata,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const findPersonnelByFirstname = `-- name: FindPersonnelByFirstname :many
SELECT id, lastname, firstname, mi, position_id, office_id, metadata, created_at, updated_at, deleted_at FROM personnel WHERE firstname = $1
`

func (q *Queries) FindPersonnelByFirstname(ctx context.Context, firstname string) ([]Personnel, error) {
	rows, err := q.db.QueryContext(ctx, findPersonnelByFirstname, firstname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Personnel
	for rows.Next() {
		var i Personnel
		if err := rows.Scan(
			&i.ID,
			&i.Lastname,
			&i.Firstname,
			&i.Mi,
			&i.PositionID,
			&i.OfficeID,
			&i.Metadata,
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

const findPersonnelByLastname = `-- name: FindPersonnelByLastname :many
SELECT id, lastname, firstname, mi, position_id, office_id, metadata, created_at, updated_at, deleted_at FROM personnel WHERE lastname = $1
`

func (q *Queries) FindPersonnelByLastname(ctx context.Context, lastname string) ([]Personnel, error) {
	rows, err := q.db.QueryContext(ctx, findPersonnelByLastname, lastname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Personnel
	for rows.Next() {
		var i Personnel
		if err := rows.Scan(
			&i.ID,
			&i.Lastname,
			&i.Firstname,
			&i.Mi,
			&i.PositionID,
			&i.OfficeID,
			&i.Metadata,
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

const listPersonnel = `-- name: ListPersonnel :many
SELECT id, lastname, firstname, mi, position_id, office_id, metadata, created_at, updated_at, deleted_at FROM personnel ORDER BY lastname
`

func (q *Queries) ListPersonnel(ctx context.Context) ([]Personnel, error) {
	rows, err := q.db.QueryContext(ctx, listPersonnel)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Personnel
	for rows.Next() {
		var i Personnel
		if err := rows.Scan(
			&i.ID,
			&i.Lastname,
			&i.Firstname,
			&i.Mi,
			&i.PositionID,
			&i.OfficeID,
			&i.Metadata,
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

const restorePersonnel = `-- name: RestorePersonnel :exec
UPDATE personnel SET deleted_at = NULL WHERE id = $1
`

func (q *Queries) RestorePersonnel(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, restorePersonnel, id)
	return err
}

const updatePersonnel = `-- name: UpdatePersonnel :exec
UPDATE personnel
SET
    lastname = $1,
    firstname = $2,
    mi = $3,
    position_id = $4,
    office_id = $5,
    metadata = $6,
    updated_at = NOW()
WHERE
    id = $7
`

type UpdatePersonnelParams struct {
	Lastname   string          `json:"lastname"`
	Firstname  string          `json:"firstname"`
	Mi         sql.NullString  `json:"mi"`
	PositionID int16           `json:"position_id"`
	OfficeID   sql.NullInt16   `json:"office_id"`
	Metadata   json.RawMessage `json:"metadata"`
	ID         int32           `json:"id"`
}

func (q *Queries) UpdatePersonnel(ctx context.Context, arg UpdatePersonnelParams) error {
	_, err := q.db.ExecContext(ctx, updatePersonnel,
		arg.Lastname,
		arg.Firstname,
		arg.Mi,
		arg.PositionID,
		arg.OfficeID,
		arg.Metadata,
		arg.ID,
	)
	return err
}
