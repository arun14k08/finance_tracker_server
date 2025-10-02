package serializers

type UserCreateResponse struct {
	ID int64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	CreatedAt int64 `json:"created_at"`
}

type UserLoginResponse struct {
	SignedToken string `json:"token"`
	ExpiresAt int64 `json:"expires_at"`
}
