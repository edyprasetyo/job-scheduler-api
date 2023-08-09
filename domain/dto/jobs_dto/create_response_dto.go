package jobs_dto

type CreateResponseDto struct {
	JobName     string `json:"jobName"`
	APIUrl      string `json:"apiUrl"`
	ExecutedAt  string `json:"executedAt"`
	CreatedTime string `json:"createdTime"`
}
