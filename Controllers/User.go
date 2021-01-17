package controllers

import (
	"duly_noted/config"
	"duly_noted/models"
	"net/http"

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
		c.JSON(http.StatusNotFound, gin.H{"msg": "error", "error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": user})
}

//POST
func SignupUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "error", "error": err.Error()})
		c.Abort()
		return
	}
	err = user.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"msg": "error", "error": err.Error()})
		c.Abort()
		return
	}
	err = user.CreateUserInstance()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"msg": "error", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": user})
}

//DELETE
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "error", "error": "User Not Found"})
		return
	}
	config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": "User Successfully Deleted"})
}
