package models

type SignupResponse struct {
	User   User     `json:"user"`
	Coords Location `json:"coords"`
}

type Location struct {
	Latitude  float64
	Longitude float64
}
