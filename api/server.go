package api

import (
	"fmt"
	"jobschedulerapi/api/middleware"
	_ "jobschedulerapi/docs"
	"jobschedulerapi/injection"
	"jobschedulerapi/util/config"
	"jobschedulerapi/util/migration"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func Run() {
	migration.Migrate()

	app := fiber.New(config.NewFiberConfig())

	app.Use(middleware.RecoveryMiddleware)
	app.Use(middleware.ErrorMiddleware)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	injection.InitPublicRouter(app)
	port := 8080
	fmt.Printf("Server is running at http://localhost:%d\n", port)
	app.Listen(fmt.Sprintf(":%d", port))
}
