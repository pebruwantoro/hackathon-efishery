package repository

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
	"gorm.io/gorm"
)

type ObjectivesUser interface {
	Create(ctx context.Context, entity *entity.ObjectiveUser) (err error)
	Update(ctx context.Context, entity *entity.ObjectiveUser) (err error)
	Delete(ctx context.Context, entity *entity.ObjectiveUser) (err error)
}

type objectiveusers struct {
	db *gorm.DB
}

func NewObjectiveUserRepository(db *gorm.DB) Objectives {
	if db == nil {
		panic("database is nil")
	}

	return &objectives{db}
}

func (r *objectiveusers) Create(ctx context.Context, entity *entity.ObjectiveUser) (err error) {
	err = r.db.WithContext(ctx).Create(entity).Error
	return
}
func (r *objectiveusers) Update(ctx context.Context, entity *entity.ObjectiveUser) (err error) {
	err = r.db.WithContext(ctx).Save(entity).Error
	return
}
func (r *objectiveusers) Delete(ctx context.Context, entity *entity.ObjectiveUser) (err error) {
	err = r.db.WithContext(ctx).Delete(entity).Error
	return
}
