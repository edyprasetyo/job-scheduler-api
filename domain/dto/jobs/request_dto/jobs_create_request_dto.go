package request_dto

import (
	"github.com/go-playground/validator/v10"
)

type JobsCreateRequestDTO struct {
	JobName    string `json:"job_name" validate:"required"`
	APIUrl     string `json:"api_url" validate:"required,uniqueJobs"`
	ExecutedAt string `json:"executed_at" validate:"required"`
}

func JobsCreateValidator() *validator.Validate {
	validator := validator.New()
	validator.RegisterValidation("uniqueJobs", uniqueJobs)
	return validator
}

func uniqueJobs(fl validator.FieldLevel) bool {
	// emailField := fl.Parent().FieldByName("Email")
	// confirmField := fl.Field()

	// if emailField.IsValid() && confirmField.String() == emailField.String() {
	// 	return true
	// }

	// return false
	return false
}
