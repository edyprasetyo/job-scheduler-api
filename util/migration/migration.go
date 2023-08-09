package migration

import (
	"jobschedulerapi/domain/model"
	"jobschedulerapi/util/config"
)

func Migrate() {
	db := config.InitDB()
	err := db.AutoMigrate(&model.TrJobsMdl{})
	if err != nil {
		panic("Failed to auto migrate models")
	}
}
