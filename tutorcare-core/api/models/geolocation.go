package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type GeolocationPosition struct {
	UserID     uuid.UUID `sql:",fk" json:"user_id"`
	LocationID int       `sql:",pk" json:"location_id"`
	Accuracy   float32   `json:"accuracy"`
	Latitude   float32   `json:"latitude"`
	Longitude  float32   `json:"longitude"`
	Timestamp  string    `json:"timestamp"`
}

type GeolocationPositionWithUser struct {
	UserID     uuid.UUID `sql:",fk" json:"user_id"`
	LocationID int       `sql:",pk" json:"location_id"`
	Accuracy   float32   `json:"accuracy"`
	Latitude   float32   `json:"latitude"`
	Longitude  float32   `json:"longitude"`
	Timestamp  string    `json:"timestamp"`
	User       User      `json:"user"`
}

type GeolocationPositionList struct {
	GeolocationPositions []GeolocationPosition `json:"geolocation_positions"`
}

type GeolocationPositionWithUserList struct {
	GeolocationPositions []GeolocationPositionWithUser `json:"geolocation_positions"`
}

func (i *GeolocationPosition) Bind(r *http.Request) error {
	if i.Accuracy == 0 || i.Latitude == 0 || i.Longitude == 0 || (i.UserID).String() == "" {
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
