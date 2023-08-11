package jobs_dto

import (
	ex "jobschedulerapi/api/exception"
	"jobschedulerapi/util/tools"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CreateRequestDto struct {
	JobName    string `json:"jobName" validate:"required"`
	APIUrl     string `json:"apiUrl" validate:"required,validAPIUrl"`
	Password   string `json:"password" validate:"required,passwordValid"`
	ExecutedAt string `json:"executedAt" validate:"required,allowedDate"`
}

func CreateRequestValidation(v *validator.Validate) []ex.ValidationError {
	v.RegisterValidation("allowedDate", allowedDate)
	v.RegisterValidation("validAPIUrl", validAPIUrl)
	v.RegisterValidation("passwordValid", passwordValid)

	return []ex.ValidationError{
		{
			Field:   "JobName",
			Message: "Job name belum diisi",
			Tag:     "required",
		},
		{
			Field:   "Password",
			Message: "Password belum diisi",
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
			Message: "API url tidak valid, harus mengandung http atau https, contoh: http://localhost:8080",
			Tag:     "validAPIUrl",
		},
		{
			Field:   "Password",
			Message: "Password tidak sesuai",
			Tag:     "passwordValid",
		},
		{
			Field:   "ExecutedAt",
			Message: "Format tanggal salah (dd/MM/yyyy HH:mm:ss)",
			Tag:     "allowedDate",
		},
	}
}

func validAPIUrl(fl validator.FieldLevel) bool {
	// emailField := fl.Parent().FieldByName("Email")
	apiUrl := fl.Field().String()
	return strings.Contains(apiUrl, "http") || strings.Contains(apiUrl, "https")
}

func allowedDate(fl validator.FieldLevel) bool {
	executedAt := fl.Field().String()
	validFormat := "dd/MM/yyyy HH:mm:ss"
	time := tools.StringToDate(executedAt, validFormat)
	return time != nil
}

func passwordValid(fl validator.FieldLevel) bool {
	validPassword := "Indorent10!"
	password := fl.Field().String()
	return password == validPassword
}
