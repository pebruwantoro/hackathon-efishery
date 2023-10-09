package entity

type ObjectiveUser struct {
	ObjectiveID uint `json:"objective_id"`
	UserID      uint `json:"user_id"`
}

func (o *ObjectiveUser) TableName() string {
	return "objective_users"
}
