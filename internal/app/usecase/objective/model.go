package objective

import "time"

type CreateObjectiveRequest struct {
	UserID      int       `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Weight      int       `json:"weight"`
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

type GetObjectiveByUUIDRequest struct {
	ID int `json:"id"`
}

type ObjectiveResponse struct {
	ID             int                  `json:"id"`
	SubobjectiveID int                  `json:"subobjective_id"`
	Name           string               `json:"name"`
	Description    string               `json:"description"`
	Weight         int                  `json:"weight"`
	DueDate        time.Time            `json:"due_date"`
	StartDate      time.Time            `json:"start_date"`
	EndDate        time.Time            `json:"end_date"`
	Subobjective   []*ObjectiveResponse `json:"subobjective"`
}

type UpdateObjectiveRequest struct {
	Id          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Weight      int       `json:"weight"`
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
