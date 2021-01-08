package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

//Initiate Database
func Init() {
	database, err := gorm.Open(postgres.Open("host=localhost port=5431 user=postgres dbname=duly_noted sslmode=disable"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //Logs all SQL queries into console
	})
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Database connected")
	DB = database
}
