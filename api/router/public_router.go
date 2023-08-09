package router

import (
	"jobschedulerapi/api/controller"

	"github.com/gofiber/fiber/v2"
)

type PublicRouter struct {
	Public fiber.Router
}

func NewPublicRouter(app *fiber.App, c *controller.JobsController) PublicRouter {
	router := app.Group("/api/v1")

	controller.JobsRoute(router, c)

	return PublicRouter{Public: router}
}
