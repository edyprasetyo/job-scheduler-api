package usecase_impl

import (
	"io"
	"jobschedulerapi/api/exception"
	"jobschedulerapi/domain/dto/jobs_dto"
	"jobschedulerapi/domain/mapper/jobs_mapper"
	"jobschedulerapi/domain/repository"
	"jobschedulerapi/domain/usecase"
	"net/http"
	"time"
)

type JobsUseCaseImpl struct {
	JobsRepository repository.JobsRepository
}

func (o *JobsUseCaseImpl) CheckAndRunJobs() error {
	jobs, err := o.JobsRepository.FetchAll("IsExecuted=0", nil)
	if err != nil {
		return err
	}
	for _, job := range jobs {
		isNeedTobeExecuted := job.ExecutedAt.Before(time.Now())
		if isNeedTobeExecuted {
			res, err := http.Get(job.APIUrl)
			if err != nil {
				return err
			}
			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return err
			}
			job.IsExecuted = true
			job.APIResponse = string(body)

		}
	}

	return nil
}

func (o *JobsUseCaseImpl) Create(dto jobs_dto.CreateRequestDto) (jobs_dto.CreateResponseDto, []exception.ValidationError, error) {

	validation_err := exception.CreateValidator(dto, jobs_dto.RegisterValidation)
	if validation_err != nil {
		return jobs_dto.CreateResponseDto{}, validation_err, nil
	}

	jobs := jobs_mapper.MapCreateRequestDto(&dto)
	err := o.JobsRepository.Insert(jobs)
	if err != nil {
		return jobs_dto.CreateResponseDto{}, nil, err
	}
	return *jobs_mapper.MapCreateResponseDto(jobs), nil, nil
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

func (o *JobsUseCaseImpl) Fetch(jobID int) (jobs_dto.FetchResponseDto, error) {
	jobs, err := o.JobsRepository.Fetch(jobID)
	if err != nil {
		return jobs_dto.FetchResponseDto{}, err
	}
	return *jobs_mapper.MapFetchResponseDto(&jobs), nil
}

func (o *JobsUseCaseImpl) FetchPendingJobs() ([]jobs_dto.FetchPendingJobsResponseDto, error) {
	jobs, err := o.JobsRepository.FetchAll("IsExecuted=0", nil)
	if err != nil {
		return []jobs_dto.FetchPendingJobsResponseDto{}, err
	}
	return jobs_mapper.MapFetchPendingJobsResponseDto(jobs), nil
}

func NewJobsUsecase(jobsRepository repository.JobsRepository) usecase.JobsUsecase {
	return &JobsUseCaseImpl{JobsRepository: jobsRepository}
}
