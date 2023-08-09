package extmodel

type ApiResponse struct {
	StatusCode int         `json:"statusCode"`
	Success    bool        `json:"success"`
	Result     interface{} `json:"result"`
}
