package repository

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
	"gorm.io/gorm"
)

type UserPoints interface {
	GetByUserUUID(ctx context.Context, uuid string) (entity entity.UserPoint, err error)
	Create(ctx context.Context, entity *entity.UserPoint) (err error)
	Update(ctx context.Context, entity *entity.UserPoint) (err error)
	Delete(ctx context.Context, entity *entity.UserPoint) (err error)
}

type userpoints struct {
	db *gorm.DB
}

func NewUserPointRepository(db *gorm.DB) UserPoints {
	if db == nil {
		panic("database is nil")
	}

	return &userpoints{db}
}

func (r *userpoints) GetByUserUUID(ctx context.Context, uuid string) (entity entity.UserPoint, err error) {
	err = r.db.WithContext(ctx).Where("user_uuid = ?", uuid).First(&entity).Error
	return
}

func (r *userpoints) GetAll(ctx context.Context) (entities []entity.UserPoint, err error) {
	err = r.db.WithContext(ctx).Find(&entities).Error
	return
}

func (r *userpoints) Create(ctx context.Context, entity *entity.UserPoint) (err error) {
	err = r.db.WithContext(ctx).Create(entity).Error
	return
}

func (r *userpoints) Update(ctx context.Context, entity *entity.UserPoint) (err error) {
	err = r.db.WithContext(ctx).Save(entity).Error
	return
}

func (r *userpoints) Delete(ctx context.Context, entity *entity.UserPoint) (err error) {
	err = r.db.WithContext(ctx).Delete(entity).Error
	return
}
