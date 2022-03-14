package database

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"main/models"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

func (db Database) GetAllPostCodes() (*models.PostCodeList, error) {
	list := &models.PostCodeList{}
	err := db.Conn.Order("post_id desc").Find(&list.PostCodes).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func (db Database) AddPostCode(postcode *models.PostCode) (models.PostCode, error) {
	postcodeOut := models.PostCode{}
	postcode.Code = generateVerificationCode()
	err := db.Conn.Create(&postcode).Error
	if err != nil {
		return postcodeOut, err
	}
	err2 := db.Conn.Where("post_id = ?", postcode.PostID).Select("*").First(&postcodeOut).Error
	if err2 != nil {
		return postcodeOut, err2
	}
	fmt.Println("New post_code record created with postID and timestamp: ", postcodeOut.PostID, postcodeOut.Timestamp)
	return postcodeOut, nil
}

func (db Database) GetPostCodeByPostId(postId int) (models.PostCode, error) {
	postcodeOut := models.PostCode{}
	err := db.Conn.Where("post_id = ?", postId).Select("*").First(&postcodeOut).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return postcodeOut, ErrNoMatch
		}
		return postcodeOut, err
	} else {
		return postcodeOut, nil
	}
}

func (db Database) GetPostCodesByUserId(userId uuid.UUID) (*models.PostCodeList, error) {
	list := &models.PostCodeList{}
	err := db.Conn.Where("user_id = ?", userId).Order("post_id desc").Find(&list.PostCodes).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func (db Database) DeletePostCode(postId int) error {
	err := db.Conn.Delete(&models.PostCode{}, postId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNoMatch
		}
		return err
	}
	fmt.Println("PostCode deleted with PostID: ", postId)
	return nil
}

func (db Database) UpdatePostCode(postId int, postcodeData models.PostCode) (models.PostCode, error) {
	postcode := models.PostCode{}

	postcode2 := models.PostCode{}
	errTwo := db.Conn.First(&postcode2, "post_id = ?", postId).Error
	if errTwo != nil {
		if errors.Is(errTwo, gorm.ErrRecordNotFound) {
			return postcode, ErrNoMatch
		}
		return postcode, errTwo
	}
	if postcode2.Code == postcodeData.Code {
		err := db.Conn.Model(&postcode).Where("post_id = ?", postId).Updates(models.PostCode{Verified: postcodeData.Verified}).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return postcode, ErrNoMatch
			}
			return postcode, err
		}
		return postcode, nil
	} else {
		return postcode, errors.New("Code is incorrect.")
	}
}

func generateVerificationCode() int {
	rand.Seed(time.Now().UnixNano())
	return 1000 + rand.Intn(999999-1000)
}
