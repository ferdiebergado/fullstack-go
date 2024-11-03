package activity

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"slices"

	"github.com/ferdiebergado/fullstack-go/internal/db"
	"github.com/ferdiebergado/fullstack-go/pkg/http/request"
	"github.com/ferdiebergado/fullstack-go/pkg/http/response"
	"github.com/ferdiebergado/fullstack-go/pkg/validator"
)

type ActivityService interface {
	CreateActivity(ctx context.Context, req db.CreateActivityParams) (*db.Activity, error)
	ListActivities(ctx context.Context, params *request.QueryParams) (*response.PaginatedData[db.ActiveActivityDetailWithCount], error)
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
	ErrUndefinedColumn  = errors.New("column does not exist")

	fields = []string{"title", "start_date", "end_date", "venue", "host"}

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

	v := validator.NewValidator(params, activityRules)
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
func (s *activityService) ListActivities(ctx context.Context, params *request.QueryParams) (*response.PaginatedData[db.ActiveActivityDetailWithCount], error) {

	searchFieldType := db.TextField

	if !slices.Contains(fields, params.SearchCol) {
		params.Search = ""
	}

	if params.SearchCol == "start_date" || params.SearchCol == "end_date" {
		searchFieldType = db.DateField
	}

	activities, err := s.queries.ListActiveActivities(ctx, *params, searchFieldType)

	if err != nil {
		return nil, err
	}

	var totalItems int64

	if len(activities) > 0 {
		totalItems = activities[0].TotalItems
	}

	totalPages := (totalItems + params.Limit - 1) / params.Limit

	paginatedData := &response.PaginatedData[db.ActiveActivityDetailWithCount]{
		Pagination: &response.PaginationMeta{
			TotalItems: totalItems,
			TotalPages: totalPages,
			Page:       params.Page,
			Limit:      params.Limit,
		},
		Data: activities,
	}

	return paginatedData, nil
}

// UpdateActivity implements ActivityService.
func (s *activityService) UpdateActivity(ctx context.Context, params db.UpdateActivityParams) error {
	v := validator.NewValidator(params, activityRules)
	validationErrors := v.Validate()

	if !v.Valid() {
		return &validator.ValidationErrorBag{Message: "Invalid activity.", ValidationErrors: validationErrors}
	}

	_, err := s.queries.FindActivity(ctx, params.ID)

	if err != nil {
		return err
	}

	return s.queries.UpdateActivity(ctx, params)
}

// DeleteActivity implements ActivityService.
func (s *activityService) DeleteActivity(ctx context.Context, id int64) error {

	_, err := s.queries.FindActiveActivity(ctx, id)

	if err != nil {
		// DEBUG:
		log.Println("error on findactivity at deleteactivity service")
		return err
	}

	return s.queries.DeleteActivity(ctx, id)
}

// CountActiveActivities implements ActivityService.
func (s *activityService) CountActiveActivities(ctx context.Context) (int64, error) {
	return s.queries.CountActiveActivities(ctx)
}
