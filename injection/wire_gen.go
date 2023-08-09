// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injection

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"jobschedulerapi/api/controller"
	"jobschedulerapi/api/router"
	"jobschedulerapi/application/repository_impl"
	"jobschedulerapi/application/usecase_impl"
	"jobschedulerapi/infrastructure/service_impl"
)

// Injectors from wire.go:

func InitPublicRouter(app *fiber.App) router.PublicRouter {
	database := service_impl.NewDatabase()
	jobsRepository := repository_impl.NewJobsRepository(database)
	jobsUsecase := usecase_impl.NewJobsUsecase(jobsRepository, database)
	jobsController := controller.NewJobsController(jobsUsecase)
	publicRouter := router.NewPublicRouter(app, jobsController)
	return publicRouter
}

// wire.go:

var jobsSet = wire.NewSet(repository_impl.NewJobsRepository, usecase_impl.NewJobsUsecase, controller.NewJobsController)
