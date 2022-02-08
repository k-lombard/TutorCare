package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Post struct {
	UserID          uuid.UUID `sql:",fk" json:"user_id"`
	CaregiverID     uuid.UUID `json:"caregiver_id"`
	PostID          int       `sql:",pk" json:"post_id"`
	Title           string    `json:"title"`
	Tags            string    `json:"tags"`
	CareDescription string    `json:"care_description"`
	CareType        string    `json:"care_type"`
	Completed       bool      `json:"completed"`
	DateOfJob       string    `json:"date_of_job"`
	StartTime       string    `json:"start_time"`
	EndTime         string    `json:"end_time"`
	DatePosted      string    `json:"date_posted"`
}

type PostWithCaregiver struct {
	UserID          uuid.UUID `sql:",fk" json:"user_id"`
	CaregiverID     uuid.UUID `json:"caregiver_id"`
	PostID          int       `sql:",pk" json:"post_id"`
	Title           string    `json:"title"`
	Tags            string    `json:"tags"`
	CareDescription string    `json:"care_description"`
	CareType        string    `json:"care_type"`
	Completed       bool      `json:"completed"`
	DateOfJob       string    `json:"date_of_job"`
	StartTime       string    `json:"start_time"`
	EndTime         string    `json:"end_time"`
	DatePosted      string    `json:"date_posted"`
	Caregiver       User      `json:"caregiver"`
	ApplicationID   int       `json:"application_id"`
	User            User      `json:"user"`
}

type PostWithApplications struct {
	UserID          uuid.UUID             `sql:",fk" json:"user_id"`
	CaregiverID     uuid.UUID             `json:"caregiver_id"`
	PostID          int                   `sql:",pk" json:"post_id"`
	Title           string                `json:"title"`
	Tags            string                `json:"tags"`
	CareDescription string                `json:"care_description"`
	CareType        string                `json:"care_type"`
	Completed       bool                  `json:"completed"`
	DateOfJob       string                `json:"date_of_job"`
	StartTime       string                `json:"start_time"`
	EndTime         string                `json:"end_time"`
	DatePosted      string                `json:"date_posted"`
	Applications    []ApplicationWithUser `json:"applications"`
	Caregiver       User                  `json:"caregiver"`
	ApplicationID   int                   `json:"application_id"`
	User            User                  `json:"user"`
}

type PostList struct {
	Posts []Post `json:"posts"`
}

type PostWithApplicationsList struct {
	Posts []PostWithApplications `json:"posts"`
}

type PostWithCaregiverList struct {
	Posts []PostWithCaregiver `json:"posts"`
}

func (i *Post) Bind(r *http.Request) error {
	if i.CareType == "" || i.DateOfJob == "" || i.StartTime == "" || i.Title == "" || i.CareDescription == "" || i.EndTime == "" || (i.UserID).String() == "" {
		return fmt.Errorf("Care_type, date_of_job, start_time, care_description, end_time, and user_id are required fields.")
	}
	return nil
}

func (i *PostWithApplications) Bind(r *http.Request) error {
	if i.CareType == "" || i.DateOfJob == "" || i.StartTime == "" || i.Title == "" || i.CareDescription == "" || i.EndTime == "" || (i.UserID).String() == "" {
		return fmt.Errorf("Care_type, date_of_job, start_time, care_description, end_time, and user_id are required fields.")
	}
	return nil
}

func (i *PostWithCaregiver) Bind(r *http.Request) error {
	if i.CareType == "" || i.DateOfJob == "" || i.StartTime == "" || i.Title == "" || i.CareDescription == "" || i.EndTime == "" || (i.UserID).String() == "" {
		return fmt.Errorf("Care_type, date_of_job, start_time, care_description, end_time, and user_id are required fields.")
	}
	return nil
}

func (*PostList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*PostWithApplicationsList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*PostWithCaregiverList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Post) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*PostWithApplications) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*PostWithCaregiver) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
