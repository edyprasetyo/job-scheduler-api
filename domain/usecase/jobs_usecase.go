package usecase

import (
	ex "jobschedulerapi/api/exception"
	"jobschedulerapi/domain/dto/jobs_dto"
)

type JobsUsecase interface {
	Create(dto jobs_dto.CreateRequestDto) (jobs_dto.CreateResponseDto, []ex.ValidationError, error)
	Delete(jobID int) error
	Fetch(jobID int) (jobs_dto.FetchResponseDto, error)
	FetchPendingJobs() ([]jobs_dto.FetchPendingJobsResponseDto, error)
}
