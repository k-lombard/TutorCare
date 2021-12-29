package database

import (
	"database/sql"
	"fmt"

	"main/models"

	"github.com/google/uuid"
)

func (db Database) GetAllUsers() (*models.UserList, error) {
	list := &models.UserList{}
	rows, err := db.Conn.Query("SELECT * FROM users ORDER BY user_id DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.DateJoined, &user.Status)
		if err != nil {
			return list, err
		}
		list.Users = append(list.Users, user)
	}
	return list, nil
}
func (db Database) AddUser(user *models.User) error {
	sqlStatement := `INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING user_id, date_joined, status;`
	var id uuid.UUID
	var status bool
	var dateJoined string
	err := db.Conn.QueryRow(sqlStatement, &user.FirstName, &user.LastName, &user.Email, &user.Password).Scan(&id, &dateJoined, &status)
	if err != nil {
		return err
	}
	fmt.Println("New user record created with ID and dateJoined: ", id, dateJoined)
	return nil
}

func (db Database) GetUserById(userId uuid.UUID) (models.User, error) {
	user := models.User{}
	query := `SELECT * FROM users WHERE user_id = $1;`
	row := db.Conn.QueryRow(query, userId)
	switch err := row.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.DateJoined, &user.Status); err {
	case sql.ErrNoRows:
		return user, ErrNoMatch
	default:
		return user, err
	}
}
func (db Database) DeleteUser(userId uuid.UUID) error {
	var id uuid.UUID
	query := `DELETE FROM users WHERE user_id = $1 RETURNING user_id;`
	err := db.Conn.QueryRow(query, userId).Scan(&id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNoMatch
		default:
			return err
		}
	}
	fmt.Println("User deleted with userID: ", id)
	return nil
}
func (db Database) UpdateUser(userId uuid.UUID, userData models.User) (models.User, error) {
	user := models.User{}
	query := `UPDATE users SET first_name=$1, last_name=$2, email=$3, password=$4 WHERE user_id=$5 RETURNING user_id, first_name, last_name, email, password, date_joined, status;`
	err := db.Conn.QueryRow(query, userData.FirstName, userData.LastName, userData.Email, userData.Password, userId).Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.DateJoined, &user.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, ErrNoMatch
		}
		return user, err
	}
	return user, nil
}
