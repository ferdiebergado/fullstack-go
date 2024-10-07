package validator

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ferdiebergado/fullstack-go/internal/db"
	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
	"github.com/ferdiebergado/fullstack-go/pkg/str"
)

type ValidationRules = map[string]string

// GetValueByJSONTagName takes a struct and returns a field's value based on the JSON tag name
func GetValueByJSONTagName[T any](st T, jsonTagName string) (any, bool) {
	v := reflect.ValueOf(st)

	// Ensure we're working with a struct
	if v.Kind() != reflect.Struct {
		return nil, false
	}

	// Iterate through the fields of the struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		jsonTag := field.Tag.Get("json")

		// Check if the json tag matches the provided tag name
		if jsonTag == jsonTagName {
			// Return the value of the corresponding field
			return v.Field(i).Interface(), true
		}
	}
	return nil, false // Return false if no matching tag is found
}

func Validate[T any](params T, validationRules ValidationRules) []myhttp.ValidationError {
	var validationErrors []myhttp.ValidationError

outerLoop:
	for field, rules := range validationRules {
		log.Print("field, rules: ", field, rules)
		for _, rule := range strings.Split(rules, "|") {
			fieldValue, _ := GetValueByJSONTagName(params, field)

			log.Println("value:", fieldValue)

			if fieldValue == nil && rule == "required" {
				validationErrors = append(validationErrors, myhttp.ValidationError{Field: field, Error: "This field is required."})
				continue outerLoop
			}

			// Split rule and possible parameters (like "min:3" -> rule = "min", param = "3")
			parts := strings.Split(rule, ":")
			ruleName := parts[0]
			var param string
			if len(parts) > 1 {
				rule = parts[0]
				param = parts[1]
			}

			strValue := GetStringValue(fieldValue)

			log.Printf("field: %s, value: %s\n", field, strValue)

			switch rule {
			case "required":
				if strings.TrimSpace(strValue) == "" {
					validationErrors = append(validationErrors, myhttp.ValidationError{Field: field, Error: "This field is required."})
					continue outerLoop
				}
			case "alphanumeric":
				re := regexp.MustCompile("^[a-zA-Z0-9 ]*$")

				if !re.MatchString(fmt.Sprint(strValue)) {
					validationErrors = append(validationErrors, myhttp.ValidationError{Field: field, Error: "Must be alphanumeric."})
				}
			case "min":
				if param != "" {
					minLen, err := strconv.Atoi(param)
					if err == nil && len(strValue) < minLen {
						validationErrors = append(validationErrors, myhttp.ValidationError{Field: field, Error: fmt.Sprintf("Must be at least %d characters long.", minLen)})
					}
				}
			case "max":
				if param != "" {
					maxLen, err := strconv.Atoi(param)
					if err == nil && len(strValue) > maxLen {
						validationErrors = append(validationErrors, myhttp.ValidationError{Field: field, Error: fmt.Sprintf("Must be no more than %d characters long.", maxLen)})
					}
				}
			case "email":
				// Basic email regex pattern, can be customized
				emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
				matched := emailRegex.MatchString(strValue)
				if !matched {
					validationErrors = append(validationErrors, myhttp.ValidationError{Field: field, Error: "Must be a valid email address."})
				}
			case "numeric":
				if _, err := strconv.Atoi(strValue); err != nil {
					validationErrors = append(validationErrors, myhttp.ValidationError{Field: field, Error: "Must be a valid number."})
				}
			case "min_num":
				if param != "" {
					minVal, err := strconv.Atoi(param)
					if err == nil {
						numVal, convErr := strconv.Atoi(strValue)
						if convErr == nil && numVal < minVal {
							validationErrors = append(validationErrors, myhttp.ValidationError{Field: field, Error: fmt.Sprintf("Must be greater than or equal to %d.", minVal)})
						}
					}
				}
			case "max_num":
				if param != "" {
					maxVal, err := strconv.Atoi(param)
					if err == nil {
						numVal, convErr := strconv.Atoi(strValue)
						if convErr == nil && numVal > maxVal {
							validationErrors = append(validationErrors, myhttp.ValidationError{Field: field, Error: fmt.Sprintf("Must be less than or equal to %d.", maxVal)})
						}
					}
				}
			case "regex":
				if param != "" {
					matched, _ := regexp.MatchString(param, strValue)
					if !matched {
						validationErrors = append(validationErrors, myhttp.ValidationError{Field: field, Error: "Invalid format"})
					}
				}
			case "date":
				_, err := time.Parse(time.DateOnly, strValue)

				if err != nil {
					validationErrors = append(validationErrors, myhttp.ValidationError{Field: field, Error: "Invalid date"})
				}
			case "after":
				if param != "" {
					// Get the value of the other date field
					otherFieldValue, _ := GetValueByJSONTagName(params, param)
					otherStrValue := GetStringValue(otherFieldValue)

					// Parse the current date and the other date
					currentDate, err := time.Parse(time.DateOnly, strValue) // Assuming it's a string
					if err != nil {
						validationErrors = append(validationErrors, myhttp.ValidationError{Field: field, Error: "Invalid date format"})
					}

					if otherStrValue != "" {
						otherDate, err := time.Parse(time.DateOnly, otherStrValue) // Assuming it's a string
						if err != nil {
							validationErrors = append(validationErrors, myhttp.ValidationError{Field: field, Error: "Invalid date format"})
						}

						// Validate if the current date is on or after the other date
						if !currentDate.Equal(otherDate) && !currentDate.After(otherDate) {
							validationErrors = append(validationErrors, myhttp.ValidationError{Field: field, Error: fmt.Sprintf("Must be on or after %s", str.SnakeToTitle(param))})
						}
					}
				}
			default:
				// TODO: Handle other validation rules if needed
				validationErrors = append(validationErrors, myhttp.ValidationError{Field: field, Error: fmt.Sprintf("Invalid rule: %s", ruleName)})
			}
		}
	}

	return validationErrors
}

func GetStringValue(val any) string {
	var strValue string

	if val == nil {
		return ""
	}

	switch v := val.(type) {
	case db.Date:
		if v.Valid {
			strValue = v.Time.Format(time.DateOnly)
		} else {
			strValue = ""
		}
	case string:
		strValue = v
	case *string:
		if v != nil {
			strValue = *v
		} else {
			strValue = ""
		}
	case int16:
		strValue = strconv.FormatInt(int64(val.(int16)), 10)
	case int32:
		strValue = strconv.FormatInt(int64(val.(int32)), 10)
	case int64:
		strValue = strconv.FormatInt(val.(int64), 10)
	default:
		strValue = val.(string)
	}

	return strValue
}
