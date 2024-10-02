package test

import (
	"strings"
	"testing"
)

func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
}

func AssertContains(t *testing.T, expected string, actual string) {
	if !strings.Contains(actual, expected) {
		t.Errorf("Expected %s do not contain actual %s", expected, actual)
	}
}
