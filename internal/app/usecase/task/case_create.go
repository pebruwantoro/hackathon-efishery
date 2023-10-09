package task

import (
	"context"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
)

func (u *usecase) Create(ctx context.Context, req CreateTaskRequest) error {
	// GET OBJECTIVE BY ID

	// MAPPING OBJECT
	task := entity.Task{
		ObjectiveID: uint(req.ObjectiveID),
		Name:        req.Name,
		Description: req.Description,
		Point:       req.Point,
		Status:      req.Status,
	}
	task.SetCreated(req.CreatedBy)
	task.SetUpdated(req.CreatedBy)

	// CREATE TASK
	err := u.taskRepository.Create(ctx, &task)
	if err != nil {
		return err
	}

	return nil
}
