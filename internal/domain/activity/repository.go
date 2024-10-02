package activity

import (
	"context"
	"database/sql"

	"github.com/ferdiebergado/fullstack-go/internal/db"
)

type ActivityRepository interface {
	Create(ctx context.Context, activityParams db.CreateActivityParams) (*db.Activity, error)
	GetActivities(ctx context.Context) ([]db.ActiveActivity, error)
	FindActiveActivity(ctx context.Context, id int32) (*db.ActiveActivity, error)
	UpdateActivity(ctx context.Context, params db.UpdateActivityParams) error
	DeleteActivity(ctx context.Context, id int32) error
}

type activityRepository struct {
	db      *sql.DB
	queries *db.Queries
}

func NewActivityRepository(conn *sql.DB, queries *db.Queries) ActivityRepository {
	return &activityRepository{
		db:      conn,
		queries: queries,
	}
}

func (r *activityRepository) Create(ctx context.Context, activityParams db.CreateActivityParams) (*db.Activity, error) {
	activity, err := r.queries.CreateActivity(ctx, activityParams)

	if err != nil {
		return nil, err
	}

	return &activity, nil
}

// FindActiveActivity implements ActivityRepository.
func (r *activityRepository) FindActiveActivity(ctx context.Context, id int32) (*db.ActiveActivity, error) {
	activity, err := r.queries.FindActivity(ctx, id)

	if err != nil {
		return nil, err
	}

	return &activity, nil
}

// GetActivities implements ActivityRepository.
func (r *activityRepository) GetActivities(ctx context.Context) ([]db.ActiveActivity, error) {
	activities, err := r.queries.ListActivities(ctx)

	if err != nil {
		return nil, err
	}

	return activities, nil
}

// UpdateActivity implements ActivityRepository.
func (r *activityRepository) UpdateActivity(ctx context.Context, params db.UpdateActivityParams) error {
	return r.queries.UpdateActivity(ctx, params)
}

// DeleteActivity implements ActivityRepository.
func (r *activityRepository) DeleteActivity(ctx context.Context, id int32) error {
	return r.queries.DeleteActivity(ctx, id)
}
