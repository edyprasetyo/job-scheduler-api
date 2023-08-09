package api

import (
	"fmt"
	"jobschedulerapi/api/middleware"
	"jobschedulerapi/injection"
	"jobschedulerapi/util/config"
	"jobschedulerapi/util/migration"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	migration.Migrate()

	app := fiber.New(config.NewFiberConfig())
	app.Use(middleware.RecoveryMiddleware)
	app.Use(middleware.ErrorMiddleware)
	injection.InitPublicRouter(app)
	port := 8080
	fmt.Printf("Server is running at http://localhost:%d\n", port)
	app.Listen(fmt.Sprintf(":%d", port))
}