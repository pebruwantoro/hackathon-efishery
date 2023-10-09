package task

import "context"

func (u *usecase) GetByObjectiveID(ctx context.Context, req GetByObjectiveIdRequest) (resp []GetTaskDetailResponse, err error) {
	tasks, err := u.taskRepository.GetByObjectiveID(ctx, uint(req.ObjectiveId))
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
