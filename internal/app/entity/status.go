package entity

type Status struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s *Status) TableName() string {
	return "status"
}
