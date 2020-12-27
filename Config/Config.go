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
	DB, err := gorm.Open(postgres.Open("host=localhost port=5431 user=postgres dbname=duly_noted sslmode=disable"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	DB.AutoMigrate(&Models.User{}, &Models.Note{})

	fmt.Println("Database connected")
}
