package database

import (
	"errors"
	"fmt"

	"main/models"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

func (db Database) GetAllLocations() (*models.GeolocationPositionList, error) {
	list := &models.GeolocationPositionList{}
	err := db.Conn.Order("location_id desc").Find(&list.GeolocationPositions).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func (db Database) GetCaregiverLocations() (*models.GeolocationPositionList, error) {
	list := &models.GeolocationPositionList{}
	err := db.Conn.Order("location_id desc").Find(&list.GeolocationPositions).Error
	if err != nil {
		return list, err
	}
	for i, geopos := range list.GeolocationPositions {

		us := models.User{}
		err2 := db.Conn.First(&us, "user_id = ?", geopos.UserID).Error
		if err2 != nil {
			return list, err2
		}
		if us.UserCategory == "caregiver" || us.UserCategory == "both" {
			list.GeolocationPositions[i].User = us
		}
	}
	return list, nil
}

func (db Database) AddGeolocationPosition(loc *models.GeolocationPosition) (models.GeolocationPosition, error) {
	geolocationPositionTemp := models.GeolocationPosition{}
	err := db.Conn.First(&geolocationPositionTemp, "user_id = ?", loc.UserID).Error
	if err != nil {
		geolocationPositionOut := models.GeolocationPosition{}
		if errors.Is(err, gorm.ErrRecordNotFound) {

			errNew := db.Conn.Create(&loc).Error
			if errNew != nil {
				return geolocationPositionOut, errNew
			}
			err2 := db.Conn.First(&geolocationPositionOut, "user_id = ?", loc.UserID).Error
			if err2 != nil {
				return geolocationPositionOut, err2
			}
			fmt.Println("New geolocation_position record created with locationID and timestamp: ", geolocationPositionOut.LocationID, geolocationPositionOut.Timestamp)
			return geolocationPositionOut, nil
		}
		return geolocationPositionOut, err
	} else {
		geolocationPositionOut := models.GeolocationPosition{}

		err := db.Conn.Model(&geolocationPositionOut).Where("user_id = ?", loc.UserID).Updates(map[string]interface{}{
			"Accuracy":  loc.Accuracy,
			"Latitude":  loc.Latitude,
			"Longitude": loc.Longitude,
		}).Error
		if err != nil {
			return geolocationPositionOut, err
		}
		fmt.Println("Geolocation_position record updated with locationID and timestamp: ", geolocationPositionOut.LocationID, geolocationPositionOut.Timestamp)
		return geolocationPositionOut, nil
	}
}

func (db Database) GetGeolocationPositionByUserId(userId uuid.UUID) (models.GeolocationPosition, error) {
	loc := models.GeolocationPosition{}
	err := db.Conn.First(&loc, "user_id = ?", userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return loc, ErrNoMatch
		}
		return loc, err
	} else {
		return loc, nil
	}
}

func (db Database) DeleteGeolocationPosition(userId uuid.UUID) error {
	out := models.GeolocationPosition{}
	err := db.Conn.Where("user_id = ?", userId).Delete(&out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNoMatch
		}
		return err
	}
	fmt.Println("GeolocationPosition deleted with userID: ", out.LocationID)
	return nil
}

func (db Database) UpdateGeolocationPosition(userId uuid.UUID, locData models.GeolocationPosition) (models.GeolocationPosition, error) {
	loc := models.GeolocationPosition{}
	loc2 := models.GeolocationPosition{}
	errTwo := db.Conn.First(&loc2, "user_id = ?", userId).Error
	if errTwo != nil {
		if errors.Is(errTwo, gorm.ErrRecordNotFound) {
			return loc, ErrNoMatch
		}
		return loc, errTwo
	}
	err := db.Conn.Model(&loc).Where("user_id = ?", userId).Updates(map[string]interface{}{
		"Accuracy":  locData.Accuracy,
		"Latitude":  locData.Latitude,
		"Longitude": locData.Longitude,
	}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return loc, ErrNoMatch
		}
		return loc, err
	}
	errOut := db.Conn.First(&loc, "user_id = ?", userId).Error
	if errOut != nil {
		if errors.Is(errOut, gorm.ErrRecordNotFound) {
			return loc, ErrNoMatch
		}
		return loc, errOut
	}
	return loc, nil
}
