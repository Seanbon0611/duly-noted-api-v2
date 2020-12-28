package main

import (
	Controllers "duly_noted/COntrollers"
	Config "duly_noted/Config"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	Config.Init()
	server.Use(func(c *gin.Context) {
		c.Set("db", Config.DB)
		c.Next()
	})
	server.GET("/api/v1/users", Controllers.GetUsers)
	server.GET("/api/v1/users/:id", Controllers.GetSingleUser)
	server.POST("/api/v1/users/create", Controllers.CreateUser)
	// 	v1.DELETE("users/delete/:id", deleteUser)
	// }
	server.Run(":3001")
	// Routes.Router()
}
