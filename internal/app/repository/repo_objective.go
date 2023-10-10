package repository

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
	"gorm.io/gorm"
)

type Objectives interface {
	GetSoloObjectiveByUserID(ctx context.Context, UserID int) (entities []entity.UserLevel, err error)
	GetPartyObjectiveByUserID(ctx context.Context, UserID int) (entities []entity.UserLevel, err error)
	GetUserObjectiveByUserID(ctx context.Context, UserID int) (entities []entity.ObjectiveUser, err error)
	GetObjectiveByID(ctx context.Context, IDs []int) (entities []entity.Objective, err error)
	GetObjectiveByName(ctx context.Context, name string) (entity entity.Objective, err error)
	Create(ctx context.Context, entity *entity.Objective) (err error)
	Update(ctx context.Context, entity *entity.Objective) (err error)
	Delete(ctx context.Context, entity *entity.Objective) (err error)
}

type objectives struct {
	db *gorm.DB
}

func NewObjectiveRepository(db *gorm.DB) Objectives {
	if db == nil {
		panic("database is nil")
	}

	return &objectives{db}
}

func (r *objectives) GetUserObjectiveByUserID(ctx context.Context, UserID int) (entities []entity.ObjectiveUser, err error) {
	err = r.db.WithContext(ctx).Where("user_id = ?", UserID).Find(&entities).Error
	return
}

func (r *objectives) GetObjectiveByID(ctx context.Context, IDs []int) (entities []entity.Objective, err error) {
	err = r.db.WithContext(ctx).Where("id IN (?)", IDs).Find(&entities).Error
	return
}

func (r *objectives) GetSoloObjectiveByUserID(ctx context.Context, UserID int) (entities []entity.UserLevel, err error) {
	err = r.db.WithContext(ctx).Where("user_id = ?", UserID).Find(&entities).Error
	return
}

func (r *objectives) GetPartyObjectiveByUserID(ctx context.Context, UserID int) (entities []entity.UserLevel, err error) {
	err = r.db.WithContext(ctx).Where("user_id = ?", UserID).Find(&entities).Error
	return
}

func (r *objectives) Create(ctx context.Context, entity *entity.Objective) (err error) {
	err = r.db.WithContext(ctx).Omit("deleted_at", "deleted_by").Create(entity).Error
	return
}
func (r *objectives) Update(ctx context.Context, entity *entity.Objective) (err error) {
	err = r.db.WithContext(ctx).Omit("deleted_at", "deleted_by").Save(entity).Error
	return
}
func (r *objectives) Delete(ctx context.Context, entity *entity.Objective) (err error) {
	err = r.db.WithContext(ctx).Delete(entity).Error
	return
}

func (r *objectives) GetObjectiveByName(ctx context.Context, name string) (entity entity.Objective, err error) {
	err = r.db.WithContext(ctx).Where("name = ?", name).Last(entity).Error
	return
}
