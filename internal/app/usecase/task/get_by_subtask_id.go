package task

import "context"

func (u *usecase) GetBySubtaskID(ctx context.Context, req GetBySubtaskIdRequest) (resp []GetTaskDetailResponse, err error) {
	tasks, err := u.taskRepository.GetBySubtaskID(ctx, uint(req.SubtaskId))
	if err != nil {
		return resp, err
	}

	for _, task := range tasks {
		data := GetTaskDetailResponse{
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

		resp = append(resp, data)
	}

	return resp, nil
}
