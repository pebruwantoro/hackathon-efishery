package user

type CreateUserRequest struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Role      int    `json:"role"`
	CreatedBy string `json:"created_by"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
