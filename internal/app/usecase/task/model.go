package task

import "time"

type CreateTaskRequest struct {
	ObjectiveID int    `json:"objective_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Point       int    `json:"point"`
	Status      string `json:"status"`
	CreatedBy   string `json:"created_by"`
}

type GetByIdRequest struct {
	Id int `json:"id"`
}

type GetTaskDetailResponse struct {
	ID          uint      `json:"id"`
	ObjectiveID uint      `json:"objective_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Point       int       `json:"point"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
}

type UpdateTaskRequest struct {
	Id          int    `json:"id"`
	ObjectiveID int    `json:"objective_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Point       int    `json:"point"`
	Status      string `json:"status"`
	UpdatedBy   string `json:"updated_by"`
}

type GetByObjectiveIdRequest struct {
	ObjectiveId int `json:"objective_id"`
}

type GetBySubtaskIdRequest struct {
	SubtaskId int `json:"subtask_id"`
}
