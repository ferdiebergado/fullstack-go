package region

import (
	"context"
	"database/sql"

	"github.com/ferdiebergado/fullstack-go/internal/db"
)

type RegionService interface {
	GetRegions(ctx context.Context) ([]db.Region, error)
}

type regionService struct {
	db      *sql.DB
	queries *db.Queries
}

func NewRegionService(database *db.Database) RegionService {
	return &regionService{db: database.Db, queries: database.Query}
}

// GetRegions implements RegionService.
func (s *regionService) GetRegions(ctx context.Context) ([]db.Region, error) {
	return s.queries.GetRegions(ctx)
}
