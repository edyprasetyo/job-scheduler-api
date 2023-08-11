package app

import (
	"jobschedulerapi/api/middleware"
	"jobschedulerapi/injection"
	"jobschedulerapi/util/config"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func NewApp() *fiber.App {
	app := fiber.New(config.NewFiberConfig())

	app.Use(middleware.Recovery)
	app.Use(middleware.Error)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	injection.InitPublicRouter(app)

	return app

}
