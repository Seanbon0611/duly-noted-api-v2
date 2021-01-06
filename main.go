package main

import (
	"duly_noted/config"
	"duly_noted/controllers"
	"duly_noted/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//Initiate server
	server := gin.Default()

	//Connection to Database
	config.Init()
	config.DB.AutoMigrate(&models.User{}, &models.Note{})

	//Routes
	server.GET("/api/v1/users", controllers.GetUsers)
	server.GET("/api/v1/users/:id", controllers.GetSingleUser)
	server.POST("/api/v1/users/create", controllers.SignupUser)
	server.POST("/api/v1/login", controllers.LoginUser)
	server.DELETE("/api/v1/users/delete/:id", controllers.DeleteUser)

	//Start Server
	server.Run(":3001")

}
