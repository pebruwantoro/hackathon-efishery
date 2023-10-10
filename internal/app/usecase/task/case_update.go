package task

import (
	"context"
	"fmt"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/entity"
)

func (u *usecase) Update(ctx context.Context, req UpdateTaskRequest) error {
	task, err := u.taskRepository.GetByID(ctx, uint(req.Id))
	if err != nil {
		return err
	}

	if req.Status == "DONE" {
		user, err := u.userLevelRepository.GetByUserID(ctx, req.UserID)
		if err != nil {
			return err
		}

		if user.ID != 0 {
			reqUserLevel := entity.UserLevel{
				ID:              user.ID,
				UserID:          user.UserID,
				HealthPoint:     user.HealthPoint,
				ExperiencePoint: task.Point,
			}
			fmt.Println("reqUserLevel: ", reqUserLevel)
			u.userLevelRepository.Update(ctx, &reqUserLevel)
		}

		fmt.Println("user: ", user)
	}

	// MAPPING TO UPDATE DATA
	task.Name = req.Name
	task.Description = req.Description
	task.Point = req.Point
	task.Status = req.Status
	task.SetUpdated(req.UpdatedBy)

	err = u.taskRepository.Update(ctx, &task)
	if err != nil {
		return err
	}

	return nil
}
