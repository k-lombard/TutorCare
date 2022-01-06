package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type GeolocationPosition struct {
	UserID     uuid.UUID `sql:",fk" json:"user_id"`
	LocationID int       `json:"location_id"`
	Accuracy   string    `json:"accuracy"`
	Latitude   float32   `json:"latitude"`
	Longitude  float32   `json:"longitude"`
	Timestamp  string    `json:"timestamp"`
}

type GeolocationPositionList struct {
	GeolocationPositions []GeolocationPosition `json:"geolocation_positions"`
}

func (i *GeolocationPosition) Bind(r *http.Request) error {
	if i.Latitude == 0 || i.Longitude == 0 || (i.UserID).String() == "" {
		return fmt.Errorf("Latitude, longitude, and user_id are required fields.")
	}
	return nil
}

func (*GeolocationPositionList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*GeolocationPosition) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
