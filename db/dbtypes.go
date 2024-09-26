package db

import (
	"database/sql"
	"encoding/json"
	"time"
)

type NullString sql.NullString

type DateOnlyTime struct {
	time.Time
}

func (d *DateOnlyTime) UnmarshalJSON(b []byte) error {
	// Trim quotes around the JSON string
	t, err := time.Parse(time.DateOnly, string(b))
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

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

func (ns *NullString) MarshalJSON() ([]byte, error) {
	return json.Marshal(ns)
}

// StringToNullString converts a string to sql.NullString.
func StringToNullString(s string) NullString {
	if s == "" {
		return NullString{String: "", Valid: false} // Set to NULL
	}
	return NullString{String: s, Valid: true} // Set to the actual string
}
