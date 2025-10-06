package services

import (
	"database/sql"
	"log"

	"github.com/arun14k08/finance_tracker_server/pkg/context"
	"github.com/arun14k08/finance_tracker_server/pkg/db"
	"github.com/arun14k08/finance_tracker_server/pkg/serializers"
	"github.com/arun14k08/finance_tracker_server/pkg/utils"
	"github.com/arun14k08/goframework/framework"
)


func CreateAccount(appCtx *context.AppContext) (ok bool) {
	currentReq := appCtx.GetRequest().(*serializers.CreateAccountRequest)
	// validate required fields
	utils.GetAccountsForUser(appCtx, appCtx.GetUser().ID)
	// check for duplicate account based on bank_name and last_four
	ALreadyExists, err := db.DBConnector.GetAccountByBankAndLastFour(appCtx.GetFiberCtx().Context(), db.GetAccountByBankAndLastFourParams{
		BankName: currentReq.BankName,
		LastFour: currentReq.LastFour,
	})
	if err != nil && err != sql.ErrNoRows {
		log.Println("Error checking for duplicate account:", err)
		framework.InternalError(appCtx.GetFiberCtx())
		return false
	}
	if ALreadyExists.ID != 0 {
		framework.BadRequest(appCtx.GetFiberCtx(), "Account with the same bank name and last four digits already exists")
		return false
	}

	// basic validation

	if currentReq == nil || currentReq.Name == "" || currentReq.AccountType == "" || currentReq.Currency == "" {
		log.Println("Missing required account fields")
		framework.BadRequest(appCtx.GetFiberCtx(), "Missing required account fields")
		return false
	}

	sanitizedBalance := utils.FormatAmount(currentReq.Balance)

	accountModel, err := db.DBConnector.CreateAccount(appCtx.GetFiberCtx().Context(), db.CreateAccountParams{
		Name:  currentReq.Name,
		UserID: appCtx.GetUser().ID,
		AccountType: sql.NullString{String: currentReq.AccountType, Valid: currentReq.AccountType != ""},
		Currency: sql.NullString{String: currentReq.Currency, Valid: currentReq.Currency != ""},
		BankName: currentReq.BankName,
		LastFour: currentReq.LastFour,
		Nickname: sql.NullString{String: currentReq.NickName, Valid: currentReq.NickName != ""},
		Notes: sql.NullString{String: currentReq.Notes, Valid: currentReq.Notes != ""},
		Balance: sanitizedBalance,
	})

	// need to add the created account to the context
	if err != nil {
		log.Println("Error creating account:", err)
		framework.InternalError(appCtx.GetFiberCtx())
		return false
	}
	utils.AddAccountInCtx(appCtx, accountModel)
	// set response in context
	appCtx.SetResponse(serializers.CreateAccountResponse{
		ID: accountModel.ID,
		UserID: accountModel.UserID,
		Name: accountModel.Name,
		Balance: accountModel.Balance,
		AccountType: accountModel.AccountType.String,
		BankName: accountModel.BankName,
		Currency: accountModel.Currency.String,
		LastFour: accountModel.LastFour,
		IsActive: accountModel.IsActive.Bool,
		NickName: accountModel.Nickname.String,
		Notes: accountModel.Notes.String,
		CreatedAt: accountModel.CreatedAt.Time.Unix(),
		UpdatedAt: accountModel.UpdatedAt.Time.Unix(),
	})
	return true
}

func GetAccounts(appCtx *context.AppContext) (ok bool) {
	utils.GetAccountsForUser(appCtx, appCtx.GetUser().ID)
	accounts := appCtx.GetAccounts()
	var response []serializers.CreateAccountResponse
	for _, account := range accounts {
		response = append(response, serializers.CreateAccountResponse{
			ID: account.ID,
			UserID: account.UserID,
			Name: account.Name,
			Balance: account.Balance,
			AccountType: account.AccountType,
			BankName: account.BankName,
			Currency: account.Currency,
			LastFour: account.LastFour,
			IsActive: account.IsActive,
			NickName: account.NickName,
			Notes: account.Notes,
			IsDefaultAccount: utils.IsDefaultAccount(appCtx, account.ID),
			CreatedAt: account.CreatedAt,
			UpdatedAt: account.UpdatedAt,
		})
	}
	appCtx.SetResponse(response)
	return true
}

func GetAccountByID(appCtx *context.AppContext, accountID int64) (ok bool) {
	account, err := db.DBConnector.GetAccountByID(appCtx.GetFiberCtx().Context(), accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			framework.NotFound(appCtx.GetFiberCtx(), "Account")
			return false
		}
		log.Println("Error fetching account by ID:", err)
		framework.InternalError(appCtx.GetFiberCtx())
		return false
	}
	if account.UserID != appCtx.GetUser().ID {
		framework.NotFound(appCtx.GetFiberCtx(), "Account")
		return false
	}
	appCtx.SetResponse(serializers.CreateAccountResponse{
		ID: account.ID,
		UserID: account.UserID,
		Name: account.Name,
		Balance: account.Balance,
		AccountType: account.AccountType.String,
		BankName: account.BankName,
		Currency: account.Currency.String,
		LastFour: account.LastFour,
		IsActive: account.IsActive.Bool,
		IsDefaultAccount: utils.IsDefaultAccount(appCtx, account.ID),
		NickName: account.Nickname.String,
		Notes: account.Notes.String,
		CreatedAt: account.CreatedAt.Time.Unix(),
		UpdatedAt: account.UpdatedAt.Time.Unix(),
	})
	return true
}

