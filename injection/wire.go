//go:build wireinject
// +build wireinject

package injection

import (
	"jobschedulerapi/api/controller"
	"jobschedulerapi/application/repository_impl"
	"jobschedulerapi/application/usecase_impl"
	"jobschedulerapi/infrastructure/service_impl"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitJobsController(db *gorm.DB) *controller.JobsController {
	wire.Build(
		service_impl.NewDatabase,
		repository_impl.NewJobsRepository,
		usecase_impl.NewJobsUsecase,
		controller.NewJobsController,
	)
	return &controller.JobsController{}

}
