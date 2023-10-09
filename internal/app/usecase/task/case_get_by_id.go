package task

import "context"

func (u *usecase) GetByID(ctx context.Context, req GetByIdRequest) (resp GetTaskDetailResponse, err error) {
	task, err := u.taskRepository.GetByID(ctx, uint(req.Id))
	if err != nil {
		return resp, err
	}

	resp = GetTaskDetailResponse{
		ID:          task.ID,
		ObjectiveID: task.ObjectiveID,
		Name:        task.Name,
		Description: task.Description,
		Point:       task.Point,
		Status:      task.Status,
		DueDate:     task.DueDate,
		StartDate:   task.StartDate,
		EndDate:     task.EndDate,
		CreatedAt:   task.CreatedAt,
		CreatedBy:   task.CreatedBy,
		UpdatedAt:   task.UpdatedAt,
		UpdatedBy:   task.UpdatedBy,
	}

	return resp, nil
}
