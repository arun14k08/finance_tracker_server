package handlers

import (
	"log"
	"strconv"

	"github.com/arun14k08/finance_tracker_server/pkg/serializers"
	"github.com/arun14k08/finance_tracker_server/pkg/services"
	"github.com/arun14k08/finance_tracker_server/pkg/utils"
	"github.com/arun14k08/goframework/framework"
	"github.com/gofiber/fiber/v2"
)

func CreateAccount(fiberCtx *fiber.Ctx) error{
	req := new(serializers.CreateAccountRequest)
    if err := fiberCtx.BodyParser(req); err != nil {
        return framework.BadRequest(fiberCtx, "Invalid request body")
    }
	appCtx, ok := utils.GetCurrentUserContext(fiberCtx)
	if !ok {
		return framework.UnAuthorized(fiberCtx, "Invalid Credentials")
	}
	appCtx.SetRequest(req)
    ok = services.CreateAccount(appCtx)
	if !ok {
		log.Print("Account creation failed")
		return nil
	}

	return framework.Created(appCtx.GetFiberCtx(), appCtx.GetResponse())
}

func GetAccounts(fiberCtx *fiber.Ctx) error {
	appCtx, ok := utils.GetCurrentUserContext(fiberCtx)
	if !ok {
		return framework.UnAuthorized(fiberCtx, "Invalid Credentials")
	}
	ok = services.GetAccounts(appCtx)
	if !ok {
		log.Print("Fetching accounts failed")
		return nil
	}
	
	return framework.Success(appCtx.GetFiberCtx(), appCtx.GetResponse())
}

func GetAccountByID(fiberCtx *fiber.Ctx) error {
	idStr := fiberCtx.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return framework.BadRequest(fiberCtx, "Invalid account ID")
	}
	appCtx, ok := utils.GetCurrentUserContext(fiberCtx)
	if !ok {
		return framework.UnAuthorized(fiberCtx, "Invalid Credentials")
	}
	ok = services.GetAccountByID(appCtx, id)
	if !ok {
		log.Print("Fetching account failed")
		return nil
	}

	return framework.Success(appCtx.GetFiberCtx(), appCtx.GetResponse())
}

func UpdateAccount(fiberCtx *fiber.Ctx) error {
	req := new(serializers.UpdateAccountRequest)
	if err := fiberCtx.BodyParser(req); err != nil {
		return framework.BadRequest(fiberCtx, "Invalid request body")
	}
	appCtx, ok := utils.GetCurrentUserContext(fiberCtx)
	if !ok {
		return framework.UnAuthorized(fiberCtx, "Invalid Credentials")
	}
	appCtx.SetRequest(req)
	ok = services.UpdateAccount(appCtx)
	if !ok {
		log.Print("Updating account failed")
		return nil
	}

	return framework.Success(appCtx.GetFiberCtx(), appCtx.GetResponse())
}

func DeleteAccount(fiberCtx *fiber.Ctx) error {
	idStr := fiberCtx.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return framework.BadRequest(fiberCtx, "Invalid account ID")
	}
	appCtx, ok := utils.GetCurrentUserContext(fiberCtx)
	if !ok {
		return framework.UnAuthorized(fiberCtx, "Invalid Credentials")
	}
	ok = services.DeleteAccount(appCtx, id)
	if !ok {
		log.Print("Deleting account failed")
		return nil
	}

	return framework.Success(appCtx.GetFiberCtx(), appCtx.GetResponse())
}