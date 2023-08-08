package request_dto

type JobsCreateRequestDTO struct {
	JobName    string `json:"job_name" binding:"required"`
	APIUrl     string `json:"api_url" binding:"required"`
	ExecutedAt string `json:"executed_at" binding:"required"`
}
