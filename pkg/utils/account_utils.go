package utils

import (
	"database/sql"

	"github.com/arun14k08/finance_tracker_server/pkg/context"
	"github.com/arun14k08/finance_tracker_server/pkg/db"
	"github.com/arun14k08/goframework/framework"
)

func AddAccountInCtx(ctx *context.AppContext, account db.Account) {
	accounts :=ctx.GetAccounts()
	accounts = append(accounts, context.Account{
		ID: account.ID,
		UserID: account.UserID,
		Name: account.Name,
		Balance: account.Balance, // converting to cents
		AccountType: account.AccountType.String,
		BankName: account.BankName,
		LastFour: account.LastFour,
		IsActive: account.IsActive.Bool,
		NickName: account.Nickname.String,
		Notes: account.Notes.String,
		CreatedAt: account.CreatedAt.Time.Unix(),
		UpdatedAt: account.UpdatedAt.Time.Unix(),
	})
	ctx.SetAccounts(accounts)
}

func GetAccountsForUser(ctx *context.AppContext, userId int64) {
	accounts, err := db.DBConnector.GetAccountsByUserId(ctx.GetFiberCtx().Context(), userId)
	if err != nil && err != sql.ErrNoRows {
		framework.InternalError(ctx.GetFiberCtx())
		return
	}
	var accountModels []context.Account
	for _, account := range accounts {
		accountModels = append(accountModels, context.Account{
			ID: account.ID,
			UserID: account.UserID,
			Name: account.Name,
			Balance: account.Balance, // converting to cents
			AccountType: account.AccountType.String,
			Currency: account.Currency.String,
			BankName: account.BankName,
			LastFour: account.LastFour,
			IsActive: account.IsActive.Bool,
			NickName: account.Nickname.String,
			Notes: account.Notes.String,
			CreatedAt: account.CreatedAt.Time.Unix(),
			UpdatedAt: account.UpdatedAt.Time.Unix(),
		})
	}
	ctx.SetAccounts(accountModels)
}

func UpdateAccount(ctx *context.AppContext, updatedAccount db.Account) {
	accounts := ctx.GetAccounts()
	for i, account := range accounts {
		if account.ID == updatedAccount.ID {
			accounts[i] = context.Account{
				ID: updatedAccount.ID,
				UserID: updatedAccount.UserID,
				Name: updatedAccount.Name,
				Balance: updatedAccount.Balance, // converting to cents
				AccountType: updatedAccount.AccountType.String,
				Currency: updatedAccount.Currency.String,
				BankName: updatedAccount.BankName,
				LastFour: updatedAccount.LastFour,
				IsActive: updatedAccount.IsActive.Bool,
				NickName: updatedAccount.Nickname.String,
				Notes: updatedAccount.Notes.String,
				CreatedAt: updatedAccount.CreatedAt.Time.Unix(),
				UpdatedAt: updatedAccount.UpdatedAt.Time.Unix(),
			}
			break
		}
	}
	ctx.SetAccounts(accounts)
}

func RemoveAccountFromCtx(ctx *context.AppContext, accountID int64) {
	accounts := ctx.GetAccounts()
	for i, account := range accounts {
		if account.ID == accountID {
			accounts = append(accounts[:i], accounts[i+1:]...)
			break
		}
	}
	ctx.SetAccounts(accounts)
}