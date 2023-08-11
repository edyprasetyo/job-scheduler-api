package jobs_mapper

import (
	"jobschedulerapi/domain/dto/jobs_dto"
	"jobschedulerapi/domain/model"
	"jobschedulerapi/util/datetime"
)

func MapCreateRequestDto(dto *jobs_dto.CreateRequestDto) *model.TrJobsMdl {
	return &model.TrJobsMdl{
		JobName:     dto.JobName,
		APIUrl:      dto.APIUrl,
		IsExecuted:  false,
		ExecutedAt:  *datetime.FromString(dto.ExecutedAt, "dd/MM/yyyy HH:mm:ss"),
		CreatedTime: datetime.Now(),
	}
}

func MapCreateResponseDto(jobs *model.TrJobsMdl) *jobs_dto.CreateResponseDto {
	return &jobs_dto.CreateResponseDto{
		JobName:     jobs.JobName,
		APIUrl:      jobs.APIUrl,
		ExecutedAt:  datetime.ToString(jobs.ExecutedAt, "dd/MM/yyyy HH:mm:ss"),
		CreatedTime: datetime.ToString(jobs.CreatedTime, "dd/MM/yyyy HH:mm:ss"),
	}
}

func MapFetchResponseDto(jobs *model.TrJobsMdl) *jobs_dto.FetchResponseDto {
	return &jobs_dto.FetchResponseDto{
		JobID:       jobs.JobID,
		JobName:     jobs.JobName,
		APIUrl:      jobs.APIUrl,
		ExecutedAt:  datetime.ToString(jobs.ExecutedAt, "dd/MM/yyyy HH:mm:ss"),
		CreatedTime: datetime.ToString(jobs.CreatedTime, "dd/MM/yyyy HH:mm:ss"),
	}
}

func MapFetchPendingJobsResponseDto(jobs []model.TrJobsMdl) []jobs_dto.FetchPendingJobsResponseDto {
	var result = make([]jobs_dto.FetchPendingJobsResponseDto, 0)
	for _, job := range jobs {
		result = append(result, jobs_dto.FetchPendingJobsResponseDto{
			JobID:       job.JobID,
			JobName:     job.JobName,
			APIUrl:      job.APIUrl,
			ExecutedAt:  datetime.ToString(job.ExecutedAt, "dd/MM/yyyy HH:mm:ss"),
			CreatedTime: datetime.ToString(job.CreatedTime, "dd/MM/yyyy HH:mm:ss"),
		})
	}
	return result
}
