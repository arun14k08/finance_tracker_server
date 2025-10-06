package services

import (
	"database/sql"
	"log"

	"github.com/arun14k08/finance_tracker_server/pkg/context"
	"github.com/arun14k08/finance_tracker_server/pkg/db"
	"github.com/arun14k08/finance_tracker_server/pkg/utils"
	"github.com/arun14k08/goframework/framework"
	"github.com/gofiber/fiber/v2"
)
func CreateUser(appCtx *context.AppContext) (ok bool) {
	user := appCtx.GetUser()
	// Email validation
	if !utils.IsValidEmail(user.Email) {
		framework.BadRequest(appCtx.GetFiberCtx(), "Email ID is invalid")
		return false
	}

	userWithEmailID, err := db.DBConnector.GetUserByEmail(appCtx.GetFiberCtx().Context(), appCtx.GetUser().Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// ✅ No user found -> safe to proceed
		} else {
			// ❌ Real DB error
			log.Printf("DB error while fetching user by email: %v", err)
			framework.InternalError(appCtx.GetFiberCtx())
			return false
		}
	} else {
		// ✅ A user was found -> email already taken
		if userWithEmailID.ID != 0 {
			framework.BadRequest(appCtx.GetFiberCtx(), "Email ID is already taken")
			return false
		}
	}

	// Password validation
	if !utils.IsStrongPassword(user.PasswordHash) {
		framework.BadRequest(appCtx.GetFiberCtx(), "Password does not meet the requirements")
		return false
	}
	generatedHash, err := utils.GetPasswordHash(user.PassWord)

	if err != nil {
		framework.InternalError(appCtx.GetFiberCtx())
		return false
	}
	user.PasswordHash = generatedHash
	role := "customer"
	if user.Role != "" {
		role = user.Role
	}

	userRow, err := db.DBConnector.CreateUser(appCtx.GetFiberCtx().Context(), db.CreateUserParams{
		Name: user.Name,
		Email: user.Email,
		PasswordHash: user.PasswordHash,
		Role: role,
	})

	if err != nil {
		framework.InternalError(appCtx.GetFiberCtx())
		return false
	}
	utils.SetUserContextWithModel(userRow, appCtx)
	accountRow, err := db.DBConnector.CreateAccount(appCtx.GetFiberCtx().Context(), db.CreateAccountParams{
		Name:  "Cash",
		UserID: appCtx.GetUser().ID,
		AccountType: sql.NullString{String: "cash", Valid: true},
		Currency: sql.NullString{String: "INR", Valid: true},
		BankName: "",
		LastFour: "0000",
		Nickname: sql.NullString{String: "Cash Account", Valid: true},
		Notes: sql.NullString{String: "Initial deposit", Valid: true},
		Balance: "0",
	})
	if err != nil {
		log.Default().Println("Error creating account:", err.Error())
		framework.InternalError(appCtx.GetFiberCtx())
		return false
	}
	log.Default().Println(accountRow.ID)
	return true
}



func GetUser(userId int64, fiberCtx *fiber.Ctx) (*context.AppContext, bool) {
	userRow, err := db.DBConnector.GetUserById(fiberCtx.Context(), userId)
	if err != nil {
		framework.InternalError(fiberCtx)
	}
	if userRow.ID == 0 {
		framework.BadRequest(fiberCtx, "User ID is invalid")
		return &context.AppContext{}, false
	}
	return  utils.GetUserContextWithUser(userRow, fiberCtx), true
}
