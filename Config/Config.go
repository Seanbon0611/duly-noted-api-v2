package Config

import (
	"fmt"

	Models "duly_noted/Models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	database, err := gorm.Open(postgres.Open("host=localhost port=5431 user=postgres dbname=duly_noted sslmode=disable"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to database")
	}
	database.AutoMigrate(&Models.User{}, &Models.Note{})

	fmt.Println("Database connected")

	DB = database
}
