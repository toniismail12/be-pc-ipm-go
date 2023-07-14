package routers

import (
	"be-pc-ipm-go/handler"
	"be-pc-ipm-go/services"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// setup cors
	services.SetupCors(app)

	app.Get("/", handler.AppName)
	app.Get("/app-name", handler.AppName)

}
