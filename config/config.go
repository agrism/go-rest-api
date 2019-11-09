package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
}

var DbConnection DatabaseConfigurations

func initConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DbConnection.DBName = os.Getenv("DB_NAME")
	DbConnection.DBUser = os.Getenv("DB_USER")
	DbConnection.DBPassword = os.Getenv("DB_PASSWORD")
	DbConnection.DBPort = os.Getenv("DB_PORT")
	DbConnection.DBHost = os.Getenv("DB_HOST")

}

func GetDbConfig() DatabaseConfigurations {

	initConfig()
	return DbConnection
}
