package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	UserID        uuid.UUID `sql:",fk" json:"user_id" gorm:"type:uuid;primaryKey;"`
	ApplicationID int       `sql:",pk" json:"application_id"`
	PostID        int       `sql:",fk" json:"post_id"`
	Message       string    `json:"message"`
	Accepted      bool      `json:"accepted"`
	DateCreated   string    `json:"date_created"  gorm:"autoCreateTime"`
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
