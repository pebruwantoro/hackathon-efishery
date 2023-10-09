package repository

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
	"gorm.io/gorm"
)

type UserLevels interface {
	GetByUserUUID(ctx context.Context, uuid string) (entity entity.UserLevel, err error)
	Create(ctx context.Context, entity *entity.UserLevel) (err error)
	Update(ctx context.Context, entity *entity.UserLevel) (err error)
	Delete(ctx context.Context, entity *entity.UserLevel) (err error)
}

type userlevels struct {
	db *gorm.DB
}

func NewUserPointRepository(db *gorm.DB) UserLevels {
	if db == nil {
		panic("database is nil")
	}

	return &userlevels{db}
}

func (r *userlevels) GetByUserUUID(ctx context.Context, uuid string) (entity entity.UserLevel, err error) {
	err = r.db.WithContext(ctx).Where("user_uuid = ?", uuid).First(&entity).Error
	return
}

func (r *userlevels) GetAll(ctx context.Context) (entities []entity.UserLevel, err error) {
	err = r.db.WithContext(ctx).Find(&entities).Error
	return
}

func (r *userlevels) Create(ctx context.Context, entity *entity.UserLevel) (err error) {
	err = r.db.WithContext(ctx).Create(entity).Error
	return
}

func (r *userlevels) Update(ctx context.Context, entity *entity.UserLevel) (err error) {
	err = r.db.WithContext(ctx).Save(entity).Error
	return
}

func (r *userlevels) Delete(ctx context.Context, entity *entity.UserLevel) (err error) {
	err = r.db.WithContext(ctx).Delete(entity).Error
	return
}
