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

type GetUserByIDRequest struct {
	Id int `json:"id"`
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
	Level           int `json:"level"`
	HealthPoint     int `json:"health_point"`
	ExperiencePoint int `json:"experience_point"`
}

type UpdateUserRequest struct {
	ID        int    `json:"uuid"`
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
