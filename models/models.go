package models

import (
	"github.com/seanbon0611/duly-noted-api-v2/config"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Email    string `json:"email" gorm:"unique; not null"`
		Password string `gorm:"not null"`
		Notes    []Note
	}

	Note struct {
		gorm.Model
		Content string `json:"content" gorm:"not null"`
		UserID  int
	}
)

//creates new instance of a user into the DB
func (user *User) CreateUserInstance() error {
	result := config.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Gets user inputted password and will convert it to a hash
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

//func that compares to see if the password and hash match
func (user *User) CheckPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}
