package context

import "github.com/gofiber/fiber/v2"

type AppContext struct {
	fiber *fiber.Ctx
	user User
	accounts []Account
	request interface{}
	response interface{}
}

type User struct {
	ID int64
	Name string
	Email string
	PassWord string
	PasswordHash string
	Role string
	CreatedAt int64
	UpdatedAt int64
}

type Account struct {
	ID int64
	UserID int64
	Name string
	Balance string
	AccountType string
	Currency string
	BankName string
	LastFour string
	IsActive bool
	NickName string
	Notes string
	CreatedAt int64
	UpdatedAt int64
}

func (ctx *AppContext) GetUser() User {
	return ctx.user
}

func (ctx *AppContext) SetUser(user User)  {
	ctx.user = user
}

func (ctx *AppContext) SetFiberCtx(fiber *fiber.Ctx) {
	ctx.fiber = fiber
}

func (ctx *AppContext) GetFiberCtx() *fiber.Ctx {
	return ctx.fiber 
}

func (ctx *AppContext) GetAccounts() []Account {
	return ctx.accounts
}

func (ctx *AppContext) SetAccounts(accounts []Account) {
	ctx.accounts = accounts
}

func (ctx *AppContext) SetRequest(req interface{}) {
	ctx.request = req
}

func (ctx *AppContext) GetRequest() interface{} {
	return ctx.request
}

func (ctx *AppContext) SetResponse(resp interface{}) {
	ctx.response = resp
}

func (ctx *AppContext) GetResponse() interface{} {
	return ctx.response
}