package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	type User struct {
		gorm.Model
		Username string
		Email    string
		Password string
	}

	type Note struct {
		gorm.Model
		UserID  uint
		Content string
		User    User
	}

	db, err := gorm.Open(postgres.Open("host=localhost port=5431 user=postgres dbname=test_db sslmode=disable"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&User{}, &Note{})
	fmt.Println("Database connected")
	server := gin.Default()

	server.GET("/:user/notes")

	server.Run(":3001")
}
