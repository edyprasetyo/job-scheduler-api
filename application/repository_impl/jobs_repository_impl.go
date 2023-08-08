package repository_impl

import (
	"jobschedulerapi/application/service"
	"jobschedulerapi/domain/model"
	"jobschedulerapi/domain/repository"
)

type JobsRepositoryImpl struct {
	DB service.Database
}

func (o *JobsRepositoryImpl) Delete(model *model.TrJobsMdl) error {
	err := o.DB.Delete(model)
	return err
}

func (o *JobsRepositoryImpl) Fetch(jobID int) (model.TrJobsMdl, error) {
	var job model.TrJobsMdl
	err := o.DB.Fetch(&job, "SELECT * FROM TrJobs WHERE JobID = @JobID", map[string]interface{}{"JobID": jobID})
	return job, err
}

func (o *JobsRepositoryImpl) FetchAll(filterExpr string, param map[string]interface{}) ([]model.TrJobsMdl, error) {
	var jobs []model.TrJobsMdl
	if filterExpr == "" {
		filterExpr = "1 = 1"
	}

	err := o.DB.Fetch(&jobs, "SELECT * FROM TrJobs WHERE "+filterExpr, param)
	return jobs, err
}

func (o *JobsRepositoryImpl) Insert(model *model.TrJobsMdl) error {
	err := o.DB.Insert(model)
	return err
}

func (o *JobsRepositoryImpl) Update(model *model.TrJobsMdl) error {
	err := o.DB.Update(model)
	return err
}

func NewJobsRepository(db service.Database) repository.JobsRepository {
	return &JobsRepositoryImpl{DB: db}
}
