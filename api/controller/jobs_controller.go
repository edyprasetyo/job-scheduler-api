package controller

import (
	"jobschedulerapi/domain/dto/jobs/request_dto"
	"jobschedulerapi/domain/usecase"
	"jobschedulerapi/util/response_handler"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobsController struct {
	JobsUseCase usecase.JobsUsecase
}

func NewJobsController(jobsUseCase usecase.JobsUsecase) *JobsController {
	return &JobsController{JobsUseCase: jobsUseCase}
}

func (o *JobsController) Route(group *gin.RouterGroup, controller *JobsController) {
	group.GET("/jobs/:jobID", controller.Fetch)
	group.POST("/jobs", controller.Create)
	group.DELETE("/jobs/:jobID", controller.Delete)
	group.GET("/jobs/pending", controller.FetchPendingJobs)
}

func (o *JobsController) Create(c *gin.Context) {
	req := request_dto.JobsCreateRequestDTO{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response_handler.ValidationError(c, err.Error())
		return
	}

	res, err := o.JobsUseCase.Create(req)
	if err != nil {
		response_handler.InternalError(c, err.Error())
		return
	}
	response_handler.Success(c, res)
}

func (o *JobsController) Delete(c *gin.Context) {
	jobIDStr := c.Param("jobID")
	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		response_handler.ValidationError(c, err.Error())
		return
	}

	err = o.JobsUseCase.Delete(jobID)
	if err != nil {
		response_handler.InternalError(c, err.Error())
		return
	}
	response_handler.Success(c, nil)
}

func (o *JobsController) Fetch(c *gin.Context) {
	jobIDStr := c.Param("jobID")
	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		response_handler.ValidationError(c, err.Error())
		return
	}

	res, err := o.JobsUseCase.Fetch(jobID)
	if err != nil {
		response_handler.InternalError(c, err.Error())
		return
	}
	response_handler.Success(c, res)
}

func (o *JobsController) FetchPendingJobs(c *gin.Context) {
	res, err := o.JobsUseCase.FetchPendingJobs()
	if err != nil {
		response_handler.InternalError(c, err.Error())
		return
	}
	response_handler.Success(c, res)
}

func (o *JobsController) CheckJob(c *gin.Context) {
	res, err := o.JobsUseCase.FetchPendingJobs()
	if err != nil {
		response_handler.InternalError(c, err.Error())
		return
	}

	response_handler.Success(c, res)
}
