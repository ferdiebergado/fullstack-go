package division

import (
	"context"
	"database/sql"

	"github.com/ferdiebergado/fullstack-go/internal/db"
)

type DivisionService interface {
	GetDivisions(ctx context.Context) ([]db.GetDivisionWithRegionRow, error)
}

type divisionService struct {
	db      *sql.DB
	queries *db.Queries
}

func NewDivisionService(database *db.Database) DivisionService {
	return &divisionService{db: database.Db, queries: database.Query}
}

// GetDivisions implements ActivityService.
func (s *divisionService) GetDivisions(ctx context.Context) ([]db.GetDivisionWithRegionRow, error) {
	return s.queries.GetDivisionWithRegion(ctx)
}
