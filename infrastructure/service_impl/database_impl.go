package service_impl

import (
	"jobschedulerapi/application/service"
	"jobschedulerapi/util/config"

	"gorm.io/gorm"
)

type DatabaseImpl struct {
	DB *gorm.DB
}

func (o *DatabaseImpl) Begin() *gorm.DB {
	tx := o.DB.Begin()
	return tx
}

func (o *DatabaseImpl) Delete(model interface{}) error {
	result := o.DB.Delete(model)
	return result.Error
}

func (o *DatabaseImpl) Execute(sql string, param ...interface{}) error {
	result := o.DB.Exec(sql, param)
	return result.Error
}

func (o *DatabaseImpl) Fetch(destClass interface{}, sql string, param map[string]interface{}) error {
	if param == nil {
		resultNil := o.DB.Raw(sql, nil).Scan(destClass)
		return resultNil.Error
	}
	result := o.DB.Raw(sql, param).Scan(destClass)
	return result.Error
}

func (o *DatabaseImpl) Insert(model interface{}) error {
	result := o.DB.Create(model)
	return result.Error
}

func (o *DatabaseImpl) Update(model interface{}) error {
	result := o.DB.Save(model)
	return result.Error
}

func NewDatabase() service.Database {
	db := config.InitDB()
	return &DatabaseImpl{DB: db}
}
