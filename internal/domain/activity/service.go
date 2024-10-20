package activity

import (
	"context"
	"database/sql"

	"github.com/ferdiebergado/fullstack-go/internal/db"
	"github.com/ferdiebergado/fullstack-go/pkg/validator"
)

type ActivityService interface {
	CreateActivity(ctx context.Context, req db.CreateActivityParams) (*db.Activity, error)
	ListActivities(ctx context.Context, args db.ListActivitiesParams) ([]db.ListActivitiesRow, error)
	ListActivitiesOrderedDesc(ctx context.Context, args db.ListActivitiesOrderedDescParams) ([]db.ListActivitiesOrderedDescRow, error)
	FindActiveActivity(ctx context.Context, id int64) (*db.FindActivityRow, error)
	UpdateActivity(ctx context.Context, params db.UpdateActivityParams) error
	DeleteActivity(ctx context.Context, id int64) error
	GetRegions(ctx context.Context) ([]db.Region, error)
	GetDivisions(ctx context.Context) ([]db.GetDivisionWithRegionRow, error)
	GetHosts(ctx context.Context) ([]db.Host, error)
	CountActivities(ctx context.Context) (int64, error)
}

type activityService struct {
	db      *sql.DB
	queries *db.Queries
}

// CountActivities implements ActivityService.
func (s *activityService) CountActivities(ctx context.Context) (int64, error) {
	return s.queries.CountActiveActivities(ctx)
}

var (
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
func (s *activityService) FindActiveActivity(ctx context.Context, id int64) (*db.FindActivityRow, error) {
	activity, err := s.queries.FindActivity(ctx, id)

	if err != nil {
		return nil, err
	}

	return &activity, nil
}

// ListActivities implements ActivityService.
func (s *activityService) ListActivities(ctx context.Context, args db.ListActivitiesParams) ([]db.ListActivitiesRow, error) {
	activities, err := s.queries.ListActivities(ctx, args)

	if err != nil {
		return nil, err
	}

	return activities, nil
}

// ListActivitiesOrderedDesc implements ActivityService.
func (s *activityService) ListActivitiesOrderedDesc(ctx context.Context, args db.ListActivitiesOrderedDescParams) ([]db.ListActivitiesOrderedDescRow, error) {
	activities, err := s.queries.ListActivitiesOrderedDesc(ctx, args)

	if err != nil {
		return nil, err
	}

	return activities, nil
}

// UpdateActivity implements ActivityService.
func (s *activityService) UpdateActivity(ctx context.Context, params db.UpdateActivityParams) error {
	v := validator.New(params, activityRules)
	validationErrors := v.Validate()

	if !v.Valid() {
		return &validator.ValidationErrorBag{Message: "Invalid activity", ValidationErrors: validationErrors}
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

// GetRegions implements ActivityService.
func (s *activityService) GetRegions(ctx context.Context) ([]db.Region, error) {
	return s.queries.GetRegions(ctx)
}

// GetDivisions implements ActivityService.
func (s *activityService) GetDivisions(ctx context.Context) ([]db.GetDivisionWithRegionRow, error) {
	return s.queries.GetDivisionWithRegion(ctx)
}

// GetHosts implements ActivityService.
func (s *activityService) GetHosts(ctx context.Context) ([]db.Host, error) {
	return s.queries.GetHosts(ctx)
}
