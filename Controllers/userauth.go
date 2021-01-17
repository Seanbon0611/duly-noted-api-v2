package controllers

import (
	"duly_noted/auth"
	"duly_noted/config"
	"duly_noted/models"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type (
	LoginPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		Token string `json:"token"`
	}
)

func Login(c *gin.Context) {
	var user models.User
	var payload LoginPayload

	//Checks Token
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "error", "error": err})
		c.Abort()
		return
	}

	//Query to find user where email == user input email param
	result := config.DB.Where("email = ?", payload.Email).First(&user)

	//if credentials are invalid, return error
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "error", "error": "Incorrect credentials"})
		c.Abort()
		return
	}

	//Checking user input password
	err = user.CheckPassword(payload.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "error", "error": "Invalid credentials"})
		c.Abort()
		return
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "thesecretestofkeys",
		Issuer:          "AuthService",
		ExpirationHours: 2,
	}

	//generate uer token
	token, err := jwtWrapper.GenerateToken(user.Email)
	//If there is an error return 500 status code
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "error", "error": "Error signing token"})
		c.Abort()
		return
	}

	tokenResponse := LoginResponse{
		Token: token,
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "token": tokenResponse, "email": user.Email, "notes": user.Notes})

	return
}
