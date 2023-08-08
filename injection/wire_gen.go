// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injection

import (
	"gorm.io/gorm"
	"jobschedulerapi/api/controller"
	"jobschedulerapi/application/repository_impl"
	"jobschedulerapi/application/usecase_impl"
	"jobschedulerapi/infrastructure/service_impl"
)

// Injectors from wire.go:

func InitJobsController(db *gorm.DB) *controller.JobsController {
	database := service_impl.NewDatabase(db)
	jobsRepository := repository_impl.NewJobsRepository(database)
	jobsUsecase := usecase_impl.NewJobsUsecase(jobsRepository)
	jobsController := controller.NewJobsController(jobsUsecase)
	return jobsController
}