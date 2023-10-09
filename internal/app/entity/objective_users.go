package entity

type ObjectiveUser struct {
	ID          int `json:"id"`
	ObjectiveID int `json:"objective_id"`
	UserID      int `json:"user_id"`
}

func (o *ObjectiveUser) TableName() string {
	return "objective_users"
}