func UpdateAccount(appCtx *context.AppContext) (ok bool) {
	currentReq := appCtx.GetRequest().(*serializers.UpdateAccountRequest)
	if currentReq == nil || currentReq.ID == 0 {
		framework.BadRequest(appCtx.GetFiberCtx(), "Missing account ID")
		return false
	}
	existingAccount, err := db.DBConnector.GetAccountByID(appCtx.GetFiberCtx().Context(), currentReq.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			framework.NotFound(appCtx.GetFiberCtx(), "Account")
			return false
		}
		log.Println("Error fetching account by ID:", err)
		framework.InternalError(appCtx.GetFiberCtx())
		return false
	}
	if existingAccount.UserID != appCtx.GetUser().ID {
		framework.NotFound(appCtx.GetFiberCtx(), "Account")
		return false
	}

	if utils.IsDefaultAccountByTypeAndLastFour(existingAccount.AccountType.String, existingAccount.LastFour) {
		framework.BadRequest(appCtx.GetFiberCtx(), "Default account cannot be updated")
		return false
	}

	// Update fields if they are provided in the request
	if currentReq.Name != "" {
		existingAccount.Name = currentReq.Name
	}
	if currentReq.AccountType != "" {
		existingAccount.AccountType = sql.NullString{String: currentReq.AccountType, Valid: true}
	}
	if currentReq.Currency != "" {
		existingAccount.Currency = sql.NullString{String: currentReq.Currency, Valid: true}
	}
	if currentReq.BankName != "" {
		existingAccount.BankName = currentReq.BankName
	}
	if currentReq.LastFour != "" {
		existingAccount.LastFour = currentReq.LastFour
	}
	if currentReq.NickName != "" {
		existingAccount.Nickname = sql.NullString{String: currentReq.NickName, Valid: true}
	}
	if currentReq.Notes != "" {
		existingAccount.Notes = sql.NullString{String: currentReq.Notes, Valid: true}
	}
	if currentReq.IsActive != existingAccount.IsActive.Bool {
		existingAccount.IsActive = sql.NullBool{Bool: currentReq.IsActive, Valid: true}
	}

	// check for duplicate account based on bank_name and last_four
	ALreadyExists, err := db.DBConnector.GetAccountByBankAndLastFour(appCtx.GetFiberCtx().Context(), db.GetAccountByBankAndLastFourParams{
		BankName: currentReq.BankName,
		LastFour: currentReq.LastFour,
		UserID: appCtx.GetUser().ID,
	})
	if err != nil && err != sql.ErrNoRows {
		log.Println("Error checking for duplicate account:", err)
		framework.InternalError(appCtx.GetFiberCtx())
		return false
	}
	if ALreadyExists.ID != currentReq.ID {
		framework.BadRequest(appCtx.GetFiberCtx(), "Account with the same bank name and last four digits already exists")
		return false
	}

	updatedAccount, err := db.DBConnector.UpdateAccount(appCtx.GetFiberCtx().Context(), db.UpdateAccountParams{
		ID: existingAccount.ID,
		Name: existingAccount.Name,
		AccountType: existingAccount.AccountType,
		Currency: existingAccount.Currency,
		BankName: existingAccount.BankName,
		LastFour: existingAccount.LastFour,
		Nickname: existingAccount.Nickname,
		Notes: existingAccount.Notes,
		IsActive: existingAccount.IsActive,
		IsDefaultAccount: utils.IsDefaultAccount(appCtx, existingAccount.ID),
	})
	if err != nil {
		log.Println("Error updating account:", err)
		framework.InternalError(appCtx.GetFiberCtx())
		return false
	}
	utils.UpdateAccount(appCtx, updatedAccount)
	appCtx.SetResponse(serializers.CreateAccountResponse{
		ID: updatedAccount.ID,
		UserID: updatedAccount.UserID,
		Name: updatedAccount.Name,
		Balance: updatedAccount.Balance,
		AccountType: updatedAccount.AccountType.String,
		BankName: updatedAccount.BankName,
		Currency: updatedAccount.Currency.String,
		LastFour: updatedAccount.LastFour,
		IsActive: updatedAccount.IsActive.Bool,
		NickName: updatedAccount.Nickname.String,
		Notes: updatedAccount.Notes.String,
		CreatedAt: updatedAccount.CreatedAt.Time.Unix(),
		UpdatedAt: updatedAccount.UpdatedAt.Time.Unix(),
	})
	return true
}

func DeleteAccount(appCtx *context.AppContext, accountID int64) (ok bool) {
	account, err := db.DBConnector.GetAccountByID(appCtx.GetFiberCtx().Context(), accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			framework.NotFound(appCtx.GetFiberCtx(), "Account")
			return false
		}
		log.Println("Error fetching account by ID:", err)
		framework.InternalError(appCtx.GetFiberCtx())
		return false
	}
	if account.UserID != appCtx.GetUser().ID {
		framework.NotFound(appCtx.GetFiberCtx(), "Account")
		return false
	}
	err = db.DBConnector.DeleteAccount(appCtx.GetFiberCtx().Context(), accountID)
	if err != nil {
		log.Println("Error deleting account:", err)
		framework.InternalError(appCtx.GetFiberCtx())
		return false
	}
	utils.RemoveAccountFromCtx(appCtx, accountID)
	appCtx.SetResponse(struct {
		Message string `json:"message"`
	}{
		Message: "Account deleted successfully",
	})
	return true
}