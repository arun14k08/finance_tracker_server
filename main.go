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
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	if config.AppProp.FrontendUrl == "" {
		if config.AppProp.Environment == "development" {
			config.AppProp.FrontendUrl = "http://localhost:5173"
		} else {
			// fallback - better to fail fast
			// log.Fatal("FRONTEND_URL must be set in production")
			// todo: need to un comment above line after deploying front-end
		}
	}

		// cors middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.AppProp.FrontendUrl,
		AllowMethods:     "GET,HEAD,PUT,POST,DELETE",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		ExposeHeaders:    "Content-Length",
		MaxAge:          3600,
	}))

		// logger middleware
	app.Use(logger.New())

	// JWT middleware
	app.Use("/api", jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{Key: []byte(config.AppProp.JwtSecret)},
	}))


	// health check route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// auth routes
	app.Post("/register", handlers.CreateUser)
	app.Post("/login", handlers.LoginUser)
	app.Post("/api/logout", handlers.LogoutUser)

	// account routes
	app.Get("/api/accounts/:id", handlers.GetAccountByID)
	app.Get("/api/accounts", handlers.GetAccounts)
	app.Post("/api/accounts", handlers.CreateAccount)
	app.Put("/api/accounts", handlers.UpdateAccount)
	app.Delete("/api/accounts/:id", handlers.DeleteAccount)

	// transaction routes
	// app.Get("/api/transactions/:id", handlers.GetTransactionByID)
	// app.Get("/api/transactions", handlers.GetTransactions)
	// app.Post("/api/transactions", handlers.CreateTransaction)
	// app.Put("/api/transactions", handlers.UpdateTransaction)
	// app.Delete("/api/transactions/:id", handlers.DeleteTransaction)

	// cron jobs
	handlers.HandleBlackListCleanUp(time.Hour)
	// server startup
	logging.LogMessage("Stating Server on PORT: " + appProp.ServerPort)
	app.Listen(":" + appProp.ServerPort)

}
