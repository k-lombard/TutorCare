package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Review struct {
	ReviewID   int       `sql:",pk" json:"review_id" gorm:"primaryKey;default:null;"`
	UserID     uuid.UUID `sql:",fk" json:"user_id" gorm:"type:uuid;default:null;"`
	ReviewerID uuid.UUID `sql:",fk" json:"reviewer_id" gorm:"type:uuid;default:null;"`
	PostID     int       `sql:",fk" json:"post_id" gorm:"default:null;"`
	Rating     int       `json:"rating" gorm:"default:null;"`
	Comment    string    `json:"comment" gorm:"default:null;"`
	ReviewDate string    `json:"review_date" gorm:"default:null;"`
}

type ReviewList struct {
	Review []Review `json:"reviews"`
}

func (i *Review) Bind(r *http.Request) error {
	if (i.UserID).String() == "" || (i.ReviewerID).String() == "" {
		return fmt.Errorf("user_id, reviewer_id, and post_id are required fields.")
	}
	return nil
}

func (*Review) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*ReviewList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
