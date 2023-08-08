package service

type Database interface {
	BeginTransaction() error
	CommitTransaction() error
	RollbackTransaction() error
	Fetch(destClass interface{}, sql string, param map[string]interface{}) error
	Insert(model interface{}) error
	Update(model interface{}) error
	Delete(model interface{}) error
	Execute(sql string, param ...interface{}) error
}
