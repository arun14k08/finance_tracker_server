package context

import "github.com/gofiber/fiber/v2"

type AppContext struct {
	fiber *fiber.Ctx
	user User
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