package database

import (
	"errors"
	"fmt"

	"main/models"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

func (db Database) GetAllPosts() (*models.PostList, error) {
	list := &models.PostList{}
	err := db.Conn.Order("post_id desc").Find(&list.Posts).Error
	if err != nil {
		return list, err
	}
	for i, post := range list.Posts {
		err := db.Conn.Where("post_id = ?", post.PostID).Select("TO_CHAR(start_date :: DATE, 'Mon dd, yyyy') as start_date, TO_CHAR(start_time :: TIME, 'hh12:mi AM') as start_time,TO_CHAR(end_date :: DATE, 'Mon dd, yyyy') as end_date, TO_CHAR(end_time :: TIME, 'hh12:mi AM') as end_time").First(&list.Posts[i]).Error
		if err != nil {
			return list, err
		}
	}
	return list, nil
}

func (db Database) GetActivePosts() (*models.PostList, error) {
	list := &models.PostList{}
	err := db.Conn.Model(&models.Post{}).Where("completed = ? AND (caregiver_id = ? OR caregiver_id IS NULL)", false, uuid.MustParse("00000000-0000-0000-0000-000000000000")).Order("post_id desc").Find(&list.Posts).Error
	if err != nil {
		return list, err
	}
	for i, post := range list.Posts {
		err := db.Conn.Raw("SELECT TO_CHAR(start_date :: DATE, 'Mon dd, yyyy') as start_date, TO_CHAR(start_time :: TIME, 'hh12:mi AM') as start_time,TO_CHAR(end_date :: DATE, 'Mon dd, yyyy') as end_date, TO_CHAR(end_time :: TIME, 'hh12:mi AM') as end_time FROM posts WHERE post_id = ?", post.PostID).Scan(&list.Posts[i]).Error
		if err != nil {
			return list, err
		}
	}
	return list, nil
}

func (db Database) GetActivePostsView(userId uuid.UUID) (*models.PostList, error) {
	list := &models.PostList{}
	err := db.Conn.Where("completed = false AND caregiver_id != ? AND caregiver_id IS NOT NULL AND (user_id = ? OR caregiver_id = ?)", "00000000-0000-0000-0000-000000000000", userId, userId).Order("post_id desc").Find(&list.Posts).Error
	if err != nil {
		return list, err
	}
	for i, post := range list.Posts {
		errDate := db.Conn.Where("post_id = ?", post.PostID).Select("TO_CHAR(start_date :: DATE, 'Mon dd, yyyy') as start_date, TO_CHAR(start_time :: TIME, 'hh12:mi AM') as start_time,TO_CHAR(end_date :: DATE, 'Mon dd, yyyy') as end_date, TO_CHAR(end_time :: TIME, 'hh12:mi AM') as end_time").First(&list.Posts[i]).Error
		if errDate != nil {
			return list, errDate
		}
		careUser := models.User{}
		errNew := db.Conn.First(&careUser, "user_id = ?", post.CaregiverID).Error
		if errNew != nil {
			return list, errNew
		}

		poster := models.User{}
		errPost := db.Conn.First(&poster, "user_id = ?", post.UserID).Error
		if errPost != nil {
			return list, errPost
		}
		list.Posts[i].User = poster
		app := models.Application{}
		errFinal := db.Conn.Where("user_id = ? AND post_id = ?", post.CaregiverID, post.PostID).Select("application_id").First(&app).Error
		if errFinal != nil {
			return list, errFinal
		}
		list.Posts[i].ApplicationID = app.ApplicationID
		list.Posts[i].Caregiver = careUser
	}
	return list, nil
}

func (db Database) GetPostsAppliedTo(caregiverId uuid.UUID) (*models.PostList, error) {
	list := &models.PostList{}
	appList := &models.ApplicationList{}
	err := db.Conn.Where("user_id = ?", caregiverId).Order("application_id desc").Find(&appList.Applications).Error
	if err != nil {
		return list, err
	}
	for _, application := range appList.Applications {
		post := models.Post{}
		err2 := db.Conn.Where("post_id = ? AND completed = false", application.PostID).Select("*").First(&post).Select("TO_CHAR(start_date :: DATE, 'Mon dd, yyyy') as start_date, TO_CHAR(start_time :: TIME, 'hh12:mi AM') as start_time,TO_CHAR(end_date :: DATE, 'Mon dd, yyyy') as end_date, TO_CHAR(end_time :: TIME, 'hh12:mi AM') as end_time").First(&post).Error
		if err2 != nil {
			return list, err2
		}
		userOut := models.User{}
		errFin := db.Conn.First(&userOut, "user_id = ?", post.UserID).Error
		if errFin != nil {
			return list, errFin
		}
		post.User = userOut
		post.Applications = append(post.Applications, application)
		list.Posts = append(list.Posts, post)
	}
	return list, nil
}

func (db Database) AddPost(post *models.Post) (models.Post, error) {
	postOut := models.Post{}
	err := db.Conn.Create(&post).Error
	if err != nil {
		return postOut, err
	}
	err2 := db.Conn.Where("post_id = ?", post.PostID).Select("*").First(&postOut).Select("TO_CHAR(start_date :: DATE, 'Mon dd, yyyy') as start_date, TO_CHAR(start_time :: TIME, 'hh12:mi AM') as start_time,TO_CHAR(end_date :: DATE, 'Mon dd, yyyy') as end_date, TO_CHAR(end_time :: TIME, 'hh12:mi AM') as end_time").First(&postOut).Error
	if err2 != nil {
		return postOut, err2
	}
	fmt.Println("New post record created with postID and timestamp: ", postOut.PostID, postOut.DatePosted)
	return postOut, nil
}

func (db Database) GetPostById(postId int) (models.Post, error) {
	postOut := models.Post{}
	err := db.Conn.Where("post_id = ?", postId).Select("*").First(&postOut).Select("TO_CHAR(start_date :: DATE, 'Mon dd, yyyy') as start_date, TO_CHAR(start_time :: TIME, 'hh12:mi AM') as start_time,TO_CHAR(end_date :: DATE, 'Mon dd, yyyy') as end_date, TO_CHAR(end_time :: TIME, 'hh12:mi AM') as end_time").First(&postOut).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return postOut, ErrNoMatch
		}
		return postOut, err
	} else {
		if postOut.CaregiverID.String() != "00000000-0000-0000-0000-000000000000" {
			newUser := models.User{}
			err4 := db.Conn.First(&newUser, "user_id = ?", postOut.CaregiverID).Error
			if err4 != nil {
				return postOut, err4
			}
			postOut.Caregiver = newUser
		}
		poster := models.User{}
		err5 := db.Conn.First(&poster, "user_id = ?", postOut.UserID).Error
		if err5 != nil {
			return postOut, err5
		}
		postOut.User = poster
		return postOut, nil
	}
}

