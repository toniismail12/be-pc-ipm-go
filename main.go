package main

import (
	"be-pc-ipm-go/config"
	"be-pc-ipm-go/routers"
	"be-pc-ipm-go/services"
)

func main() {

	// start fiber
	app := services.CreateApp()

	// router
	routers.Setup(app)

	// get port
	port := config.Env("PORT")

	// run
	app.Listen(":" + port)

}
