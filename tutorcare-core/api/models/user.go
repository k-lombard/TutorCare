package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID `sql:",pk" json:"user_id" gorm:"type:uuid;primaryKey;default:null;"`
	FirstName    string    `json:"first_name" gorm:"default:null"`
	LastName     string    `json:"last_name" gorm:"default:null"`
	Email        string    `json:"email" gorm:"default:null"`
	Password     string    `json:"password" gorm:"default:null"`
	DateJoined   string    `json:"date_joined" gorm:"default:null"`
	Status       bool      `json:"status" gorm:"default:null"`
	UserCategory string    `json:"user_category" gorm:"default:null"`
	Experience   string    `json:"experience" gorm:"default:null"`
	Bio          string    `json:"bio" gorm:"default:null"`
	Preferences  string    `json:"preferences" gorm:"default:null"`
	Country      string    `json:"country" gorm:"default:null"`
	State        string    `json:"state" gorm:"default:null"`
	City         string    `json:"city" gorm:"default:null"`
	Zipcode      string    `json:"zipcode" gorm:"default:null"`
	Address      string    `json:"address" gorm:"default:null"`
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
