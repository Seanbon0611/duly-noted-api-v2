package Routes

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	server := gin.Default()
	v1 := server.Group("/api/v1")
	{
		v1.GET("/users", Controllers.getUsers)
		// 	v1.GET("/users/id", getSingleUser)
		v1.POST("/users/create", createUser)
		// 	v1.DELETE("users/delete/:id", deleteUser)
		// }
		server.Run(":3001")
	}
	return server
}