func (db Database) GetPostsByUserId(userId uuid.UUID) (*models.PostList, error) {
	list := &models.PostList{}
	err := db.Conn.Where("user_id = ?", userId).Order("post_id desc").Find(&list.Posts).Error
	if err != nil {
		return list, err
	}
	for i, post := range list.Posts {
		errDate := db.Conn.Where("post_id = ?", post.PostID).Select("TO_CHAR(start_date :: DATE, 'Mon dd, yyyy') as start_date, TO_CHAR(start_time :: TIME, 'hh12:mi AM') as start_time,TO_CHAR(end_date :: DATE, 'Mon dd, yyyy') as end_date, TO_CHAR(end_time :: TIME, 'hh12:mi AM') as end_time").First(&list.Posts[i]).Error
		if errDate != nil {
			return list, errDate
		}
		list2 := &models.ApplicationList{}
		err2 := db.Conn.Where("post_id = ?", post.PostID).Order("application_id desc").Find(&list2.Applications).Error
		if err2 != nil {
			return list, err2
		}
		for k, application := range list2.Applications {
			newUser := models.User{}
			err4 := db.Conn.First(&newUser, "user_id = ?", application.UserID).Error
			if err4 != nil {
				return list, err4
			}
			list2.Applications[k].User = newUser
		}
		if post.CaregiverID.String() != "00000000-0000-0000-0000-000000000000" {
			newUser2 := models.User{}
			err3 := db.Conn.First(&newUser2, "user_id = ?", post.CaregiverID).Error
			if err3 != nil {
				return list, err3
			}
			list.Posts[i].Caregiver = newUser2
		}
		list.Posts[i].Applications = list2.Applications
	}
	return list, nil
}

