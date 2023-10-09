package entity

type ObjectiveUser struct {
	ObjectiveUUID string `json:"objective_uuid"`
	UserUUID      string `json:"user_uuid"`
}

func (o *ObjectiveUser) TableName() string {
	return "objective_users"
}
