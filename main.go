package main

import (
	"log"
	"time"

	"github.com/arun14k08/finance_tracker_server/pkg/db"
	"github.com/arun14k08/finance_tracker_server/pkg/handlers"
	"github.com/arun14k08/goframework/config"
	"github.com/arun14k08/goframework/framework"
	"github.com/arun14k08/goframework/logging"
	"github.com/arun14k08/goframework/service"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {
	service.Start(framework.FrameWorkService{})
	if err := db.Connect(); err != nil {
        log.Fatal("Database connection failed:", err)
    }
    defer db.Close()
	appProp := config.GetAppProp()
	app := fiber.New()
	
	// JWT middleware
	app.Use("/api", jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{Key: []byte(config.AppProp.JwtSecret)},
	}))
	// logger middleware
	app.Use(logger.New())
	// routes
	app.Post("/register", handlers.CreateUser)
	app.Post("/login", handlers.LoginUser)
	app.Post("/api/logout", handlers.LogoutUser)

	// cron jobs
	handlers.HandleBlackListCleanUp(time.Hour)
	// server startup
	logging.LogMessage("Stating Server on PORT: " + appProp.ServerPort)
	app.Listen(":" + appProp.ServerPort)

}
