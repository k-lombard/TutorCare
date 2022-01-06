package models

import "github.com/google/uuid"

type GeolocationPosition struct {
	UserID    uuid.UUID `sql:",fk" json:"user_id"`
	Accuracy  string    `json:"accuracy"`
	Latitude  float32   `json:"latitude"`
	Longitude float32   `json:"longitude"`
	Timestamp string    `json:"timestamp"`
}
