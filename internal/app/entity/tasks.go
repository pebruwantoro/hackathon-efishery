package entity

import "time"

type Task struct {
	ID          int       `json:"id"`
	ObjectiveID int       `json:"objective_id"`
	SubtaskID   int       `json:"subtask_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Point       int       `json:"point"`
	UserID      int       `json:"user_id"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedAt   time.Time `json:"deleted_at"`
	DeletedBy   string    `json:"deleted_by"`
}

func (t *Task) SetCreated(created string) {
	t.CreatedAt = time.Now()
	t.CreatedBy = created
}

func (t *Task) SetUpdated(updated string) {
	t.UpdatedAt = time.Now()
	t.UpdatedBy = updated
}

func (t *Task) TableName() string {
	return "tasks"
}
