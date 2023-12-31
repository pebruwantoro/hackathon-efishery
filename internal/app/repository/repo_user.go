package repository

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
	"gorm.io/gorm"
)

type Users interface {
	GetByID(ctx context.Context, id int) (entity entity.User, err error)
	GetByEmail(ctx context.Context, email string) (entity entity.User, err error)
	GetByUsername(ctx context.Context, username string) (entity entity.User, err error)
	GetAll(ctx context.Context) (entities []entity.User, err error)
	Create(ctx context.Context, entity *entity.User) (err error)
	Update(ctx context.Context, entity *entity.User) (err error)
	Delete(ctx context.Context, entity *entity.User) (err error)
}

type users struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Users {
	if db == nil {
		panic("database is nil")
	}

	return &users{db}
}

func (r *users) GetByID(ctx context.Context, id int) (entity entity.User, err error) {
	err = r.db.WithContext(ctx).Where("id = ?", id).First(&entity).Error
	return
}

func (r *users) GetAll(ctx context.Context) (entities []entity.User, err error) {
	err = r.db.WithContext(ctx).Find(&entities).Error
	return
}

func (r *users) Create(ctx context.Context, entity *entity.User) (err error) {
	err = r.db.WithContext(ctx).Omit("deleted_at", "deleted_by").Create(entity).Error
	return
}

func (r *users) Update(ctx context.Context, entity *entity.User) (err error) {
	err = r.db.WithContext(ctx).Omit("deleted_at", "deleted_by").Save(entity).Error
	return
}

func (r *users) Delete(ctx context.Context, entity *entity.User) (err error) {
	err = r.db.WithContext(ctx).Delete(entity).Error
	return
}

func (r *users) GetByEmail(ctx context.Context, email string) (entity entity.User, err error) {
	err = r.db.WithContext(ctx).Where("email = ?", email).First(&entity).Error
	return
}

func (r *users) GetByUsername(ctx context.Context, username string) (entity entity.User, err error) {
	err = r.db.WithContext(ctx).Where("username = ?", username).First(&entity).Error
	return
}
