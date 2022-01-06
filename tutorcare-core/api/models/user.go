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

type UserWithTokens struct {
	UserID       uuid.UUID `sql:",pk" json:"user_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	DateJoined   string    `json:"date_joined"`
	Status       bool      `json:"status"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

type TokenDetails struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessUuid   string `json:"access_uuid"`
	RefreshUuid  string `json:"refresh_uuid"`
	AtExpires    int64  `json:"at_expires"`
	RtExpires    int64  `json:"rt_expires"`
}

type AccessDetails struct {
	AccessUuid string `json:"access_uuid"`
	UserId     uint64 `json:"user_id"`
}

type UserList struct {
	Users []User `json:"users"`
}

func (i *User) Bind(r *http.Request) error {
	if i.Email == "" || i.Password == "" {
		return fmt.Errorf("Email and password are required fields.")
	}
	return nil
}

func (*UserList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*TokenDetails) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
