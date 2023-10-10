package task

import (
	"context"
	"time"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
)

func (u *usecase) Create(ctx context.Context, req CreateTaskRequest) error {
	// GET OBJECTIVE BY ID

	var startDate, endDate, dueDate time.Time
	if req.StartDate != "" {
		startTime, _ := time.Parse(time.RFC3339, req.StartDate)
		startDate = startTime
	}
	if req.EndDate != "" {
		endTime, _ := time.Parse(time.RFC3339, req.EndDate)
		endDate = endTime
	}
	if req.DueDate != "" {
		dueTime, _ := time.Parse(time.RFC3339, req.DueDate)
		dueDate = dueTime
	}
	// MAPPING OBJECT
	task := entity.Task{
		ObjectiveID: req.ObjectiveID,
		Name:        req.Name,
		Description: req.Description,
		Point:       req.Point,
		Status:      req.Status,
		ParentID:    req.ParentId,
		UserID:      req.UserID,
		StartDate:   startDate,
		EndDate:     endDate,
		DueDate:     dueDate,
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
