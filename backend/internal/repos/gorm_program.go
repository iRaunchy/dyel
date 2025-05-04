package repos

import (
	"context"

	"github.com/iraunchy/dyel/backend/db"
	"gorm.io/gorm"
)

// GORMProgramRepo implements ProgramRepo using GORM.
type GORMProgramRepo struct {
	DB *gorm.DB
}

// NewGORMProgramRepo wires in a *gorm.DB instance.
func NewGORMProgramRepo(dbConn *gorm.DB) *GORMProgramRepo {
	return &GORMProgramRepo{DB: dbConn}
}

func (r *GORMProgramRepo) Create(ctx context.Context, p *db.Program) (*db.Program, error) {
	if err := r.DB.WithContext(ctx).Create(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (r *GORMProgramRepo) Get(ctx context.Context, id string) (*db.Program, error) {
	var p db.Program
	if err := r.DB.WithContext(ctx).
		Preload("Days.Exercises").
		First(&p, "id = ?", id).
		Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *GORMProgramRepo) List(ctx context.Context) ([]db.Program, error) {
	var list []db.Program
	if err := r.DB.WithContext(ctx).
		Preload("Days.Exercises").
		Find(&list).
		Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *GORMProgramRepo) Update(ctx context.Context, p *db.Program) (*db.Program, error) {
	if err := r.DB.WithContext(ctx).Save(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (r *GORMProgramRepo) Delete(ctx context.Context, id string) error {
	return r.DB.WithContext(ctx).
		Delete(&db.Program{}, "id = ?", id).
		Error
}
