package utils

import (
	"regexp"

	"github.com/arun14k08/finance_tracker_server/pkg/context"
	"github.com/arun14k08/finance_tracker_server/pkg/db"
	"github.com/arun14k08/finance_tracker_server/pkg/serializers"
	"github.com/arun14k08/goframework/framework"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)


func GetUserContext(userReq *serializers.UserRequest, fiberCtx *fiber.Ctx) context.AppContext{
	var appCtx  context.AppContext
	passwordHash, err := GetPasswordHash(userReq.Password)
	if err != nil {
		
		framework.InternalError(fiberCtx)
	}

	appCtx.SetUser(context.User{
		Name: userReq.Name,
		Email: userReq.Email,
		PassWord: userReq.Password,
		PasswordHash: passwordHash,
		Role: userReq.Role,
	})
	appCtx.SetFiberCtx(fiberCtx)
	return appCtx
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func IsStrongPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[@$!%*?&]`).MatchString(password)

	return hasUpper && hasLower && hasDigit && hasSpecial
}

func GetUserContextWithUser(user db.User, fiberCtx *fiber.Ctx) context.AppContext {
	var appCtx  context.AppContext
	appCtx.SetUser(context.User{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
		CreatedAt: user.CreatedAt.Time.UnixMilli(),
		UpdatedAt: user.UpdatedAt.Time.UnixMilli(),
	})
	appCtx.SetFiberCtx(fiberCtx)
	return appCtx
}

func SetUserContextWithModel(user db.CreateUserRow, appCtx *context.AppContext) {
	appCtx.SetUser(context.User{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
		CreatedAt: user.CreatedAt.Time.UnixMilli(),
		UpdatedAt: user.UpdatedAt.Time.UnixMilli(),
	})
}

func GetPasswordHash(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return "" , err
	}
	return string(passwordHash), nil
}

