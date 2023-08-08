package jobs_mapper

import (
	"jobschedulerapi/domain/dto/jobs/request_dto"
	"jobschedulerapi/domain/dto/jobs/response_dto"
	"jobschedulerapi/domain/model"
	"jobschedulerapi/util/tools"

	"time"
)

func MapCreateRequestDto(dto *request_dto.JobsCreateRequestDTO) *model.TrJobsMdl {
	return &model.TrJobsMdl{
		JobName:     dto.JobName,
		APIUrl:      dto.APIUrl,
		IsExecuted:  false,
		ExecutedAt:  tools.StringToDate(dto.ExecutedAt, "dd/MM/yyyy HH:mm:ss"),
		CreatedTime: time.Now(),
	}
}

func MapCreateResponseDto(jobs *model.TrJobsMdl) *response_dto.JobsCreateResponseDTO {
	return &response_dto.JobsCreateResponseDTO{
		JobName:     jobs.JobName,
		APIUrl:      jobs.APIUrl,
		ExecutedAt:  tools.DateToString(jobs.ExecutedAt, "dd/MM/yyyy HH:mm:ss"),
		CreatedTime: tools.DateToString(jobs.CreatedTime, "dd/MM/yyyy HH:mm:ss"),
	}
}

func MapFetchResponseDto(jobs *model.TrJobsMdl) *response_dto.JobsFetchResponseDTO {
	return &response_dto.JobsFetchResponseDTO{
		JobName:     jobs.JobName,
		APIUrl:      jobs.APIUrl,
		ExecutedAt:  tools.DateToString(jobs.ExecutedAt, "dd/MM/yyyy HH:mm:ss"),
		CreatedTime: tools.DateToString(jobs.CreatedTime, "dd/MM/yyyy HH:mm:ss"),
	}
}

func MapFetchAllResponseDto(jobs []model.TrJobsMdl) []response_dto.JobsFetchAllResponseDTO {
	var result []response_dto.JobsFetchAllResponseDTO
	for _, job := range jobs {
		result = append(result, response_dto.JobsFetchAllResponseDTO{
			JobName:     job.JobName,
			APIUrl:      job.APIUrl,
			ExecutedAt:  tools.DateToString(job.ExecutedAt, "dd/MM/yyyy HH:mm:ss"),
			CreatedTime: tools.DateToString(job.CreatedTime, "dd/MM/yyyy HH:mm:ss"),
		})
	}
	return result
}
