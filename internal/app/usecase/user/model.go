package user

type CreateUserRequest struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Role      int    `json:"role"`
	CreatedBy string `json:"created_by"`
}

type CreateUserResponse struct{}
