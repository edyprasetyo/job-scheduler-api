package jobs_dto

type FetchResponseDto struct {
	JobID       int64  `json:"job_id"`
	JobName     string `json:"job_name"`
	APIUrl      string `json:"api_url"`
	ExecutedAt  string `json:"executed_at"`
	CreatedTime string `json:"created_time"`
}
