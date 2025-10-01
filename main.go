package main

import (
	"github.com/arun14k08/finance_tracker_server/pkg/handlers"
	"github.com/arun14k08/goframework/config"
	"github.com/arun14k08/goframework/framework"
	"github.com/arun14k08/goframework/logging"
	"github.com/arun14k08/goframework/service"
	"github.com/gofiber/fiber/v2"
)



func main() {
	service.Start(framework.FrameWorkService{})
	appProp := config.GetAppProp()
	app := fiber.New()
	app.Post("/users", handlers.CreateUser)
	app.Get("/users", handlers.GetUser)

	logging.LogMessage("Stating Server on PORT: " + appProp.ServerPort)
	app.Listen(":" + appProp.ServerPort)
}
