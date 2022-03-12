package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Profile struct {
	UserID        uuid.UUID `sql:",fk" json:"user_id" gorm:"type:uuid;default:null;"`
	ProfileID     int       `sql:",pk" json:"profile_id" gorm:"primaryKey;default:null;"`
	ProfilePic    bool      `json:"profile_pic" gorm:"default:null"`
	Bio           string    `json:"bio" gorm:"default:null"`
	BadgeList     string    `json:"badge_list" gorm:"default:null"`
	Age           int       `json:"age" gorm:"default:null"`
	Gender        string    `json:"gender" gorm:"default:null"`
	Language      string    `json:"language" gorm:"default:null"`
	Experience    string    `json:"experience" gorm:"default:null"`
	Education     string    `json:"education" gorm:"default:null"`
	Skills        string    `json:"skills" gorm:"default:null"`
	ServiceTypes  string    `json:"service_types" gorm:"default:null"`
	AgeGroups     string    `json:"age_groups" gorm:"default:null"`
	Covid19       bool      `json:"covid19" gorm:"default:null"`
	Cpr           bool      `json:"cpr" gorm:"default:null"`
	FirstAid      bool      `json:"first_aid" gorm:"default:null"`
	Smoker        bool      `json:"smoker" gorm:"default:null"`
	JobsCompleted int       `json:"jobs_completed" gorm:"default:null"`
	RateRange     string    `json:"rate_range" gorm:"default:null"`
	Rating        float64   `json:"rating" gorm:"default:null"`
}

func (i *Profile) Bind(r *http.Request) error {
	if (i.UserID).String() == "" {
		return fmt.Errorf("user_id are required fields.")
	}
	return nil
}

func (*Profile) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
