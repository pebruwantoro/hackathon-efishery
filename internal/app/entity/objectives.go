package entity

import "time"

type Objective struct {
	ID          int       `json:"id"`
	ParentID    int       `json:"parent_id"`
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
}

func (o *Objective) SetCreated(created string) {
	o.CreatedAt = time.Now()
	o.CreatedBy = created
}

func (o *Objective) SetUpdated(updated string) {
	o.UpdatedAt = time.Now()
	o.UpdatedBy = updated
}

func (o *Objective) TableName() string {
	return "objectives"
}
