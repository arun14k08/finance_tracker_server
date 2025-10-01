package services

import (
	"github.com/arun14k08/finance_tracker_server/pkg/context"
	"github.com/arun14k08/finance_tracker_server/pkg/db"
	"github.com/arun14k08/finance_tracker_server/pkg/handlers/utils"
	"github.com/arun14k08/finance_tracker_server/pkg/models"
	"github.com/arun14k08/finance_tracker_server/pkg/repo"
	"github.com/arun14k08/goframework/framework"
	"github.com/gofiber/fiber/v2"
)
func CreateUser(appCtx context.AppContext) (context.AppContext, bool) {
	user := appCtx.GetUser()

	// Email validation
	if !utils.IsValidEmail(user.Email) {
		framework.BadRequest(appCtx.GetFiberCtx(), "Email ID is invalid")
		return appCtx, false
	}

	// Check if email exists
	userWithEmailID, err := repo.GetUserWithEmail(user.Email)
	if err != nil {
		framework.InternalError(appCtx.GetFiberCtx())
		return appCtx, false
	}
	if userWithEmailID.ID != 0 {
		framework.BadRequest(appCtx.GetFiberCtx(), "Email ID is already taken")
		return appCtx, false
	}

	// Password validation
	if !utils.IsStrongPassword(user.PasswordHash) {
		framework.BadRequest(appCtx.GetFiberCtx(), "Password does not meet the requirements")
		return appCtx, false
	}

	// // Insert user
	// userModel := models.User{
	// 	Name:         user.Name,
	// 	Email:        user.Email,
	// 	PasswordHash: user.PasswordHash,
	// }
	

	// err = repo.CreateUser(userModel)
	if err != nil {
		framework.InternalError(appCtx.GetFiberCtx())
		return appCtx, false
	}

	return appCtx, true
}



func GetUser(userId int64, fiberCtx *fiber.Ctx) context.AppContext {
	user, err := repo.GetUser(userId)
	if err != nil {
		framework.InternalError(fiberCtx)
	}
	return  utils.GetUserContextWithModel(user, fiberCtx)
}
