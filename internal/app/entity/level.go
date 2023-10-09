package entity

type Level struct {
	ID     uint `json:"id"`
	Level  uint `json:"level"`
	MaxExp int  `json:"max_exp"`
}

func (u *Level) TableName() string {
	return "levels"
}
