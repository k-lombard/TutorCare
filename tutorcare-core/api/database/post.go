package database

import (
	"database/sql"
	"fmt"

	"main/models"

	"github.com/google/uuid"
)

func (db Database) GetAllPosts() (*models.PostList, error) {
	list := &models.PostList{}
	err := db.Conn.Order("post_id desc").Find(&list.Posts).Error
	if err != nil {
		return list, err
	}
	for _, post := range list.Posts {
		err := db.Conn.Where("post_id = ?", post.PostID).Select("TO_CHAR(date_of_job :: DATE, 'Mon dd, yyyy')", "TO_CHAR(start_time :: TIME, 'hh12:mi AM')", "TO_CHAR(end_time :: TIME, 'hh12:mi AM')").First(&post.DateOfJob, &post.StartTime, &post.EndTime).Error
		if err != nil {
			return list, err
		}
	}
	return list, nil
}

func (db Database) GetActivePosts() (*models.PostList, error) {
	list := &models.PostList{}
	err := db.Conn.Where("completed = false AND caregiver_id = ?", "00000000-0000-0000-0000-000000000000").Order("post_id desc").Find(&list.Posts).Error
	if err != nil {
		return list, err
	}
	for _, post := range list.Posts {
		err := db.Conn.Where("post_id = ?", post.PostID).Select("TO_CHAR(date_of_job :: DATE, 'Mon dd, yyyy')", "TO_CHAR(start_time :: TIME, 'hh12:mi AM')", "TO_CHAR(end_time :: TIME, 'hh12:mi AM')").First(&post.DateOfJob, &post.StartTime, &post.EndTime).Error
		if err != nil {
			return list, err
		}
	}
	return list, nil
}

func (db Database) GetActivePostsWithCaregiver(userId uuid.UUID) (*models.PostList, error) {
	list := &models.PostList{}
	err := db.Conn.Where("completed = false AND caregiver_id != ? AND user_id = ?", "00000000-0000-0000-0000-000000000000", userId).Order("post_id desc").Find(&list.Posts).Error
	if err != nil {
		return list, err
	}
	for _, post := range list.Posts {
		careUser := models.User{}
		errNew := db.Conn.First(&careUser, "user_id = ?", post.CaregiverID).Error
		if errNew != nil {
			return list, errNew
		}
		errFinal := db.Conn.Where("user_id = ? AND post_id = ?", post.CaregiverID, post.PostID).Select("application_id").First(&post.ApplicationID).Error
		if errFinal != nil {
			return list, errFinal
		}
		post.Caregiver = careUser
	}
	return list, nil
}

func (db Database) GetActivePostsForCaregiverView(caregiverId uuid.UUID) (*models.PostList, error) {
	list := &models.PostList{}
	err := db.Conn.Where("completed = false AND caregiver_id != ? AND caregiver_id = ?", "00000000-0000-0000-0000-000000000000", caregiverId).Order("post_id desc").Find(&list.Posts).Error
	if err != nil {
		return list, err
	}
	for _, post := range list.Posts {
		careUser := models.User{}
		errNew := db.Conn.First(&careUser, "user_id = ?", post.UserID).Error
		if errNew != nil {
			return list, errNew
		}
		post.User = careUser
	}
	return list, nil
}

