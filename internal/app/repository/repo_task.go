package repository

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
	"gorm.io/gorm"
)

type Tasks interface {
	GetByID(ctx context.Context, id uint) (entity entity.Task, err error)
	GetByUserID(ctx context.Context, userID int) (entities []entity.Task, err error)
	GetAll(ctx context.Context) (entities []entity.Task, err error)
	Create(ctx context.Context, entity *entity.Task) (err error)
	Update(ctx context.Context, entity *entity.Task) (err error)
	Delete(ctx context.Context, entity *entity.Task) (err error)
	GetByObjectiveID(ctx context.Context, objectiveId uint) (entity []entity.Task, err error)
	GetBySubtaskID(ctx context.Context, subtaskId uint) (entity []entity.Task, err error)
}

type tasks struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) Tasks {
	if db == nil {
		panic("database is nil")
	}

	return &tasks{db}
}

func (r *tasks) GetByUserID(ctx context.Context, userID int) (entities []entity.Task, err error) {
	err = r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&entities).Error
	return
}

func (r *tasks) GetByID(ctx context.Context, id uint) (entity entity.Task, err error) {
	err = r.db.WithContext(ctx).Where("id = ?", id).First(&entity).Error
	return
}

func (r *tasks) GetAll(ctx context.Context) (entities []entity.Task, err error) {
	err = r.db.WithContext(ctx).Find(&entities).Error
	return
}

func (r *tasks) Create(ctx context.Context, entity *entity.Task) (err error) {
	err = r.db.WithContext(ctx).Omit("deleted_at", "deleted_by").Create(entity).Error
	return
}

func (r *tasks) Update(ctx context.Context, entity *entity.Task) (err error) {
	err = r.db.WithContext(ctx).Omit("deleted_at", "deleted_by").Save(entity).Error
	return
}

func (r *tasks) Delete(ctx context.Context, entity *entity.Task) (err error) {
	err = r.db.WithContext(ctx).Delete(entity).Error
	return
}

func (r *tasks) GetByObjectiveID(ctx context.Context, objectiveId uint) (entity []entity.Task, err error) {
	err = r.db.WithContext(ctx).Where("objective_id = ?", objectiveId).Find(&entity).Error
	return
}

func (r *tasks) GetBySubtaskID(ctx context.Context, subtaskId uint) (entity []entity.Task, err error) {
	err = r.db.WithContext(ctx).Where("subtask_id = ?", subtaskId).Find(&entity).Error
	return
}
