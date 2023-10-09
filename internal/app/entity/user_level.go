package entity

type UserLevel struct {
	ID              uint `json:"id"`
	UserID          uint `json:"user_id"`
	LevelID         uint `json:"level"`
	HealthPoint     int  `json:"health_point"`
	ExperiencePoint int  `json:"experience_point"`
}

func (u *UserLevel) TableName() string {
	return "user_levels"
}
