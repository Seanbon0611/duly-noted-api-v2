package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	models "duly_noted/models"
)

var db *gorm.DB

func getUsers(c *gin.Context) {
	var users []models.User
	db.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No users found"})
		return
	}
	for _, user := range users {

		users = append(users, models.User{Username: user.Username, Email: user.Email})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

func init() {
	db, err := gorm.Open(postgres.Open("host=localhost port=5431 user=postgres dbname=duly_noted sslmode=disable"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.User{}, &models.Note{})

	fmt.Println("Database connected")
}

func main() {

	server := gin.Default()

	v1 := server.Group("/api/v1")
	{
		v1.GET("/users", getUsers)
		// 	v1.GET("/users/id", getSingleUser)
		// 	v1.POST("/users/create", createUser)
		// 	v1.DELETE("users/delete/:id", deleteUser)
		// }
		server.Run(":3001")
	}
}
