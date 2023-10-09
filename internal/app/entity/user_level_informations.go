package entity

type UserLevelInformation struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	MaxExp int    `json:"max_exp"`
}

func (u *UserLevelInformation) TableName() string {
	return "user_level_informations"
}
