package database

import (
	"database/sql"
	"fmt"

	"main/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (db Database) GetAllUsers() (*models.UserList, error) {
	list := &models.UserList{}
	if err := db.Conn.Order("user_id desc").Find(&list.Users).Error; err != nil {
		return list, err
	}
	return list, nil
}

func (db Database) AddUser(user *models.User) (models.User, error) {
	userOut := models.User{}
	hash, errOne := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if errOne != nil {
		fmt.Println(errOne)
	}
	user.Password = string(hash)
	if err := db.Conn.Create(&user).Error; err != nil {
		return userOut, err
	}
	if err2 := db.Conn.First(&userOut, "email = ?", user.Email).Error; err2 != nil {
		return userOut, err2
	}

	fmt.Println("New user record created with ID and dateJoined: ", userOut.UserID, userOut.DateJoined)
	return userOut, nil
}

func (db Database) GetUserById(userId uuid.UUID) (models.User, error) {
	user := models.User{}
	switch err := db.Conn.First(&user, "user_id = ?", userId).Error; err {
	case sql.ErrNoRows:
		return user, ErrNoMatch
	default:
		return user, err
	}
}

func (db Database) DeleteUser(userId uuid.UUID) error {
	user := &models.User{}
	err := db.Conn.Delete(&user, "user_id = ?", userId).Error
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNoMatch
		default:
			return err
		}
	}
	fmt.Println("User deleted with userID: ", user.UserID)
	return nil
}

func (db Database) UpdateUser(userId uuid.UUID, userData models.User) (models.User, error) {
	user := models.User{}
	hash, errOne := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.MinCost)
	if errOne != nil {
		fmt.Println(errOne)
	}
	user2 := models.User{}
	errTwo := db.Conn.First(&user2, "user_id = ?", userId).Error
	if errTwo != nil {
		if errTwo == sql.ErrNoRows {
			return user, ErrNoMatch
		}
		return user, errTwo
	}
	isMatch := comparePasswords(user2.Password, []byte(userData.Password))
	if isMatch == true {
		hash = []byte(user2.Password)
	}
	err := db.Conn.Model(&user).Updates(models.User{UserID: userId, FirstName: userData.FirstName, LastName: userData.LastName, Email: userData.Email, Password: string(hash), UserCategory: userData.UserCategory, Experience: userData.Experience, Bio: userData.Bio, Preferences: userData.Preferences, City: userData.City, Zipcode: userData.Zipcode, Address: userData.Address}).Error
	if err != nil {
		if err == sql.ErrNoRows {
			return user, ErrNoMatch
		}
		return user, err
	}
	return user, nil
}

func (db Database) UpdateUserProfile(userId uuid.UUID, userData models.User) (models.User, error) {
	user := models.User{}
	user2 := models.User{}
	errTwo := db.Conn.First(&user2, "user_id = ?", userId).Error
	if errTwo != nil {
		if errTwo == sql.ErrNoRows {
			return user, ErrNoMatch
		}
		return user, errTwo
	}
	err := db.Conn.Model(&user).Updates(models.User{UserID: userId, Email: userData.Email, UserCategory: userData.UserCategory, Experience: userData.Experience, Bio: userData.Bio, Preferences: userData.Preferences, City: userData.City, Zipcode: userData.Zipcode, Address: userData.Address}).Error
	if err != nil {
		if err == sql.ErrNoRows {
			return user, ErrNoMatch
		}
		return user, err
	}
	return user, nil
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
