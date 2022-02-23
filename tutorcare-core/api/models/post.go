package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Post struct {
	UserID             uuid.UUID     `sql:",fk" json:"user_id" gorm:"type:uuid;default:null;"`
	CaregiverID        uuid.UUID     `json:"caregiver_id" gorm:"type:uuid;default:null;"`
	PostID             int           `sql:",pk" json:"post_id" gorm:"primaryKey;default:null;"`
	Title              string        `json:"title" gorm:"default:null"`
	Tags               string        `json:"tags" gorm:"default:null"`
	CareDescription    string        `json:"care_description" gorm:"default:null"`
	CareType           string        `json:"care_type" gorm:"default:null"`
	Completed          bool          `json:"completed" gorm:"default:null"`
	StartDate          string        `json:"start_date" gorm:"default:null"`
	StartTime          string        `json:"start_time" gorm:"default:null"`
	EndDate            string        `json:"end_date" gorm:"default:null"`
	EndTime            string        `json:"end_time" gorm:"default:null"`
	DatePosted         string        `json:"date_posted" gorm:"default:null"`
	PosterCompleted    bool          `json:"poster_completed" gorm:"default:null"`
	CaregiverCompleted bool          `json:"caregiver_completed" gorm:"default:null"`
	Applications       []Application `json:"applications" gorm:"-"`
	Caregiver          User          `json:"caregiver" gorm:"-"`
	ApplicationID      int           `json:"application_id" gorm:"-"`
	User               User          `json:"user" gorm:"-"`
}

type PostList struct {
	Posts []Post `json:"posts"`
}

func (i *Post) Bind(r *http.Request) error {
	if i.CareType == "" || i.StartDate == "" || i.StartTime == "" || i.Title == "" || i.CareDescription == "" || i.EndDate == "" || i.EndTime == "" || (i.UserID).String() == "" {
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
