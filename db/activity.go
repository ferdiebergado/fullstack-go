package db

import (
	"encoding/json"
)

type CustomUpdateActivityParams struct {
	Title     string
	StartDate DateOnlyTime
	EndDate   DateOnlyTime
	Venue     NullString
	Host      NullString
	Metadata  json.RawMessage
	ID        int32
}

// Custom unmarshal to handle date parsing
func (u *UpdateActivityParams) UnmarshalJSON(data []byte) error {
	var temp CustomUpdateActivityParams

	return json.Unmarshal(data, &temp)
}

func (u *UpdateActivityParams) MarshalJSON() ([]byte, error) {
	var temp CustomUpdateActivityParams

	return json.Marshal(&temp)
}
