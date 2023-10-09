package entity

type ObjectiveUser struct {
	ID          uint `json:"id"`
	ObjectiveID uint `json:"objective_id"`
	UserID      uint `json:"user_id"`
}

func (o *ObjectiveUser) TableName() string {
	return "objective_users"
}
