package entity

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (r *Role) TableName() string {
	return "roles"
}
