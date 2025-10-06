package handlers

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/arun14k08/finance_tracker_server/pkg/db"
	"github.com/arun14k08/finance_tracker_server/pkg/serializers"
	"github.com/arun14k08/finance_tracker_server/pkg/services"
	"github.com/arun14k08/finance_tracker_server/pkg/utils"
	"github.com/arun14k08/goframework/config"
	"github.com/arun14k08/goframework/framework"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateUser(fiberCtx *fiber.Ctx) error {
	userReq := new(serializers.UserCreateRequest)
	if err := fiberCtx.BodyParser(userReq); err != nil {
		return framework.BadRequest(fiberCtx, "Bad Request")
	}
	appCtx := utils.GetUserContext(userReq, fiberCtx)
	ok := services.CreateUser(appCtx)
	if !ok {
		return nil
	}
	return framework.Created(appCtx.GetFiberCtx(), serializers.UserCreateResponse{
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

func LoginUser(fiberCtx *fiber.Ctx) error {
	loginReq := new(serializers.UserLoginRequest)
	if err := fiberCtx.BodyParser(loginReq); err != nil {
		return framework.BadRequest(fiberCtx, "Bad Request")
	}
	
	user, err := db.DBConnector.GetUserByEmail(fiberCtx.Context(), loginReq.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return framework.UnAuthorized(fiberCtx, "Invalid Email is provided")
		}
		return framework.InternalError(fiberCtx)
	}

	if !utils.CheckPassword(loginReq.PassWord, user.PasswordHash) {
		return framework.UnAuthorized(fiberCtx, "Invalid Password")
	}

	sessionExpiresAt := time.Now().Add(time.Hour * 24).Unix() // expires in 24h

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":  sessionExpiresAt,
		"jti":     uuid.NewString(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.AppProp.JwtSecret))
	if err != nil {
		return framework.Error(fiberCtx, 500, "login_error", "Could not login at the moment, Please try again later")
	}

	return framework.Success(fiberCtx, serializers.UserLoginResponse{
		SignedToken: signedToken,
		ExpiresAt: sessionExpiresAt,
		Name: user.Name,
		Email: user.Email,
	})

}

func LogoutUser(fiberCtx *fiber.Ctx) error{
	user := fiberCtx.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)

    jti, ok := claims["jti"].(string)
	if !ok {
		return framework.BadRequest(fiberCtx, "Invalid Credentials")
	}

	tokenBlacklist, err := db.DBConnector.GetBlackListByJti(fiberCtx.Context(), jti)
	if err != nil && err != sql.ErrNoRows {
		return framework.InternalError(fiberCtx)
	}
	if tokenBlacklist.Jti != "" {
		return framework.BadRequest(fiberCtx, "Already logged out")
	}

	expUnix := int64(claims["exp"].(float64))
    expiresAt := time.Unix(expUnix, 0)

	_, err = db.DBConnector.CreateBlackList(fiberCtx.Context(), db.CreateBlackListParams{
		Jti: jti,
		ExpiresAt: expiresAt,
	})

	if err != nil {
		return framework.InternalError(fiberCtx)
	}
	
	return framework.SuccessWithMsg(fiberCtx, "Logout Successful", nil)
}

func HandleBlackListCleanUp(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		log.Println("Starting token_blacklist cleanup!")
        for range ticker.C {
            rowsAffected, err := db.DBConnector.DeleteExpiredBlackList(context.Background()) 
            if err != nil {
                log.Println("Error cleaning up token_blacklist:", err)
				return
            } 
			log.Printf("Deleted %d expired tokens\n", rowsAffected)
        }
    }()

}