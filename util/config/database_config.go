package config

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	envData := GetEnv()
	dsn := "sqlserver://" + envData.DB_USER + ":" + envData.DB_PASSWORD + "@" + envData.DB_HOST + "?database=" + envData.DB_NAME + "&encrypt=disable"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	return db
}
