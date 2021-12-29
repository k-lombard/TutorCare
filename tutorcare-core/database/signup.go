package database

import (
	"database/sql"
	"fmt"
	"main/models"
)

func (db Database) Signup(user *models.User) (models.User, error) {
	userOut := models.User{}
	row := db.Conn.QueryRow("SELECT email FROM users WHERE email=$1", &user.Email)
	switch err := row.Scan(&userOut.Email); err {
	case sql.ErrNoRows:
		fmt.Println("Email doesn't already exist; proceed with registration.")
		return userOut, nil
	default:
		fmt.Println("Email already exists")
		return userOut, nil
	}
}

func (db Database) Login(user *models.User) (models.User, bool) {

	userOut := models.User{}

	db.Conn.QueryRow("SELECT email, password FROM users WHERE email=$1", user.Email).Scan(&userOut.Email, &userOut.Password)
	// hash, errOne := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	// if errOne != nil {
	// 	fmt.Println(errOne)
	// }
	isMatch := comparePasswords(userOut.Password, []byte(user.Password))
	if isMatch == true {
		return userOut, true
	} else {
		return userOut, false
	}

}
