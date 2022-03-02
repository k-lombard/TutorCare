package database

import (
	"errors"
	"fmt"

	"main/models"

	"gorm.io/gorm"

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
	if err2 := db.Conn.Where("email = ?", user.Email).First(&userOut).Error; err2 != nil {
		return userOut, err2
	}

	fmt.Println("New user record created with ID and dateJoined: ", userOut.UserID, userOut.DateJoined)

	err3 := db.addUserProfile(userOut.UserID)
	if err3 != nil {
		return userOut, err3
	}

	return userOut, nil
}

func (db Database) addUserProfile(userId uuid.UUID) error {
	userProfileOut := models.Profile{}
	userProfileOut.UserID = userId
	if err := db.Conn.Create(&userProfileOut).Error; err != nil {
		return err
	}
	if err2 := db.Conn.Where("user_id = ?", userId).First(&userProfileOut).Error; err2 != nil {
		return err2
	}
	fmt.Println("New user profile record created with userID and profileID", userProfileOut.UserID, userProfileOut.ProfileID)
	return nil
}

func (db Database) GetUserById(userId uuid.UUID) (models.User, error) {
	user := models.User{}
	err := db.Conn.First(&user, "user_id = ?", userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, ErrNoMatch
		} else {
			return user, err
		}
	}
	return user, nil
}

func (db Database) GetUserProfileById(userId uuid.UUID) (models.Profile, error) {
	userProfile := models.Profile{}
	err := db.Conn.First(&userProfile, "user_id = ?", userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userProfile, ErrNoMatch
		} else {
			return userProfile, err
		}
	}
	return userProfile, nil
}

func (db Database) DeleteUser(userId uuid.UUID) error {
	user := &models.User{}
	err := db.Conn.Delete(&user, "user_id = ?", userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNoMatch
		} else {
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
		if errors.Is(errTwo, gorm.ErrRecordNotFound) {
			return user, ErrNoMatch
		}
		return user, errTwo
	}
	isMatch := comparePasswords(user2.Password, []byte(userData.Password))
	if isMatch == true {
		hash = []byte(user2.Password)
	}
	err := db.Conn.Model(&user).Where("user_id = ?", userId).Updates(models.User{FirstName: userData.FirstName, LastName: userData.LastName, Email: userData.Email, Password: string(hash), UserCategory: userData.UserCategory, Experience: userData.Experience, Bio: userData.Bio, Preferences: userData.Preferences, City: userData.City, Zipcode: userData.Zipcode, Address: userData.Address}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, ErrNoMatch
		}
		return user, err
	}
	errFinal := db.Conn.First(&user, "user_id = ?", userId).Error
	if errFinal != nil {
		if errors.Is(errFinal, gorm.ErrRecordNotFound) {
			return user, ErrNoMatch
		}
		return user, errFinal
	}
	return user, nil
}

func (db Database) UpdateUserData(userId uuid.UUID, userData models.User) (models.User, error) {
	user := models.User{}
	user2 := models.User{}
	errTwo := db.Conn.First(&user2, "user_id = ?", userId).Error
	if errTwo != nil {
		if errors.Is(errTwo, gorm.ErrRecordNotFound) {
			return user, ErrNoMatch
		}
		return user, errTwo
	}
	err := db.Conn.Model(&user).Where("user_id = ?", userId).Updates(models.User{Email: userData.Email, UserCategory: userData.UserCategory, Experience: userData.Experience, Bio: userData.Bio, Preferences: userData.Preferences}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, ErrNoMatch
		}
		return user, err
	}
	errFinal := db.Conn.First(&user, "user_id = ?", userId).Error
	if errFinal != nil {
		if errors.Is(errFinal, gorm.ErrRecordNotFound) {
			return user, ErrNoMatch
		}
		return user, errFinal
	}
	return user, nil
}

func (db Database) UpdateUserProfile(userId uuid.UUID, userData models.Profile) (models.Profile, error) {
	user := models.Profile{}
	user2 := models.Profile{}
	errTwo := db.Conn.First(&user2, "user_id = ?", userId).Error
	if errTwo != nil {
		if errors.Is(errTwo, gorm.ErrRecordNotFound) {
			return user, ErrNoMatch
		}
		return user, errTwo
	}
	err := db.Conn.Model(&user).Where("user_id = ?", userId).Updates(
		models.Profile{
			ProfilePic:    userData.ProfilePic,
			Bio:           userData.Bio,
			BadgeList:     userData.BadgeList,
			Age:           userData.Age,
			Gender:        userData.Gender,
			Language:      userData.Language,
			Experience:    userData.Experience,
			Education:     userData.Education,
			Skills:        userData.Skills,
			ServiceTypes:  userData.ServiceTypes,
			AgeGroups:     userData.AgeGroups,
			Covid19:       userData.Covid19,
			Smoker:        userData.Smoker,
			JobsCompleted: userData.JobsCompleted,
			RateRange:     userData.RateRange,
			Rating:        userData.Rating,
		}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, ErrNoMatch
		}
		return user, err
	}
	errFinal := db.Conn.First(&user, "user_id = ?", userId).Error
	if errFinal != nil {
		if errors.Is(errFinal, gorm.ErrRecordNotFound) {
			return user, ErrNoMatch
		}
		return user, errFinal
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
