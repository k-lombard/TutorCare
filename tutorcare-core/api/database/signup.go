package database

import (
	"errors"
	"fmt"
	"main/models"

	"gorm.io/gorm"
)

func (db Database) Signup(user *models.User) bool {
	userOut := models.User{}
	err := db.Conn.Select("email").First(&userOut, "email = ?", user.Email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Email doesn't already exist; proceed with registration.")
			return true
		}
		return false
	} else {
		fmt.Println("Email already exists")
		return false
	}
}

func (db Database) ValidateEmail(email string) {
	db.Conn.Model(&models.User{}).Where("email = ?", email).Updates(models.User{Status: true})
}

func (db Database) Login(user *models.User) (models.User, bool) {

	userOut := models.User{}

	db.Conn.First(&userOut, "email = ?", user.Email)
	isMatch := comparePasswords(userOut.Password, []byte(user.Password))
	if isMatch == true {
		return userOut, true
	} else {
		return userOut, false
	}

}
