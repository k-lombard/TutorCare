package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID          uuid.UUID     `sql:",fk" json:"user_id"`
	CaregiverID     uuid.UUID     `json:"caregiver_id"`
	PostID          int           `sql:",pk" json:"post_id" gorm:"primaryKey"`
	Title           string        `json:"title"`
	Tags            string        `json:"tags"`
	CareDescription string        `json:"care_description"`
	CareType        string        `json:"care_type"`
	Completed       bool          `json:"completed"`
	DateOfJob       string        `json:"date_of_job"`
	StartTime       string        `json:"start_time"`
	EndTime         string        `json:"end_time"`
	DatePosted      string        `json:"date_posted" gorm:"autoCreateTime"`
	Applications    []Application `json:"applications" gorm:"-"`
	Caregiver       User          `json:"caregiver" gorm:"-"`
	ApplicationID   int           `json:"application_id" gorm:"-"`
	User            User          `json:"user" gorm:"-"`
}

type PostList struct {
	Posts []Post `json:"posts"`
}

func (i *Post) Bind(r *http.Request) error {
	if i.CareType == "" || i.DateOfJob == "" || i.StartTime == "" || i.Title == "" || i.CareDescription == "" || i.EndTime == "" || (i.UserID).String() == "" {
		return fmt.Errorf("Care_type, date_of_job, start_time, care_description, end_time, and user_id are required fields.")
	}
	return nil
}
func (*PostList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Post) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
