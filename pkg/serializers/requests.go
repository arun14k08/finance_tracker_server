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

type CreateAccountRequest struct {
	Name string `json:"name"`
	AccountType string `json:"account_type"`
	Currency    string `json:"currency"`
	BankName    string `json:"bank_name"`
	LastFour    string `json:"last_four"`
	Balance     float64 `json:"balance"`
	NickName    string `json:"nick_name,omitempty"`
	Notes       string `json:"notes,omitempty"`
	IsActive    bool   `json:"is_active,omitempty"`
}

type UpdateAccountRequest struct {
	ID int64 `json:"id"`
	Name string `json:"name,omitempty"`
	AccountType string `json:"account_type,omitempty"`
	Currency    string `json:"currency,omitempty"`
	BankName    string `json:"bank_name,omitempty"`
	LastFour    string `json:"last_four,omitempty"`
	NickName    string `json:"nick_name,omitempty"`
	Notes       string `json:"notes,omitempty"`
	IsActive    bool  `json:"is_active,omitempty"`
}