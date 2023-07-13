package services

import (
	"be-pc-ipm-go/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupCors(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     config.AllowHeaders(),
		AllowMethods:     config.AllowMethods(),
		AllowCredentials: true,
	}))
}
