package controllers

import (
	"net/http"

	config "github.com/seanbon0611/duly-noted-api-v2/Config"
	"github.com/seanbon0611/duly-noted-api-v2/models"

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
		c.JSON(http.StatusBadRequest, gin.H{"msg": "error", "error": "JSON error"})
		c.Abort()
		return
	}

	note := models.Note{UserID: input.UserID, Content: input.Content}
	result := config.DB.Create(&note)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "error", "error": "Error Creating Note"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": note})
}

func GetUserNotes(c *gin.Context) {
	var notes []models.Note

	if err := config.DB.Where("user_id = ?", c.Param("id")).Find(&notes).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "error", "error": "No notes found"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": notes})
}
