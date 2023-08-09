package controller

import (
	"jobschedulerapi/api/handler"
	"jobschedulerapi/domain/dto/jobs/request_dto"
	"jobschedulerapi/domain/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type JobsController struct {
	JobsUseCase usecase.JobsUsecase
}

func NewJobsController(jobsUseCase usecase.JobsUsecase) *JobsController {
	return &JobsController{JobsUseCase: jobsUseCase}
}

func JobsRoute(r fiber.Router, c *JobsController) {
	group := r.Group("/jobs")
	group.Get("/:jobID", c.Fetch)
	group.Post("", c.Create)
	group.Delete("/:jobID", c.Delete)
	group.Get("/all/pending", c.FetchPendingJobs)
}

func (o *JobsController) Create(c *fiber.Ctx) error {
	req := request_dto.JobsCreateRequestDTO{}
	err := c.BodyParser(&req)
	if err != nil {
		return handler.ValidationError(c, err.Error())
	}

	res, err := o.JobsUseCase.Create(req)
	if err != nil {
		return handler.InternalError(c, err.Error())
	}
	return handler.Success(c, res)
}

func (o *JobsController) Delete(c *fiber.Ctx) error {
	jobIDStr := c.Params("jobID")
	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		return handler.ValidationError(c, err.Error())
	}

	err = o.JobsUseCase.Delete(jobID)
	if err != nil {
		return handler.InternalError(c, err.Error())

	}
	return handler.Success(c, nil)
}

func (o *JobsController) Fetch(c *fiber.Ctx) error {
	jobIDStr := c.Params("jobID")
	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		return handler.ValidationError(c, err.Error())
	}

	res, err := o.JobsUseCase.Fetch(jobID)
	if err != nil {
		return handler.InternalError(c, err.Error())

	}
	return handler.Success(c, res)
}

func (o *JobsController) FetchPendingJobs(c *fiber.Ctx) error {
	res, err := o.JobsUseCase.FetchPendingJobs()
	if err != nil {
		return handler.InternalError(c, err.Error())
	}
	return handler.Success(c, res)
}

func (o *JobsController) CheckJob(c *fiber.Ctx) error {
	res, err := o.JobsUseCase.FetchPendingJobs()
	if err != nil {
		return handler.InternalError(c, err.Error())
	}

	return handler.Success(c, res)
}
