package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type NullString sql.NullString

// Custom UnmarshalJSON to handle string -> sql.NullString conversion
func (ns *NullString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		ns.String = *s
		ns.Valid = true
	} else {
		ns.Valid = false
	}
	return nil
}

// Custom unmarshal to handle date parsing
func (p *CreateActivityParams) UnmarshalJSON(data []byte) error {
	var temp struct {
		Title     string          `json:"title"`
		StartDate string          `json:"start_date"`
		EndDate   string          `json:"end_date"`
		Venue     sql.NullString  `json:"venue"`
		Host      sql.NullString  `json:"host"`
		Metadata  json.RawMessage `json:"metadata"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	startDate, err := time.Parse(time.DateOnly, temp.StartDate)
	if err != nil {
		return fmt.Errorf("invalid start_date format: %v", err)
	}
	endDate, err := time.Parse(time.DateOnly, temp.EndDate)
	if err != nil {
		return fmt.Errorf("invalid end_date format: %v", err)
	}

	// Assign parsed data to the original struct
	p.Title = temp.Title
	p.StartDate = startDate
	p.EndDate = endDate
	p.Venue = temp.Venue
	p.Host = temp.Host
	p.Metadata = temp.Metadata

	return nil
}
