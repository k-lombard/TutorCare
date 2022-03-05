package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type GeolocationPosition struct {
	UserID     uuid.UUID `sql:",fk" json:"user_id" gorm:"type:uuid;default:null;"`
	LocationID int       `sql:",pk" json:"location_id" gorm:"primaryKey;default:null;"`
	Accuracy   float64   `json:"accuracy" gorm:"default:null"`
	Latitude   float64   `json:"latitude" gorm:"default:null"`
	Longitude  float64   `json:"longitude" gorm:"default:null"`
	Timestamp  string    `json:"timestamp" gorm:"default:null"`
	User       User      `json:"user" gorm:"-"`
}

func (GeolocationPosition) TableName() string {
	return "geolocation"
}

type GeolocationPositionList struct {
	GeolocationPositions []GeolocationPosition `json:"geolocation_positions"`
}

func (i *GeolocationPosition) Bind(r *http.Request) error {
	if i.Latitude == 0 || i.Longitude == 0 || (i.UserID).String() == "" {
		return fmt.Errorf("Accuracy, latitude, longitude, and user_id are required fields.")
	}
	return nil
}

func (*GeolocationPositionList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*GeolocationPosition) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
