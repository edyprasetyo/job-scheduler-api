package repository

import (
	model "jobschedulerapi/domain/model"
)

type JobsRepository interface {
	Fetch(jobID int) (model.TrJobsMdl, error)
	Insert(model *model.TrJobsMdl) error
	Update(model *model.TrJobsMdl) error
	Delete(model *model.TrJobsMdl) error
	FetchAll(filterExpr string, param map[string]interface{}) ([]model.TrJobsMdl, error)
}
