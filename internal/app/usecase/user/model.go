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

type GetUserByUUIDRequest struct {
	Id string `json:"id"`
}

type GetUserDetailResponse struct {
	User  User      `json:"user"`
	Point UserPoint `json:"point"`
}

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type UserPoint struct {
	Level    int `json:"level"`
	TotalHp  int `json:"total_hp"`
	TotalExp int `json:"total_exp"`
}

type UpdateUserRequest struct {
	UUID      string `json:"uuid"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	UpdatedBy string `json:"updated_by"`
}

type UpdateUserResponse struct {
	UUID      string `json:"uuid"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	UpdatedBy string `json:"updated_by"`
}
