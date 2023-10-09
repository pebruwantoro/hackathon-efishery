package task

import "context"

func (u *usecase) Update(ctx context.Context, req UpdateTaskRequest) error {
	task, err := u.taskRepository.GetByID(ctx, uint(req.Id))
	if err != nil {
		return err
	}

	// MAPPING TO UPDATE DATA
	task.Name = req.Name
	task.Description = req.Description
	task.Point = req.Point
	task.Status = req.Status
	task.DueDate = req.DueDate
	task.StartDate = req.StartDate
	task.EndDate = req.EndDate
	task.SetUpdated(req.UpdatedBy)

	err = u.taskRepository.Update(ctx, &task)
	if err != nil {
		return err
	}

	return nil
}