func (db Database) GetPostsByUserIdCompleted(userId uuid.UUID) (*models.PostList, error) {
	list := &models.PostList{}
	err := db.Conn.Where("user_id = ? AND completed = true", userId).Order("post_id desc").Find(&list.Posts).Error
	if err != nil {
		return list, err
	}
	for i, post := range list.Posts {
		errDate := db.Conn.Where("post_id = ?", post.PostID).Select("TO_CHAR(start_date :: DATE, 'Mon dd, yyyy') as start_date, TO_CHAR(start_time :: TIME, 'hh12:mi AM') as start_time,TO_CHAR(end_date :: DATE, 'Mon dd, yyyy') as end_date, TO_CHAR(end_time :: TIME, 'hh12:mi AM') as end_time").First(&list.Posts[i]).Error
		if errDate != nil {
			return list, errDate
		}
		list2 := &models.ApplicationList{}
		err2 := db.Conn.Where("post_id = ?", post.PostID).Order("application_id desc").Find(&list2.Applications).Error
		if err2 != nil {
			return list, err2
		}
		for k, application := range list2.Applications {
			newUser := models.User{}
			err4 := db.Conn.First(&newUser, "user_id = ?", application.UserID).Error
			if err4 != nil {
				return list, err4
			}
			list2.Applications[k].User = newUser
		}
		if post.CaregiverID.String() != "00000000-0000-0000-0000-000000000000" {
			newUser2 := models.User{}
			err3 := db.Conn.First(&newUser2, "user_id = ?", post.CaregiverID).Error
			if err3 != nil {
				return list, err3
			}
			list.Posts[i].Caregiver = newUser2
		}
		list.Posts[i].Applications = list2.Applications
	}
	return list, nil
}

func (db Database) DeletePost(postId int) error {
	errApp := db.Conn.Where("post_id = ?", postId).Delete(models.Application{}).Error
	if errApp != nil {
		if errors.Is(errApp, gorm.ErrRecordNotFound) {
			err := db.Conn.Delete(&models.Post{}, postId).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return ErrNoMatch
				} else {
					return err
				}
			}
			fmt.Println("Post deleted with PostID: ", postId)
			return nil
		} else {
			return errApp
		}
	}
	err := db.Conn.Delete(&models.Post{}, postId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNoMatch
		}
		return err
	}
	fmt.Println("Post deleted with PostID: ", postId)
	return nil
}

func (db Database) UpdatePost(postId int, postData models.Post) (models.Post, error) {
	post := models.Post{}

	post2 := models.Post{}
	errTwo := db.Conn.First(&post2, "post_id = ?", postId).Error
	if errTwo != nil {
		if errors.Is(errTwo, gorm.ErrRecordNotFound) {
			return post, ErrNoMatch
		}
		return post, errTwo
	}
	err := db.Conn.Model(&post).Where("post_id = ?", postId).Updates(map[string]interface{}{
		"Title":              postData.Title,
		"Tags":               postData.Tags,
		"CareDescription":    postData.CareDescription,
		"StartDate":          postData.StartDate,
		"StartTime":          postData.StartTime,
		"EndDate":            postData.EndDate,
		"EndTime":            postData.EndTime,
		"PosterCompleted":    postData.PosterCompleted,
		"CaregiverCompleted": postData.CaregiverCompleted,
	}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return post, ErrNoMatch
		}
		return post, err
	}
	if (post2.CaregiverCompleted && postData.PosterCompleted) || (post2.PosterCompleted && postData.CaregiverCompleted) {
		errLast := db.Conn.Model(&post).Where("post_id = ?", postId).Updates(map[string]interface{}{
			"Completed": true,
		}).Error
		if errLast != nil {
			if errors.Is(errLast, gorm.ErrRecordNotFound) {
				return post, ErrNoMatch
			}
			return post, errLast
		}
	}
	post.Completed = true
	return post, nil
}

func (db Database) AddApplicationToPost(postId int, postData models.Post, appUserId uuid.UUID) (models.Post, error) {
	post := models.Post{}
	post2 := models.Post{}
	errTwo := db.Conn.First(&post2, "post_id = ?", postId).Error
	if errTwo != nil {
		if errors.Is(errTwo, gorm.ErrRecordNotFound) {
			return post, ErrNoMatch
		}
		return post, errTwo
	}
	if appUserId != post2.UserID {
		err := db.Conn.Model(&post).Where("post_id = ?", postId).Updates(map[string]interface{}{
			"CaregiverID": postData.CaregiverID,
		}).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return post, ErrNoMatch
			}
			return post, err
		}
		return post, nil
	} else {
		return post, ErrSameUser
	}
}
