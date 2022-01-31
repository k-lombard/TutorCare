package database

import (
	"database/sql"
	"fmt"
	"main/models"
)

func (db Database) Signup(user *models.User) bool {
	userOut := models.User{}
	row := db.Conn.QueryRow("SELECT email FROM users WHERE email=$1", &user.Email)
	switch err := row.Scan(&userOut.Email); err {
	case sql.ErrNoRows:
		fmt.Println("Email doesn't already exist; proceed with registration.")
		return true
	default:
		fmt.Println("Email already exists")
		return false
	}
}

func (db Database) ValidateEmail(email string) {
	db.Conn.QueryRow(`UPDATE users SET status=$1 WHERE email=$2;`, true, email)
}

func (db Database) Login(user *models.User) (models.UserWithTokens, bool) {

	userOut := models.UserWithTokens{}

	db.Conn.QueryRow("SELECT * FROM users WHERE email = $1;", user.Email).Scan(&userOut.UserID, &userOut.FirstName, &userOut.LastName, &userOut.Email, &userOut.Password, &userOut.DateJoined, &userOut.Status, &userOut.UserCategory, &userOut.Experience, &userOut.Bio)
	isMatch := comparePasswords(userOut.Password, []byte(user.Password))
	if isMatch == true {
		return userOut, true
	} else {
		return userOut, false
	}

}
