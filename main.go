package main

import (
	"time"

	config "github.com/seanbon0611/duly-noted-api-v2/Config"
	controllers "github.com/seanbon0611/duly-noted-api-v2/Controllers"
	"github.com/seanbon0611/duly-noted-api-v2/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {

	tracer.Start(
		tracer.WithEnv("development"),
		tracer.WithService("duly-noted"),
		tracer.WithServiceVersion("abc123"),
	)
	defer tracer.Stop()
	//Initiate server
	server := gin.Default()
	server.Use(gintrace.Middleware("duly-noted"))

	//Cors
	server.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"https://angry-bohr-320bd2.netlify.app", "http://localhost:3000"},
		AllowMethods:  []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
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
	server.DELETE("/api/v1/notes/delete/:id", controllers.DeleteNote)

	//Start Server
	server.Run(":3001")

}
