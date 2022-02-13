package database

import (
	"database/sql"
	"fmt"
	"main/models"
)

func (db Database) Signup(user *models.User) bool {
	userOut := models.User{}
	switch err := db.Conn.Select("email").First(&userOut.Email, "email = ?", &user.Email).Error; err {
	case sql.ErrNoRows:
		fmt.Println("Email doesn't already exist; proceed with registration.")
		return true
	default:
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
