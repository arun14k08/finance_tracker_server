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

type CreateAccountResponse struct {
	ID int64 `json:"id"`
	UserID int64 `json:"user_id"`
	Name string `json:"name"`
	Balance string `json:"balance"`
	AccountType string `json:"account_type"`
	BankName string `json:"bank_name,omitempty"`
	LastFour string `json:"last_four,omitempty"`
	IsActive bool `json:"is_active"`
	Currency string `json:"currency"`
	NickName string `json:"nick_name,omitempty"`
	Notes string `json:"notes,omitempty"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}
