package models

import (
	"gorm.io/gorm"
)

func models() {

	type User struct {
		gorm.Model
		Username string
		Email    string
		Password string
	}

	type Note struct {
		gorm.Model
		UserID  uint
		Content string
		User    User
	}
}
