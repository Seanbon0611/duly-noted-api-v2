package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

//Initiate Database
func Init() {

	//envirnment variables for credentials
	dbHost := goDotEnvVariable("HOST")
	dbPort := goDotEnvVariable("PORT")
	dbUser := goDotEnvVariable("DBUSER")
	dbName := goDotEnvVariable("DBNAME")
	dbPass := goDotEnvVariable("PASSWORD")
	dbSsl := goDotEnvVariable("SSLMODE")
	dbCred := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, dbUser, dbName, dbPass, dbSsl)

	database, err := gorm.Open(postgres.Open(dbCred), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //Logs all SQL queries into console
	})
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Database connected")
	DB = database
}
