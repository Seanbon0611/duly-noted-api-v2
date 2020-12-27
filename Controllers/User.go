package Controllers

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
