package Controllers

import (
	"net/http"

	config "github.com/seanbon0611/duly-noted-api-v2/Config"
	"github.com/seanbon0611/duly-noted-api-v2/models"

	"github.com/gin-gonic/gin"
)

//GET
func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"msg": "error", "error": "No users found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": users})
}

func GetSingleUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "error", "error": "User not found "})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": user})
}

//POST
func SignupUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "error", "error": "Error with JSON"})
		c.Abort()
		return
	}
	err = user.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"msg": "error", "error": "Error could not signup"})
		c.Abort()
		return
	}
	err = user.CreateUserInstance()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"msg": "error", "error": "Error registering user"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": user})
}

// //UPDATE
// func UpdateUser(c *gin.Context) {
// 	var user models.User

// 	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"msg": "error", "error": "Cannot find user"})
// 		c.Abort()
// 		return
// 	}
// }

//DELETE
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "error", "error": "User Not Found"})
		c.Abort()
		return
	}
	config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": "User Successfully Deleted"})
}
