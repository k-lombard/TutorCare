package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type User struct {
	UserID     uuid.UUID `sql:",pk" json:"user_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	DateJoined string    `json:"date_joined"`
	Status     bool      `json:"status"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
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
