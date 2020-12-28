package Controllers

import (
	Config "duly_noted/Config"
	Models "duly_noted/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//GET
func GetUsers(c *gin.Context) {
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

func GetSingleUser(c *gin.Context) {
	var user Models.User
	if err := Config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

//POST
func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	new_user := Models.User{Username: input.Username, Email: input.Email, Password: input.Password}
	Config.DB.Create(&new_user)
	c.JSON(http.StatusOK, gin.H{"data": new_user})
}

//DELETE
func DeleteUser(c *gin.Context) {
	var user Models.User

	if err := Config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User Not Found"})
		return
	}
	Config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User Successfully Deleted", "data": "true"})
}
