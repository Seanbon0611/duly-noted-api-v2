package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	Models "duly_noted/Models"
)

var db *gorm.DB

func getUsers(c *gin.Context) {
	var users []Models.User
	db.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No users found"})
		return
	}
	for _, user := range users {

		users = append(users, Models.User{Username: user.Username, Email: user.Email})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

func createUser(c *gin.Context) {
	var user Models.User
	db.Save(user)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User Successfully Created"})
}

func init() {
	db, err := gorm.Open(postgres.Open("host=localhost port=5431 user=postgres dbname=duly_noted sslmode=disable"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&Models.User{}, &Models.Note{})

	fmt.Println("Database connected")
}

func main() {

	server := gin.Default()

	v1 := server.Group("/api/v1")
	{
		v1.GET("/users", getUsers)
		// 	v1.GET("/users/id", getSingleUser)
		v1.POST("/users/create", createUser)
		// 	v1.DELETE("users/delete/:id", deleteUser)
		// }
		server.Run(":3001")
	}
}
