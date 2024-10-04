package activity

import (
	"context"
	"errors"

	"github.com/ferdiebergado/fullstack-go/internal/db"
)

type ActivityService interface {
	CreateActivity(ctx context.Context, req db.CreateActivityParams) (*db.Activity, error)
	ListActivities(ctx context.Context) ([]db.ActiveActivity, error)
	FindActiveActivity(ctx context.Context, id int32) (*db.ActiveActivity, error)
	UpdateActivity(ctx context.Context, params db.UpdateActivityParams) error
	DeleteActivity(ctx context.Context, id int32) error
}

type activityService struct {
	queries *db.Queries
}

func NewActivityService(queries *db.Queries) ActivityService {
	return &activityService{queries: queries}
}

func (s *activityService) CreateActivity(ctx context.Context, params db.CreateActivityParams) (*db.Activity, error) {
	if params.Title == "" {
		return nil, errors.New("title is required")
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
