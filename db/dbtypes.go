package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

func (d *DateOnlyTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d)
}

// Scan implements the sql.Scanner interface for DateOnlyTime
func (d *DateOnlyTime) Scan(value interface{}) error {
	if value == nil {
		*d = DateOnlyTime{} // Handle NULL value by setting a zero value
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*d = DateOnlyTime{Time: v}
		return nil
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type *DateOnlyTime", value)
	}
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

// Scan implements the sql.Scanner interface for NullString
func (ns *NullString) Scan(value interface{}) error {
	if value == nil {
		ns.String, ns.Valid = "", false
		return nil
	}
	switch v := value.(type) {
	case string:
		ns.String, ns.Valid = v, true
		return nil
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type *NullString", value)
	}
}

// StringToNullString converts a string to sql.NullString.
func StringToNullString(s string) NullString {
	if s == "" {
		return NullString{String: "", Valid: false} // Set to NULL
	}
	return NullString{String: s, Valid: true} // Set to the actual string
}
