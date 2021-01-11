package main

import (
	"duly_noted/config"
	"duly_noted/controllers"
	"duly_noted/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//Initiate server
	server := gin.Default()

	//Cors
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//Connection to Database
	config.Init()
	config.DB.AutoMigrate(&models.User{}, &models.Note{})

	//Routes

	//User
	server.GET("/api/v1/users", controllers.GetUsers)
	server.GET("/api/v1/users/:id", controllers.GetSingleUser)
	server.POST("/api/v1/users/create", controllers.SignupUser)
	server.DELETE("/api/v1/users/delete/:id", controllers.DeleteUser)
	//Auth
	server.POST("/api/v1/login", controllers.Login)

	//Note
	server.POST("/api/v1/notes/create", controllers.CreateNote)

	//Start Server
	server.Run(":3001")

}
