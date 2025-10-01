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

func CreateUser(fiberCtx *fiber.Ctx) error {
	log.Printf("Hitting /users POST")
	userReq := new(serializers.UserRequest)
	if err := fiberCtx.BodyParser(userReq); err != nil {
		return framework.BadRequest(fiberCtx, "Bad Request")
	}
	appCtx := utils.GetUserContext(userReq, fiberCtx)
	ok := services.CreateUser(&appCtx)
	if !ok {
		return nil
	}
	return framework.Success(appCtx.GetFiberCtx(), serializers.UserCreateResponse{
		ID:        appCtx.GetUser().ID,
		Name:      appCtx.GetUser().Name,
		Email:     appCtx.GetUser().Email,
		Role: appCtx.GetUser().Role,
		CreatedAt: appCtx.GetUser().CreatedAt,
	})
}


func GetUser(fiberCtx *fiber.Ctx) error {
	id := fiberCtx.Query("id")

    userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
        return framework.BadRequest(fiberCtx, "Invalid user ID")
    }
	appCtx, OK := services.GetUser(userID, fiberCtx)
	if !OK {
		return nil
	}

	framework.Success(appCtx.GetFiberCtx(), serializers.UserCreateResponse{
			ID: appCtx.GetUser().ID,
			Name: appCtx.GetUser().Name,
			Email: appCtx.GetUser().Email,
			Role: appCtx.GetUser().Role,
			CreatedAt: appCtx.GetUser().CreatedAt,
		})
	return  nil
}