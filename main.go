package main

import (
	Controllers "duly_noted/COntrollers"
	Config "duly_noted/Config"

	"github.com/gin-gonic/gin"
)

func main() {
	//Initiate server
	server := gin.Default()

	//Connection to Database
	Config.Init()

	//Routes
	server.GET("/api/v1/users", Controllers.GetUsers)
	server.GET("/api/v1/users/:id", Controllers.GetSingleUser)
	server.POST("/api/v1/users/create", Controllers.CreateUser)
	server.DELETE("/api/v1/users/delete/:id", Controllers.DeleteUser)

	//Start Server
	server.Run(":3001")

}
