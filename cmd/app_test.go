package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/ferdiebergado/fullstack-go/internal/db"
	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
	"github.com/ferdiebergado/fullstack-go/pkg/test"
)

var (
	payload = &db.CreateActivityParams{
		Title:     "New Activity",
		StartDate: db.NewDate(time.Now()),
		EndDate:   db.NewDate(time.Now().AddDate(0, 0, 1)),
		Venue:     nil,
		Host:      nil,
		Metadata:  json.RawMessage(`{}`),
	}
)

func setupTestRouter(t *testing.T) *myhttp.Router {
	t.Helper()

	// conn, tx := setupTestDB(t)
	conn := db.OpenDb()

	q := db.New(conn)
	router := NewApp(conn, q)

	return router
}

func createActivity(t *testing.T) db.Activity {
	t.Helper()

	log.Print("creating activity... ")
	// conn, tx := setupTestDB(t)
	conn := db.OpenDb()

	queries := db.New(conn)
	activity, err := queries.CreateActivity(context.Background(), *payload)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("created.")

	return activity
}

func TestCreateActivity(t *testing.T) {
	router := setupTestRouter(t)
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest(http.MethodPost, "/api/activities", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	test.AssertEqual(t, http.StatusCreated, rr.Code)
}

func TestGetActivity(t *testing.T) {
	router := setupTestRouter(t)
	activity := createActivity(t)

	id := strconv.Itoa(int(activity.ID))

	req, err := http.NewRequest(http.MethodGet, "/api/activities/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	test.AssertEqual(t, http.StatusOK, rr.Code)
}

func TestUpdateActivity(t *testing.T) {
	payload := &db.UpdateActivityParams{
		Title:     "Updated Activity",
		StartDate: db.NewDate(time.Now()),
		EndDate:   db.NewDate(time.Now().AddDate(0, 0, 2)),
		Venue:     nil,
		Host:      nil,
		Metadata:  json.RawMessage(`{}`),
		ID:        1,
	}
	router := setupTestRouter(t)
	body, _ := json.Marshal(payload)

	activity := createActivity(t)

	id := strconv.Itoa(int(activity.ID))

	req, err := http.NewRequest(http.MethodPut, "/api/activities/"+id, bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	test.AssertEqual(t, http.StatusOK, rr.Code)
}

func TestDeleteActivity(t *testing.T) {
	router := setupTestRouter(t)
	activity := createActivity(t)

	id := strconv.Itoa(int(activity.ID))

	req, err := http.NewRequest(http.MethodDelete, "/api/activities/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	test.AssertEqual(t, http.StatusNoContent, rr.Code)
}

func TestListActiveActivities(t *testing.T) {
	router := setupTestRouter(t)

	req, err := http.NewRequest(http.MethodGet, "/activities", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Accept", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	test.AssertEqual(t, http.StatusOK, rr.Code)

	test.AssertContains(t, rr.Header().Get("Content-Type"), "application/json")
}
