package util

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	envData := getEnv()
	dsn := "sqlserver://" + envData.DB_USER + ":" + envData.DB_PASSWORD + "@" + envData.DB_HOST + "?database=" + envData.DB_NAME + "&encrypt=disable"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	return db
}
