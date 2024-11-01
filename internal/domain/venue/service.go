package venue

import (
	"context"
	"database/sql"

	"github.com/ferdiebergado/fullstack-go/internal/db"
	"github.com/ferdiebergado/fullstack-go/pkg/validator"
)

type VenueService interface {
	GetVenues(ctx context.Context) ([]db.GetVenuesRow, error)
	CreateVenue(ctx context.Context, params db.CreateVenueParams) (*db.Venue, error)
}

type venueService struct {
	db      *sql.DB
	queries *db.Queries
}

var validationRules = validator.ValidationRules{
	"name":        "required",
	"division_id": "required|numeric|min_num:1|max_num:227",
}

func NewVenueService(database *db.Database) VenueService {
	return &venueService{db: database.Db, queries: database.Query}
}

// CreateVenue implements VenueService.
func (s *venueService) CreateVenue(ctx context.Context, params db.CreateVenueParams) (*db.Venue, error) {
	v := validator.NewValidator(params, validationRules)
	validationErrors := v.Validate()

	if !v.Valid() {
		return nil, &validator.ValidationErrorBag{Message: "Invalid venue", ValidationErrors: validationErrors}
	}

	venue, err := s.queries.CreateVenue(ctx, params)

	if err != nil {
		return nil, err
	}

	return &venue, nil
}

// GetVenues implements VenueService.
func (s *venueService) GetVenues(ctx context.Context) ([]db.GetVenuesRow, error) {
	return s.queries.GetVenues(ctx)
}
