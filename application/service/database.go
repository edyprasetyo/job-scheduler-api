package service

import (
	"gorm.io/gorm"
)

type Database interface {
	Begin() *gorm.DB
	Fetch(destClass interface{}, sql string, param map[string]interface{}) error
	Insert(model interface{}) error
	Update(model interface{}) error
	Delete(model interface{}) error
	Execute(sql string, param ...interface{}) error
}
