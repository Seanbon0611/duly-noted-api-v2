package main

import (
	"time"

	config "github.com/seanbon0611/duly-noted-api-v2/Config"
	controllers "github.com/seanbon0611/duly-noted-api-v2/Controllers"
	"github.com/seanbon0611/duly-noted-api-v2/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//Initiate server
	server := gin.Default()

	//Cors
	server.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000, https://angry-bohr-320bd2.netlify.app/, https://duly-noted-api-cfac6.ondigitalocean.app/duly-noted-api-v-2"},
		AllowMethods:  []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        3 * time.Hour,
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
	server.GET("/api/v1/notes/:id", controllers.GetUserNotes)
	server.POST("/api/v1/notes/create", controllers.CreateNote)

	//Start Server
	server.Run(":3001")

}
