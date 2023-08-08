package usecase_impl

import (
	"jobschedulerapi/domain/dto/jobs/request_dto"
	"jobschedulerapi/domain/dto/jobs/response_dto"
	"jobschedulerapi/domain/mapper/jobs_mapper"
	"jobschedulerapi/domain/repository"
	"jobschedulerapi/domain/usecase"
)

type JobsUseCaseImpl struct {
	JobsRepository repository.JobsRepository
}

func (o *JobsUseCaseImpl) Create(dto request_dto.JobsCreateRequestDTO) (response_dto.JobsCreateResponseDTO, error) {
	jobs := jobs_mapper.MapCreateRequestDto(&dto)
	err := o.JobsRepository.Insert(jobs)
	if err != nil {
		return response_dto.JobsCreateResponseDTO{}, err
	}
	return *jobs_mapper.MapCreateResponseDto(jobs), nil
}

func (o *JobsUseCaseImpl) Delete(jobID int) error {
	jobs, err := o.JobsRepository.Fetch(jobID)
	if err != nil {
		return err
	}
	err = o.JobsRepository.Delete(&jobs)
	if err != nil {
		return err
	}
	return nil
}

func (o *JobsUseCaseImpl) Fetch(jobID int) (response_dto.JobsFetchResponseDTO, error) {
	jobs, err := o.JobsRepository.Fetch(jobID)
	if err != nil {
		return response_dto.JobsFetchResponseDTO{}, err
	}
	return *jobs_mapper.MapFetchResponseDto(&jobs), nil
}

func (o *JobsUseCaseImpl) FetchPendingJobs() ([]response_dto.JobsFetchPendingJobsResponseDTO, error) {
	jobs, err := o.JobsRepository.FetchAll("IsExecuted=0", nil)
	if err != nil {
		return []response_dto.JobsFetchPendingJobsResponseDTO{}, err
	}
	return jobs_mapper.MapFetchAllResponseDto(jobs), nil
}

func NewJobsUsecase(jobsRepository repository.JobsRepository) usecase.JobsUsecase {
	return &JobsUseCaseImpl{JobsRepository: jobsRepository}
}
