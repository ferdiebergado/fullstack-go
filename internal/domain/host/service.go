package host

import (
	"context"
	"database/sql"

	"github.com/ferdiebergado/fullstack-go/internal/db"
)

type HostService interface {
	GetHosts(ctx context.Context) ([]db.Host, error)
	CreateHost(ctx context.Context, name string) (db.Host, error)
}

type hostService struct {
	db      *sql.DB
	queries *db.Queries
}

func NewHostService(database *db.Database) HostService {
	return &hostService{db: database.Db, queries: database.Query}
}

// CreateHost implements HostService.
func (h *hostService) CreateHost(ctx context.Context, name string) (db.Host, error) {
	return h.queries.CreateHost(ctx, name)
}

// GetHosts implements HostService.
func (h *hostService) GetHosts(ctx context.Context) ([]db.Host, error) {
	return h.queries.GetHosts(ctx)
}
