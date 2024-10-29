// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"encoding/json"
	"time"
)

type ActiveActivity struct {
	ID        int64           `json:"id"`
	Title     string          `json:"title"`
	StartDate Date            `json:"start_date"`
	EndDate   Date            `json:"end_date"`
	VenueID   int32           `json:"venue_id"`
	HostID    int32           `json:"host_id"`
	Metadata  json.RawMessage `json:"metadata"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt sql.NullTime    `json:"deleted_at"`
}

type ActiveActivityDetail struct {
	ID        int64           `json:"id"`
	Title     string          `json:"title"`
	StartDate Date            `json:"start_date"`
	EndDate   Date            `json:"end_date"`
	Venue     string          `json:"venue"`
	Region    string          `json:"region"`
	Host      string          `json:"host"`
	Metadata  json.RawMessage `json:"metadata"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt sql.NullTime    `json:"deleted_at"`
}

type Activity struct {
	ID        int64           `json:"id"`
	Title     string          `json:"title"`
	StartDate Date            `json:"start_date"`
	EndDate   Date            `json:"end_date"`
	VenueID   int32           `json:"venue_id"`
	HostID    int32           `json:"host_id"`
	Metadata  json.RawMessage `json:"metadata"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt sql.NullTime    `json:"deleted_at"`
}

type ActivityDetail struct {
	ID        int64           `json:"id"`
	Title     string          `json:"title"`
	StartDate Date            `json:"start_date"`
	EndDate   Date            `json:"end_date"`
	Venue     string          `json:"venue"`
	Region    string          `json:"region"`
	Host      string          `json:"host"`
	Metadata  json.RawMessage `json:"metadata"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt sql.NullTime    `json:"deleted_at"`
}

type Division struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	RegionID int16  `json:"region_id"`
}

type Host struct {
	ID        int64           `json:"id"`
	Name      string          `json:"name"`
	Metadata  json.RawMessage `json:"metadata"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt sql.NullTime    `json:"deleted_at"`
}

type Office struct {
	ID        int32           `json:"id"`
	Name      *string         `json:"name"`
	ShortName string          `json:"short_name"`
	Metadata  json.RawMessage `json:"metadata"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt sql.NullTime    `json:"deleted_at"`
}

type Personnel struct {
	ID         int32           `json:"id"`
	Lastname   string          `json:"lastname"`
	Firstname  string          `json:"firstname"`
	Mi         sql.NullString  `json:"mi"`
	PositionID int16           `json:"position_id"`
	OfficeID   sql.NullInt16   `json:"office_id"`
	Metadata   json.RawMessage `json:"metadata"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	DeletedAt  sql.NullTime    `json:"deleted_at"`
}

type Position struct {
	ID        int32           `json:"id"`
	Title     *string         `json:"title"`
	Metadata  json.RawMessage `json:"metadata"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt sql.NullTime    `json:"deleted_at"`
}

type Region struct {
	ID       int16  `json:"id"`
	RegionID int16  `json:"region_id"`
	Name     string `json:"name"`
}

type Travel struct {
	ID         int64           `json:"id"`
	StartDate  Date            `json:"start_date"`
	EndDate    Date            `json:"end_date"`
	Status     int16           `json:"status"`
	Remarks    *string         `json:"remarks"`
	Metadata   json.RawMessage `json:"metadata"`
	ActivityID int32           `json:"activity_id"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	DeletedAt  sql.NullTime    `json:"deleted_at"`
}

type Venue struct {
	ID         int32           `json:"id"`
	Name       string          `json:"name"`
	DivisionID int32           `json:"division_id"`
	Metadata   json.RawMessage `json:"metadata"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	DeletedAt  sql.NullTime    `json:"deleted_at"`
}
