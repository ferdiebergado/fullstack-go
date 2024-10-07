package db

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Date is a custom type for storing Go's time.Time as a PostgreSQL DATE.
type Date struct {
	Time  time.Time
	Valid bool // Valid is true if the Date is not NULL
}

// Scan implements the sql.Scanner interface to read a DATE value from the database.
func (d *Date) Scan(value interface{}) error {
	if value == nil {
		d.Time, d.Valid = time.Time{}, false
		return nil
	}

	// Convert the value to time.Time
	switch v := value.(type) {
	case time.Time:
		d.Time = v
		d.Valid = true
	default:
		return fmt.Errorf("dbtype.Date: cannot scan type %T into Date", value)
	}

	return nil
}

// Value implements the driver.Valuer interface to write the Date value to the database.
func (d Date) Value() (driver.Value, error) {
	if !d.Valid {
		return nil, nil
	}

	// Return only the date part (truncate time part)
	return d.Time.Format(time.DateOnly), nil
}

// MarshalJSON implements the json.Marshaler interface for JSON serialization.
func (d Date) MarshalJSON() ([]byte, error) {
	if !d.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(d.Time.Format(time.DateOnly))
}

// UnmarshalJSON implements the json.Unmarshaler interface for JSON deserialization.
func (d *Date) UnmarshalJSON(data []byte) error {
	// If the JSON is null, mark the Date as invalid (NULL).
	if string(data) == "null" {
		d.Time, d.Valid = time.Time{}, false
		return nil
	}

	// Try to parse the date string from the JSON.
	var dateString string
	if err := json.Unmarshal(data, &dateString); err != nil {
		return err
	}

	// Parse the string into a time.Time.
	parsedTime, err := time.Parse(time.DateOnly, dateString)
	if err != nil {
		return err
	}

	d.Time, d.Valid = parsedTime, true
	return nil
}

// NewDate creates a new Date instance with a valid time.
func NewDate(t time.Time) Date {
	return Date{Time: t, Valid: true}
}

// NullDate returns a Date instance with a NULL value.
func NullDate() Date {
	return Date{Valid: false}
}

func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10) // Convert the Int64 field to a string
}

func NullInt64ToString(ni sql.NullInt64) string {
	if ni.Valid {
		return strconv.FormatInt(ni.Int64, 10) // Convert the Int64 field to a string
	}
	return "" // Return an empty string if the value is null
}

// Function to get the value of sql.NullInt64
func GetNullInt64Value(ni sql.NullInt64) (int64, error) {
	if ni.Valid {
		return ni.Int64, nil
	}
	return 0, fmt.Errorf("value is null")
}

func NewSqlNullInt64(n int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: n,
		Valid: true,
	}
}
