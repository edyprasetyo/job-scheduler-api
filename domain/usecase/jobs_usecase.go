package usecase

import (
	"jobschedulerapi/domain/dto/jobs/request_dto"
	"jobschedulerapi/domain/dto/jobs/response_dto"
)

type JobsUsecase interface {
	Create(dto request_dto.JobsCreateRequestDTO) (response_dto.JobsCreateResponseDTO, error)
	Delete(jobID int) error
	Fetch(jobID int) (response_dto.JobsFetchResponseDTO, error)
	FetchPendingJobs() ([]response_dto.JobsFetchPendingJobsResponseDTO, error)
}
