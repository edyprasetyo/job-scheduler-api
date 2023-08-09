//go:build wireinject
// +build wireinject

package injection

import (
	"jobschedulerapi/api/controller"
	"jobschedulerapi/api/router"
	"jobschedulerapi/application/repository_impl"
	"jobschedulerapi/application/usecase_impl"
	"jobschedulerapi/infrastructure/service_impl"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var jobsSet = wire.NewSet(
	repository_impl.NewJobsRepository,
	usecase_impl.NewJobsUsecase,
	controller.NewJobsController,
)

func InitPublicRouter(app *fiber.App) router.PublicRouter {
	wire.Build(
		service_impl.NewDatabase,
		jobsSet,
		router.NewPublicRouter,
	)
	return router.PublicRouter{}

}
