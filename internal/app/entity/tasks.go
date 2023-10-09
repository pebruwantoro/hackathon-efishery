package entity

import "time"

type Task struct {
	UUID          string    `json:"uuid"`
	ObjectiveUUID string    `json:"objective_uuid"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Level         string    `json:"level"`
	Point         string    `json:"point"`
	Status        string    `json:"status"`
	DueDate       time.Time `json:"due_date"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	CreatedAt     time.Time `json:"created_at"`
	CreatedBy     string    `json:"created_by"`
	UpdatedAt     time.Time `json:"updated_at"`
	UpdatedBy     string    `json:"updated_by"`
	DeletedAt     time.Time `json:"deleted_at"`
	DeletedBy     string    `json:"deleted_by"`
}

func (t *Task) TableName() string {
	return "tasks"
}
