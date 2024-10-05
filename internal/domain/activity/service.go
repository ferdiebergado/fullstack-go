package activity

import (
	"context"
	"database/sql"

	"github.com/ferdiebergado/fullstack-go/internal/db"
	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
	"github.com/ferdiebergado/fullstack-go/pkg/validator"
)

type ActivityService interface {
	CreateActivity(ctx context.Context, req db.CreateActivityParams) (*db.Activity, error)
	ListActivities(ctx context.Context) ([]db.ActiveActivity, error)
	FindActiveActivity(ctx context.Context, id int32) (*db.ActiveActivity, error)
	UpdateActivity(ctx context.Context, params db.UpdateActivityParams) error
	DeleteActivity(ctx context.Context, id int32) error
}

type activityService struct {
	db      *sql.DB
	queries *db.Queries
}

func NewActivityService(database *db.Database) ActivityService {
	return &activityService{db: database.Db, queries: database.Query}
}

func (s *activityService) CreateActivity(ctx context.Context, params db.CreateActivityParams) (*db.Activity, error) {

	validationRules := validator.ValidationRules{
		"title":      "required|min:2|max:150",
		"start_date": "required|date",
		"end_date":   "required|date|after:start_date",
		"venue":      "max:60",
		"host":       "max:60",
	}

	validationErrors := validator.Validate(params, validationRules)

	if len(validationErrors) > 0 {
		return nil, &myhttp.ValidationErrorBag{Message: "Invalid activity", ValidationErrors: validationErrors}
	}

	activityParams := db.CreateActivityParams{
		Title:     params.Title,
		StartDate: params.StartDate,
		EndDate:   params.EndDate,
		Venue:     params.Venue,
		Host:      params.Host,
		Metadata:  params.Metadata,
	}

	activity, err := s.queries.CreateActivity(ctx, activityParams)

	if err != nil {
		return nil, err
	}

	return &activity, nil
}

// FindActiveActivity implements ActivityService.
func (s *activityService) FindActiveActivity(ctx context.Context, id int32) (*db.ActiveActivity, error) {
	activity, err := s.queries.FindActivity(ctx, id)

	if err != nil {
		return nil, err
	}

	return &activity, nil
}

// ListActivities implements ActivityService.
func (s *activityService) ListActivities(ctx context.Context) ([]db.ActiveActivity, error) {
	activities, err := s.queries.ListActivities(ctx)

	if err != nil {
		return nil, err
	}

	return activities, nil
}

// UpdateActivity implements ActivityService.
func (s *activityService) UpdateActivity(ctx context.Context, params db.UpdateActivityParams) error {
	return s.queries.UpdateActivity(ctx, params)
}

// DeleteActivity implements ActivityService.
func (s *activityService) DeleteActivity(ctx context.Context, id int32) error {

	_, err := s.queries.FindActivity(ctx, id)

	if err != nil {
		return err
	}

	return s.queries.DeleteActivity(ctx, id)
}
