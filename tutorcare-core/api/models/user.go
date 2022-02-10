package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID       uuid.UUID `sql:",pk" json:"user_id" gorm:"primaryKey"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	DateJoined   string    `json:"date_joined" gorm:"autoCreateTime"`
	Status       bool      `json:"status"`
	UserCategory string    `json:"user_category"`
	Experience   string    `json:"experience"`
	Bio          string    `json:"bio"`
	Preferences  string    `json:"preferences"`
	Country      string    `json:"country"`
	State        string    `json:"state"`
	City         string    `json:"city"`
	Zipcode      string    `json:"zipcode"`
	Address      string    `json:"address"`
	AccessToken  string    `json:"access_token" gorm:"-"`
	RefreshToken string    `json:"refresh_token" gorm:"-"`
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
	UserId     string `json:"user_id"`
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
