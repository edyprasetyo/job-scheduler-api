package jobs_dto

type FetchResponseDto struct {
	JobID       int64  `json:"jobID"`
	JobName     string `json:"jobName"`
	APIUrl      string `json:"apiUrl"`
	ExecutedAt  string `json:"executedAt"`
	CreatedTime string `json:"createdTime"`
}
