package models

import (
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Username string `json:"username" binding:"min=3, max=25" gorm:"varchar(27)unique; not null"`
		Email    string `json:"email"gorm:"unique; not null"`
		Password string `gorm:"not null`
	}

	Note struct {
		gorm.Model
		UserID  uint   `json:"user_id"`
		Content string `json:"content`
		User    User
	}
)
