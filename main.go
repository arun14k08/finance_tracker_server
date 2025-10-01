package main

import (
	"log"

	"github.com/arun14k08/finance_tracker_server/pkg/db"
	"github.com/arun14k08/finance_tracker_server/pkg/handlers"
	"github.com/arun14k08/goframework/config"
	"github.com/arun14k08/goframework/framework"
	"github.com/arun14k08/goframework/logging"
	"github.com/arun14k08/goframework/service"
	"github.com/gofiber/fiber/v2"
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
	app.Post("/users", handlers.CreateUser)
	app.Get("/users", handlers.GetUser)

	logging.LogMessage("Stating Server on PORT: " + appProp.ServerPort)
	app.Listen(":" + appProp.ServerPort)
}
