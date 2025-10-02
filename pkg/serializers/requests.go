package serializers

type UserCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"`
}

type UserLoginRequest struct {
	Email string `json:"email"`
	PassWord string `json:"password"`
}