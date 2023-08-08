package router

import (
	"jobschedulerapi/injection"
	"jobschedulerapi/util"
	"jobschedulerapi/util/migration"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	db := util.InitDB()
	migration.Migrate(db)

	publicRouter := r.Group("/api/v1")

	jobsController := injection.InitJobsController(db)
	jobsController.Route(publicRouter, jobsController)

	return r
}
