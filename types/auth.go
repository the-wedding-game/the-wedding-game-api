package types

type UserRole string

const (
	Admin  UserRole = "ADMIN"
	Player UserRole = "PLAYER"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	Username string   `json:"username"`
	Role     UserRole `json:"role"`
}

type LoginResponse struct {
	User        UserResponse `json:"user"`
	AccessToken string       `json:"access_token"`
}
