package config

import (
	"fmt"
	"os"

	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

//Initiate Database
func Init() {
	os.Setenv("HOST", "dulynoted.c3jtcwuysqgr.us-west-1.rds.amazonaws.com")
	os.Setenv("PORT", "5432")
	os.Setenv("DBNAME", "postgres")
	os.Setenv("DBUSER", "postgres")
	os.Setenv("PASSWORD", "J3tBlu30611")
	os.Setenv("SSLMODE", "disable")

	dbHost := os.Getenv("HOST")
	dbPort := os.Getenv("PORT")
	dbUser := os.Getenv("DBUSER")
	dbName := os.Getenv("DBNAME")
	dbPass := os.Getenv("PASSWORD")
	dbSsl := os.Getenv("SSLMODE")
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
