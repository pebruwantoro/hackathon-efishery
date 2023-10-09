package repository

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
	"gorm.io/gorm"
)

type Roles interface {
	GetByID(ctx context.Context, id int) (entity entity.Role, err error)
}

type roles struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) Roles {
	if db == nil {
		panic("database is nil")
	}

	return &roles{db}
}

func (r *roles) GetByID(ctx context.Context, id int) (entity entity.Role, err error) {
	err = r.db.WithContext(ctx).Last(&entity, id).Error
	return
}
