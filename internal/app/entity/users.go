package entity

import "time"

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	RoleId    uint      `json:"role_id"`
	Password  string    `json:"password"`
	Salary    float64   `json:"salary"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

func (u *User) SetCreated(created string) {
	u.CreatedAt = time.Now()
	u.CreatedBy = created
}

func (u *User) SetUpdated(updated string) {
	u.UpdatedAt = time.Now()
	u.UpdatedBy = updated
}

func (u *User) TableName() string {
	return "users"
}
