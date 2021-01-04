package models

import (
	"duly_noted/config"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Username string `json:"username" gorm:"varchar(27)unique; not null"`
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

func (user *User) CreateUserInstance() error {
	result := config.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}
