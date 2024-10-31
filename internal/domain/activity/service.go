package activity

import (
	"context"
	"database/sql"
	"errors"
	"net/url"
	"strconv"

	"github.com/ferdiebergado/fullstack-go/internal/db"
	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
	"github.com/ferdiebergado/fullstack-go/pkg/validator"
)

const (
	queryParamPage    = "page"
	queryParamLimit   = "limit"
	queryParamSortCol = "sortCol"
	queryParamSortDir = "sortDir"
	queryParamSearch  = "search"
	recordsPerPage    = 5
	sortColumn        = "start_date"
	sortDir           = 1
)

type ActivityService interface {
	CreateActivity(ctx context.Context, req db.CreateActivityParams) (*db.Activity, error)
	ListActivities(ctx context.Context, urlValues url.Values) (*myhttp.PaginatedData[db.ListActiveActivitiesRow], error)
	FindActiveActivity(ctx context.Context, id int64) error
	FindActiveActivityDetails(ctx context.Context, id int64) (*db.ActiveActivityDetail, error)
	UpdateActivity(ctx context.Context, params db.UpdateActivityParams) error
	DeleteActivity(ctx context.Context, id int64) error
	CountActiveActivities(ctx context.Context) (int64, error)
}

type activityService struct {
	db      *sql.DB
	queries *db.Queries
}

var (
	ErrActivityNotFound = errors.New("activity not found")

	activityRules = validator.ValidationRules{
		"title":      "required|min:2|max:300",
		"start_date": "required|date",
		"end_date":   "required|date|after:start_date",
		"venue_id":   "required|numeric",
		"host_id":    "required|numeric",
	}
)

func NewActivityService(database *db.Database) ActivityService {
	return &activityService{db: database.Db, queries: database.Query}
}

func (s *activityService) CreateActivity(ctx context.Context, params db.CreateActivityParams) (*db.Activity, error) {

	v := validator.New(params, activityRules)
	validationErrors := v.Validate()

	if !v.Valid() {
		return nil, &validator.ValidationErrorBag{Message: "Invalid activity", ValidationErrors: validationErrors}
	}

	activityParams := db.CreateActivityParams{
		Title:     params.Title,
		StartDate: params.StartDate,
		EndDate:   params.EndDate,
		VenueID:   params.VenueID,
		HostID:    params.HostID,
		Metadata:  params.Metadata,
	}

	activity, err := s.queries.CreateActivity(ctx, activityParams)

	if err != nil {
		return nil, err
	}

	return &activity, nil
}

// FindActiveActivity implements ActivityService.
func (s *activityService) FindActiveActivity(ctx context.Context, id int64) error {
	_, err := s.queries.FindActiveActivity(ctx, id)

	if err != nil {
		return ErrActivityNotFound
	}

	return nil
}

// FindActiveActivity implements ActivityService.
func (s *activityService) FindActiveActivityDetails(ctx context.Context, id int64) (*db.ActiveActivityDetail, error) {
	activity, err := s.queries.FindActiveActivityDetails(ctx, id)

	if err != nil {
		return nil, err
	}

	return &activity, nil
}

// ListActivities implements ActivityService.
func (s *activityService) ListActivities(ctx context.Context, urlValues url.Values) (*myhttp.PaginatedData[db.ListActiveActivitiesRow], error) {
	page := GetPage(urlValues)
	limit := GetLimit(urlValues)
	offset := (page - 1) * limit

	sortCol := GetSortCol(urlValues)
	sortDir := GetSortDir(urlValues)

	search := GetSearch(urlValues)

	var activities []db.ListActiveActivitiesRow
	var err error
	var totalItems int64

	order := "ASC"

	if sortDir == -1 {
		order = "DESC"
	}

	args := db.ListActiveActivitiesParams{
		Limit:   limit,
		Offset:  offset,
		Column1: &sortCol,
		Column2: &order,
		Column5: &search,
	}

	activities, err = s.queries.ListActiveActivities(ctx, args)

	if err != nil {
		return nil, err
	}

	if len(activities) > 0 {
		totalItems = activities[0].TotalItems
	}

	totalPages := (totalItems + limit - 1) / limit

	paginatedData := &myhttp.PaginatedData[db.ListActiveActivitiesRow]{
		Pagination: &myhttp.PaginationMeta{
			TotalItems: totalItems,
			TotalPages: totalPages,
			Page:       page,
			Limit:      limit,
		},
		Data: activities,
	}

	return paginatedData, nil
}

// UpdateActivity implements ActivityService.
func (s *activityService) UpdateActivity(ctx context.Context, params db.UpdateActivityParams) error {
	v := validator.New(params, activityRules)
	validationErrors := v.Validate()

	if !v.Valid() {
		return &validator.ValidationErrorBag{Message: "Invalid activity", ValidationErrors: validationErrors}
	}

	_, err := s.queries.FindActivity(ctx, params.ID)

	if err != nil {
		return err
	}

	return s.queries.UpdateActivity(ctx, params)
}

// DeleteActivity implements ActivityService.
func (s *activityService) DeleteActivity(ctx context.Context, id int64) error {

	_, err := s.queries.FindActivity(ctx, id)

	if err != nil {
		return err
	}

	return s.queries.DeleteActivity(ctx, id)
}

// CountActiveActivities implements ActivityService.
func (s *activityService) CountActiveActivities(ctx context.Context) (int64, error) {
	return s.queries.CountActiveActivities(ctx)
}

func GetPage(urlValues url.Values) int64 {
	// TODO: Validate query params
	s := urlValues.Get(queryParamPage)

	page, err := strconv.ParseInt(s, 0, 64)

	if err != nil || page < 1 {
		return 1
	}

	return page
}

func GetLimit(urlValues url.Values) int64 {
	// TODO: Validate query params
	s := urlValues.Get(queryParamLimit)

	limit, err := strconv.ParseInt(s, 0, 64)

	if err != nil || limit < 1 {
		return recordsPerPage
	}

	return limit
}

func GetSortCol(urlValues url.Values) string {
	// TODO: Validate query params
	sortCol := urlValues.Get(queryParamSortCol)

	if sortCol == "" {
		return sortColumn
	}

	return sortCol
}

func GetSortDir(urlValues url.Values) int {
	// TODO: Validate query params
	s := urlValues.Get(queryParamSortDir)

	sortDir, err := strconv.Atoi(s)

	if err != nil {
		return sortDir
	}

	if sortDir != 1 && sortDir != -1 {
		return sortDir
	}

	return sortDir
}

func GetSearch(urlValues url.Values) string {
	// TODO: Validate query params
	searchText := urlValues.Get(queryParamSearch)

	if searchText == "" {
		return searchText
	}

	return "%" + searchText + "%"
}
