package entity

type UserLevel struct {
	ID       uint `json:"id"`
	UserID   uint `json:"user_id"`
	Level    int  `json:"level"`
	TotalHp  int  `json:"total_hp"`
	TotalExp int  `json:"total_exp"`
}

func (u *UserLevel) TableName() string {
	return "user_levels"
}
