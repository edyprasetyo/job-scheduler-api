package jobs_dto

import (
	ex "jobschedulerapi/api/exception"
	"jobschedulerapi/util/tools"

	"github.com/go-playground/validator/v10"
)

type CreateRequestDto struct {
	JobName    string `json:"job_name" validate:"required"`
	APIUrl     string `json:"api_url" validate:"required,uniqueJobs"`
	ExecutedAt string `json:"executed_at" validate:"required,allowedDate"`
}

func CreateRequestValidation(v *validator.Validate) []ex.ValidationError {
	v.RegisterValidation("uniqueJobs", uniqueJobs)
	v.RegisterValidation("allowedDate", allowedDate)

	return []ex.ValidationError{
		{
			Field:   "JobName",
			Message: "Job name belum diisi",
			Tag:     "required",
		},
		{
			Field:   "APIUrl",
			Message: "API url belum diisi",
			Tag:     "required",
		},
		{
			Field:   "ExecutedAt",
			Message: "Executed at belum diisi",
			Tag:     "required",
		},
		{
			Field:   "APIUrl",
			Message: "API url sudah ada",
			Tag:     "uniqueJobs",
		},
		{
			Field:   "ExecutedAt",
			Message: "Format tanggal salah (dd/MM/yyyy HH:mm:ss)",
			Tag:     "allowedDate",
		},
	}
}

func uniqueJobs(fl validator.FieldLevel) bool {
	// emailField := fl.Parent().FieldByName("Email")
	// confirmField := fl.Field()
	return true
}

func allowedDate(fl validator.FieldLevel) bool {
	executedAt := fl.Field()
	validFormat := "dd/MM/yyyy HH:mm:ss"
	time := tools.StringToDate(executedAt.String(), validFormat)
	return time != nil
}
