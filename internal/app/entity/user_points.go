package entity

type UserPoint struct {
	UUID     string `json:"uuid"`
	UserUUID string `json:"user_uuid"`
	Level    int    `json:"level"`
	TotalHp  int    `json:"total_hp"`
	TotalExp int    `json:"total_exp"`
}

func (u *UserPoint) TableName() string {
	return "user_points"
}