func (db Database) GetPostsAppliedTo(caregiverId uuid.UUID) (*models.PostList, error) {
	list := &models.PostList{}
	appList := &models.ApplicationList{}
	err := db.Conn.Order("application_id desc").Find(&appList.Applications).Error
	if err != nil {
		return list, err
	}
	for _, application := range appList.Applications {
		post := models.Post{}
		err2 := db.Conn.Where("post_id = ? AND completed = false AND caregiver_id = ?", application.PostID, "00000000-0000-0000-0000-000000000000").Select("*", "TO_CHAR(date_of_job :: DATE, 'Mon dd, yyyy')", "TO_CHAR(start_time :: TIME, 'hh12:mi AM')", "TO_CHAR(end_time :: TIME, 'hh12:mi AM')").First(&post, &post.DateOfJob, &post.StartTime, &post.EndTime).Error
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
	err2 := db.Conn.Where("post_id = ?", post.PostID).Select("*", "TO_CHAR(date_of_job :: DATE, 'Mon dd, yyyy')", "TO_CHAR(start_time :: TIME, 'hh12:mi AM')", "TO_CHAR(end_time :: TIME, 'hh12:mi AM')").First(&postOut, &postOut.DateOfJob, &postOut.StartTime, &postOut.EndTime).Error
	if err2 != nil {
		return postOut, err2
	}
	fmt.Println("New post record created with postID and timestamp: ", postOut.PostID, postOut.DatePosted)
	return postOut, nil
}

func (db Database) GetPostById(postId int) (models.Post, error) {
	postOut := models.Post{}
	switch err := db.Conn.Where("post_id = ? AND caregiver_id != ?", postId, "00000000-0000-0000-0000-000000000000").Select("*", "TO_CHAR(date_of_job :: DATE, 'Mon dd, yyyy')", "TO_CHAR(start_time :: TIME, 'hh12:mi AM')", "TO_CHAR(end_time :: TIME, 'hh12:mi AM')").First(&postOut, &postOut.DateOfJob, &postOut.StartTime, &postOut.EndTime).Error; err {
	case sql.ErrNoRows:
		return postOut, ErrNoMatch
	default:
		newUser := models.User{}
		err4 := db.Conn.First(&newUser, "user_id = ?", postOut.CaregiverID).Error
		if err4 != nil {
			return postOut, err4
		}
		postOut.Caregiver = newUser
		return postOut, nil
	}
}

func (db Database) GetPostsByUserId(userId uuid.UUID) (*models.PostList, error) {
	list := &models.PostList{}
	err := db.Conn.Where("user_id = ? AND caregiver_id != ?", userId, "00000000-0000-0000-0000-000000000000").Order("post_id desc").Find(&list.Posts).Error
	if err != nil {
		return list, err
	}
	for _, post := range list.Posts {
		err := db.Conn.Where("post_id = ?", post.PostID).Select("TO_CHAR(date_of_job :: DATE, 'Mon dd, yyyy')", "TO_CHAR(start_time :: TIME, 'hh12:mi AM')", "TO_CHAR(end_time :: TIME, 'hh12:mi AM')").First(&post.DateOfJob, &post.StartTime, &post.EndTime).Error
		if err != nil {
			return list, err
		}
		list2 := &models.ApplicationList{}
		err2 := db.Conn.Where("post_id = ?", post.PostID).Order("application_id desc").Find(&list2.Applications).Error
		if err2 != nil {
			return list, err2
		}
		for _, application := range list2.Applications {
			newUser := models.User{}
			err4 := db.Conn.First(&newUser, "user_id = ?", application.UserID).Error
			if err4 != nil {
				return list, err4
			}
			application.User = newUser
		}
		newUser2 := models.User{}
		err3 := db.Conn.First(&newUser2, "user_id = ?", post.CaregiverID).Error
		if err3 != nil {
			return list, err3
		}
		post.Caregiver = newUser2
		post.Applications = list2.Applications
	}
	return list, nil
}

func (db Database) DeletePost(postId int) error {
	err := db.Conn.Delete(&models.Post{}, postId).Error
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNoMatch
		default:
			return err
		}
	}
	fmt.Println("Post deleted with PostID: ", postId)
	return nil
}

func (db Database) UpdatePost(postId int, postData models.Post) (models.Post, error) {
	post := models.Post{}

	post2 := models.Post{}
	errTwo := db.Conn.First(&post2, "post_id = ?", postId).Error
	if errTwo != nil {
		if errTwo == sql.ErrNoRows {
			return post, ErrNoMatch
		}
		return post, errTwo
	}
	err := db.Conn.Model(&post).Updates(models.Post{PostID: postId, Title: postData.Title, Tags: postData.Tags, CareDescription: postData.CareDescription, StartTime: postData.StartTime, EndTime: postData.EndTime}).Error
	if err != nil {
		if err == sql.ErrNoRows {
			return post, ErrNoMatch
		}
		return post, err
	}
	return post, nil
}

func (db Database) AddApplicationToPost(postId int, postData models.Post) (models.Post, error) {
	post := models.Post{}
	post2 := models.Post{}
	errTwo := db.Conn.First(&post2, "post_id = ?", postId).Error
	if errTwo != nil {
		if errTwo == sql.ErrNoRows {
			return post, ErrNoMatch
		}
		return post, errTwo
	}
	if postData.CaregiverID != post2.UserID {
		err := db.Conn.Model(&post).Updates(models.Post{PostID: postId, CaregiverID: postData.CaregiverID}).Error
		if err != nil {
			if err == sql.ErrNoRows {
				return post, ErrNoMatch
			}
			return post, err
		}
		return post, nil
	} else {
		return post, ErrSameUser
	}
}
