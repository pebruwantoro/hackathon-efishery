package entity

type Status struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (s *Status) TableName() string {
	return "status"
}
