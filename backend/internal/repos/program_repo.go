package repos

import (
	"context"
	"github.com/iraunchy/dyel/backend/db"
)

type ProgramRepo interface {
	Create(ctx context.Context, p *db.Program) (*db.Program, error)
	Get(ctx context.Context, id string) (*db.Program, error)
	List(ctx context.Context) ([]db.Program, error)
	Update(ctx context.Context, p *db.Program) (*db.Program, error)
	Delete(ctx context.Context, id string) error
}
