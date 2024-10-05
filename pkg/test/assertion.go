package test

import (
	"fmt"
	"reflect"
	"testing"

	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
)

// AssertEqual asserts that two values are equal.
func AssertEqual(t *testing.T, expected, actual interface{}, msg ...string) {
	if !reflect.DeepEqual(expected, actual) {
		message := formatMessage("expected", expected, actual, msg...)
		t.Errorf(message)
	}
}

// AssertNotEqual asserts that two values are not equal.
func AssertNotEqual(t *testing.T, expected, actual interface{}, msg ...string) {
	if reflect.DeepEqual(expected, actual) {
		message := formatMessage("not expected", expected, actual, msg...)
		t.Errorf(message)
	}
}

// AssertNoError asserts that an error is nil.
func AssertNoError(t *testing.T, err error, msg ...string) {
	if err != nil {
		message := formatMessage("no error", nil, err, msg...)
		t.Errorf(message)
	}
}

// AssertError asserts that an error is not nil.
func AssertError(t *testing.T, err error, msg ...string) {
	if err == nil {
		message := formatMessage("error", "non-nil error", err, msg...)
		t.Errorf(message)
	}
}

// AssertContains asserts that a string contains a substring.
func AssertContains(t *testing.T, s, substr string, msg ...string) {
	if !contains(s, substr) {
		message := formatMessage(fmt.Sprintf("'%s' to contain", substr), substr, s, msg...)
		t.Errorf(message)
	}
}

// AssertLen asserts that a collection has the expected length.
func AssertLen(t *testing.T, collection interface{}, length int, msg ...string) {
	actualLen := reflect.ValueOf(collection).Len()
	if actualLen != length {
		message := formatMessage("length", length, actualLen, msg...)
		t.Errorf(message)
	}
}

// AssertValidationError asserts that a validation error bag contains specific errors.
func AssertValidationError(t *testing.T, actualErrors []string, expectedErrors []string, msg ...string) {
	AssertEqual(t, expectedErrors, actualErrors, msg...)
}

// AssertApiResponse asserts that an ApiResponse struct contains the expected values.
func AssertApiResponse(t *testing.T, response myhttp.ApiResponse, expectedSuccess bool, expectedData interface{}, expectedErrors []string) {
	AssertEqual(t, expectedSuccess, response.Success, "Success field mismatch")
	AssertEqual(t, expectedData, response.Data, "Data field mismatch")
	AssertEqual(t, expectedErrors, response.Errors, "Errors field mismatch")
}

// Helper function to check if a string contains a substring.
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}

// Helper function to format error messages.
func formatMessage(expectationType string, expected, actual interface{}, msg ...string) string {
	if len(msg) > 0 {
		return msg[0]
	}
	return fmt.Sprintf("Expected %v but got %v for %s", expected, actual, expectationType)
}
