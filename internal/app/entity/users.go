package entity

import "time"

type User struct {
	UUID       string    `json:"uuid"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	AccessRole string    `json:"access_role"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  string    `json:"updated_by"`
	DeletedAt  time.Time `json:"deleted_at"`
	DeletedBy  string    `json:"deleted_by"`
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
