package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvData struct {
	DB_HOST     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
}

func GetEnv() EnvData {
	goEnv := os.Getenv("GO_ENVIRONMENT")
	suffix := ""
	if goEnv == "PRODUCTION" || goEnv == "production" {
		suffix = ".production"
	} else if goEnv == "STAGING" || goEnv == "staging" {
		suffix = ".staging"
	}

	err := godotenv.Load("env/.env" + suffix)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return EnvData{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
	}
}
