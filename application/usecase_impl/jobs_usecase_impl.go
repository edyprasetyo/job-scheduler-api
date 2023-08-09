package usecase_impl

import (
	ex "jobschedulerapi/api/exception"
	"jobschedulerapi/application/service"
	"jobschedulerapi/domain/dto/jobs_dto"
	"jobschedulerapi/domain/mapper/jobs_mapper"
	"jobschedulerapi/domain/repository"
	"jobschedulerapi/domain/usecase"
)

type JobsUseCaseImpl struct {
	JobsRepository repository.JobsRepository
	DB             service.Database
}

func (o *JobsUseCaseImpl) Create(dto jobs_dto.CreateRequestDto) (jobs_dto.CreateResponseDto, []ex.ValidationError, error) {
	tx := o.DB.Begin()
	validation_err := ex.CreateValidator(dto, jobs_dto.CreateRequestValidation)
	if validation_err != nil {
		tx.Rollback()
		return jobs_dto.CreateResponseDto{}, validation_err, nil
	}

	jobs := jobs_mapper.MapCreateRequestDto(&dto)
	err := o.JobsRepository.Insert(jobs)
	if err != nil {
		tx.Rollback()
		return jobs_dto.CreateResponseDto{}, nil, err
	}
	tx.Commit()
	return *jobs_mapper.MapCreateResponseDto(jobs), nil, nil
}

func (o *JobsUseCaseImpl) Delete(jobID int) error {
	tx := o.DB.Begin()
	jobs, err := o.JobsRepository.Fetch(jobID)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = o.JobsRepository.Delete(&jobs)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
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

func NewJobsUsecase(jobsRepository repository.JobsRepository, db service.Database) usecase.JobsUsecase {
	return &JobsUseCaseImpl{JobsRepository: jobsRepository, DB: db}
}
