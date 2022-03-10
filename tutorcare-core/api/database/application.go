package database

import (
	"errors"
	"fmt"

	"main/models"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

func (db Database) GetAllApplications() (*models.ApplicationList, error) {
	list := &models.ApplicationList{}
	err := db.Conn.Order("application_id desc").Find(&list.Applications).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func (db Database) AddApplication(app *models.Application) (models.Application, error) {
	appOut := models.Application{}
	postOut := models.Post{}
	errLatest := db.Conn.First(&postOut, "post_id = ?", app.PostID).Error
	if errLatest != nil {
		return appOut, errLatest
	}
	if postOut.UserID != app.UserID {
		err := db.Conn.Create(&app).Error
		if err != nil {
			return appOut, err
		}
		return *app, nil
	} else {
		return appOut, ErrSameUser
	}
}

func (db Database) GetApplicationById(appId int) (models.Application, error) {
	appOut := models.Application{}
	err := db.Conn.First(&appOut, "application_id = ?", appId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return appOut, ErrNoMatch
		}
		return appOut, err
	} else {
		userOut := models.User{}
		err3 := db.Conn.First(&userOut, "user_id = ?", appOut.UserID).Error
		if err3 != nil {
			return appOut, err3
		}
		appOut.User = userOut
		return appOut, nil
	}
}

func (db Database) GetApplicationsByPostId(postId int) (*models.ApplicationList, error) {
	list := &models.ApplicationList{}
	err := db.Conn.Where("post_id = ?", postId).Order("application_id desc").Find(&list.Applications).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func (db Database) GetApplicationsByUserId(userId uuid.UUID) (*models.ApplicationList, error) {
	list := &models.ApplicationList{}
	err := db.Conn.Where("user_id = ?", userId).Order("application_id desc").Find(&list.Applications).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func (db Database) DeleteApplication(applicationId int) error {
	err := db.Conn.Delete(&models.Application{}, applicationId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNoMatch
		}
		return err
	}
	fmt.Println("Application deleted with ApplicationID: ", applicationId)
	return nil
}

func (db Database) UpdateApplication(applicationId int, appData models.Application) (models.Application, error) {
	app := models.Application{}
	app2 := models.Application{}
	errTwo := db.Conn.First(&app2, "application_id = ?", applicationId).Error
	if errTwo != nil {
		if errors.Is(errTwo, gorm.ErrRecordNotFound) {
			return app, ErrNoMatch
		}
		return app, errTwo
	}
	err := db.Conn.Model(&app).Where("application_id = ?", applicationId).Updates(map[string]interface{}{
		"Message":  appData.Message,
		"Accepted": appData.Accepted,
	}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return app, ErrNoMatch
		}
		return app, err
	}
	errOut := db.Conn.First(&app, "application_id = ?", applicationId).Error
	if errOut != nil {
		if errors.Is(errOut, gorm.ErrRecordNotFound) {
			return app, ErrNoMatch
		}
		return app, errOut
	}
	return app, nil
}
