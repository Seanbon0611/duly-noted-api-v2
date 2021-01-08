package controllers

import (
	"duly_noted/config"
	"duly_noted/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateNoteInput struct {
	UserID  int    `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

//POST
func CreateNote(c *gin.Context) {
	var input CreateNoteInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	note := models.Note{UserID: input.UserID, Content: input.Content}
	result := config.DB.Create(&note)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Creating Note"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": note})
}
