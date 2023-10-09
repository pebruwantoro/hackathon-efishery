package entity

type Level struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Point int    `json:"point"`
}

func (l *Level) TableName() string {
	return "levels"
}
