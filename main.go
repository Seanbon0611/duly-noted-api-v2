package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	type User struct {
		gorm.Model
		Username string `gorm:"unique; not null"`
		Email    string `gorm:"unique; not null"`
		Password string `gorm:"not null`
	}

	type Note struct {
		gorm.Model
		UserID  uint
		Content string
		User    User
	}

	db, err := gorm.Open(postgres.Open("host=localhost port=5431 user=postgres dbname=duly_noted sslmode=disable"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&User{}, &Note{})

	test_user := User{Username: "testuser3", Password: "password"}
	db.Create(&test_user)
	fmt.Println(test_user.ID)
	fmt.Println("Database connected")
	server := gin.Default()

	server.GET("/:username/notes")

	server.Run(":3001")
}
