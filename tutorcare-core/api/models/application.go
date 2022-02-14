package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Application struct {
	UserID        uuid.UUID `sql:",fk" json:"user_id" gorm:"type:uuid;default:null;"`
	ApplicationID int       `sql:",pk" json:"application_id" gorm:"primaryKey;default:null;"`
	PostID        int       `sql:",fk" json:"post_id" gorm:"default:null"`
	Message       string    `json:"message" gorm:"default:null"`
	Accepted      bool      `json:"accepted" gorm:"default:null"`
	DateCreated   string    `json:"date_created" gorm:"default:null"`
	User          User      `json:"user" gorm:"-"`
}

type ApplicationList struct {
	Applications []Application `json:"applications"`
}

func (i *Application) Bind(r *http.Request) error {
	if i.PostID == 0 || i.Message == "" || (i.UserID).String() == "" {
		return fmt.Errorf("Post_id, message, and user_id are required fields.")
	}
	return nil
}

func (*ApplicationList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Application) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
