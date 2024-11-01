package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ferdiebergado/fullstack-go/internal/db"
	"github.com/ferdiebergado/fullstack-go/pkg/http/response"
	"github.com/ferdiebergado/fullstack-go/pkg/test"
	router "github.com/ferdiebergado/go-express"
)

const apiEndpoint = "/api/activities"

var (
	payload = &db.CreateActivityParams{
		Title:     "New Activity",
		StartDate: db.NewDate(time.Now()),
		EndDate:   db.NewDate(time.Now().AddDate(0, 0, 1)),
		VenueID:   1,
		HostID:    1,
		Metadata:  json.RawMessage(`{}`),
	}

	conn = db.OpenDb()
)

func setupTestRouter(t *testing.T) *router.Router {
	t.Helper()

	q := db.New(conn)
	database := db.NewDatabase(conn, q)
	router := NewApp(database)

	return router
}

func createActivity(t *testing.T) db.Activity {
	t.Helper()

	conn := db.OpenDb()

	queries := db.New(conn)
	activity, err := queries.CreateActivity(context.Background(), *payload)

	if err != nil {
		log.Fatal(err)
	}

	return activity
}

func TestCreateActivity(t *testing.T) {
	router := setupTestRouter(t)
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest(http.MethodPost, apiEndpoint, bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	test.AssertEqual(t, http.StatusCreated, rr.Code)
}

func TestCreateActivityInvalid(t *testing.T) {
	router := setupTestRouter(t)
	payload.Title = ""
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest(http.MethodPost, apiEndpoint, bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	test.AssertEqual(t, http.StatusBadRequest, rr.Code)

	// Check if the JSON response matches ApiResponse struct
	var response response.ApiResponse[[]db.ActiveActivityDetail]
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	t.Log("response:", response)

	test.AssertLen(t, response.Meta.Errors, 1)
	test.AssertEqual(t, "title", response.Meta.Errors[0].Field)
}

func TestGetActivity(t *testing.T) {
	t.Parallel()

	router := setupTestRouter(t)
	activity := createActivity(t)

	url := fmt.Sprintf("%s/%d", apiEndpoint, activity.ID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	test.AssertEqual(t, http.StatusOK, rr.Code)
}

func TestUpdateActivity(t *testing.T) {
	t.Parallel()

	payload := &db.UpdateActivityParams{
		Title:     "Updated Activity",
		StartDate: db.NewDate(time.Now()),
		EndDate:   db.NewDate(time.Now().AddDate(0, 0, 2)),
		VenueID:   1,
		HostID:    2,
		Metadata:  json.RawMessage(`{}`),
		ID:        1,
	}

	router := setupTestRouter(t)
	body, _ := json.Marshal(payload)

	activity := createActivity(t)

	url := fmt.Sprintf("%s/%d", apiEndpoint, activity.ID)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	test.AssertEqual(t, http.StatusOK, rr.Code)
}

func TestDeleteActivity(t *testing.T) {
	t.Parallel()

	router := setupTestRouter(t)
	activity := createActivity(t)

	url := fmt.Sprintf("%s/%d", apiEndpoint, activity.ID)

	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	test.AssertEqual(t, http.StatusNoContent, rr.Code)
}

func TestListActiveActivities(t *testing.T) {
	t.Parallel()
	router := setupTestRouter(t)

	req, err := http.NewRequest(http.MethodGet, apiEndpoint, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	test.AssertEqual(t, http.StatusOK, rr.Code)
	test.AssertContains(t, rr.Header().Get("Content-Type"), "application/json")
}
