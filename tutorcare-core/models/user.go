package models

import (
	"fmt"
	"net/http"
)

type User struct {
	UserID     int    `json:"user_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	DateJoined string `json:"date_joined"`
}

type UserList struct {
	Users []User `json:"users"`
}

func (i *User) Bind(r *http.Request) error {
	if i.FirstName == "" || i.LastName == "" || i.Email == "" {
		return fmt.Errorf("Firstname, Lastname, Email, and userID are required fields.")
	}
	return nil
}

func (*UserList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
