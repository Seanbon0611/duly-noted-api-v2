package main

import (
	"net/http"

	Config "duly_noted/Config"
	Models "duly_noted/Models"

	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	var users []Models.User
	Config.DB.Find(&users)

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
	Config.DB.Save(user)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User Successfully Created"})
}

func main() {

	server := gin.Default()
	Config.Init()
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
