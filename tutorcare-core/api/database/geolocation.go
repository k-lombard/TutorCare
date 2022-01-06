package database

import (
	"database/sql"
	"fmt"

	"main/models"

	"github.com/google/uuid"
)

func (db Database) GetAllLocations() (*models.GeolocationPositionList, error) {
	list := &models.GeolocationPositionList{}
	rows, err := db.Conn.Query("SELECT * FROM geolocation ORDER BY location_id DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var loc models.GeolocationPosition
		err := rows.Scan(&loc.UserID, &loc.LocationID, &loc.Accuracy, &loc.Latitude, &loc.Longitude, &loc.Timestamp)
		if err != nil {
			return list, err
		}
		list.GeolocationPositions = append(list.GeolocationPositions, loc)
	}
	return list, nil
}
func (db Database) AddGeolocationPosition(loc *models.GeolocationPosition) (models.GeolocationPosition, error) {
	sqlStatement := `INSERT INTO geolocation (user_id, accuracy, latitude, longitude) VALUES ($1, $2, $3, $4) RETURNING location_id, timestamp;`
	var id int
	var timestamp string
	geolocationPositionOut := models.GeolocationPosition{}

	err := db.Conn.QueryRow(sqlStatement, &loc.UserID, &loc.Accuracy, &loc.Latitude, &loc.Longitude).Scan(&id, &timestamp)

	if err != nil {
		return geolocationPositionOut, err
	}
	err2 := db.Conn.QueryRow(`SELECT * FROM geolocation WHERE user_id = $1;`, &loc.UserID).Scan(&geolocationPositionOut.UserID, &geolocationPositionOut.LocationID, &geolocationPositionOut.Accuracy, &geolocationPositionOut.Latitude, &geolocationPositionOut.Longitude, &geolocationPositionOut.Timestamp)
	if err2 != nil {
		return geolocationPositionOut, err2
	}
	fmt.Println("New geolocation_position record created with locationID and timestamp: ", id, timestamp)
	return geolocationPositionOut, nil
}

func (db Database) GetGeolocationPositionByUserId(userId uuid.UUID) (models.GeolocationPosition, error) {
	loc := models.GeolocationPosition{}
	query := `SELECT * FROM geolocation WHERE user_id = $1;`
	row := db.Conn.QueryRow(query, userId)
	switch err := row.Scan(&loc.UserID, &loc.LocationID, &loc.Accuracy, &loc.Latitude, &loc.Longitude, &loc.Timestamp); err {
	case sql.ErrNoRows:
		return loc, ErrNoMatch
	default:
		return loc, err
	}
}
func (db Database) DeleteGeolocationPosition(userId uuid.UUID) error {
	var id uuid.UUID
	query := `DELETE FROM geolocation WHERE user_id = $1 RETURNING user_id;`
	err := db.Conn.QueryRow(query, userId).Scan(&id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNoMatch
		default:
			return err
		}
	}
	fmt.Println("GeolocationPosition deleted with userID: ", id)
	return nil
}
func (db Database) UpdateGeolocationPosition(userId uuid.UUID, locData models.GeolocationPosition) (models.GeolocationPosition, error) {
	loc := models.GeolocationPosition{}
	query := `UPDATE geolocation SET accuracy=$1, latitude=$2, longitude=$3 WHERE user_id=$4 RETURNING user_id, location_id, accuracy, latitude, longitude, timestamp;`

	query2 := `SELECT * FROM geolocation WHERE user_id = $1;`
	loc2 := models.GeolocationPosition{}
	errTwo := db.Conn.QueryRow(query2, userId).Scan(&loc2.UserID, &loc2.LocationID, &loc2.Accuracy, &loc2.Latitude, &loc2.Longitude, &loc2.Timestamp)
	if errTwo != nil {
		if errTwo == sql.ErrNoRows {
			return loc, ErrNoMatch
		}
		return loc, errTwo
	}
	err := db.Conn.QueryRow(query, locData.Accuracy, locData.Latitude, locData.Longitude, userId).Scan(&loc.UserID, &loc.LocationID, &loc.Accuracy, &loc.Latitude, &loc.Longitude, &loc.Timestamp)
	if err != nil {
		if err == sql.ErrNoRows {
			return loc, ErrNoMatch
		}
		return loc, err
	}
	return loc, nil
}
