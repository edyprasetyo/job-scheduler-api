package jobs_dto

type CreateResponseDto struct {
	JobName     string `json:"job_name"`
	APIUrl      string `json:"api_url"`
	ExecutedAt  string `json:"executed_at"`
	CreatedTime string `json:"created_time"`
}
