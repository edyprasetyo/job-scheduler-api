package ex

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string
	Message string
	Tag     string
}

func getAllValidation(err error) []ValidationError {
	// err =
	// Key: 'CreateRequestDto.JobName' Error:Field validation for 'JobName' failed on the 'required' tag
	// Key: 'CreateRequestDto.APIUrl' Error:Field validation for 'APIUrl' failed on the 'uniqueJobs' tag
	if err == nil {
		return nil
	}
	validationError := make([]ValidationError, 0)

	errors := strings.Split(err.Error(), "\n")
	for _, error := range errors {
		if error == "" {
			continue
		}
		splittedError := strings.Split(error, " ")
		validationError = append(validationError, ValidationError{
			Field:   strings.ReplaceAll(splittedError[5], "'", ""),
			Message: "",
			Tag:     strings.ReplaceAll(splittedError[9], "'", ""),
		})
	}
	return validationError
}

func CreateValidator(dto interface{}, registerValidation func(*validator.Validate) []ValidationError) []ValidationError {
	validator := validator.New()
	errMessage := registerValidation(validator)
	err := validator.Struct(dto)
	listValidation := getAllValidation(err)
	for i, validation := range listValidation {
		for _, allError := range errMessage {
			if validation.Tag == allError.Tag && validation.Field == allError.Field {
				listValidation[i].Message = allError.Message
			}
		}
	}
	return listValidation
}
