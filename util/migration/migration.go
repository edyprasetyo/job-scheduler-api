package migration

import (
	"jobschedulerapi/domain/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	//  db.AutoMigrate(&model.TrJobsMdl{})
	err := db.AutoMigrate(&model.TrJobsMdl{})
	if err != nil {
		panic("Failed to auto migrate models")
	}
}
