package database

import (
	"main/models"

	"github.com/google/uuid"
)

func (db Database) AddReview(review *models.Review) (models.Review, error) {
	reviewOut := models.Review{}

	if err := db.Conn.Create(&review).Error; err != nil {
		return reviewOut, err
	}
	if err2 := db.Conn.Where("review_id = ?", review.ReviewID).First(&reviewOut).Error; err2 != nil {
		return reviewOut, err2
	}
	return reviewOut, nil
}

func (db Database) GetReviewsByUserId(userId uuid.UUID) (*models.ReviewList, error) {
	list := &models.ReviewList{}
	err := db.Conn.Where("user_id = ?", userId).Order("review_id desc").Find(&list.Review).Error
	if err != nil {
		return list, err
	}
	for i, review := range list.Review {
		errDate := db.Conn.Where("review_id = ?", review.ReviewID).Select("*").First(&list.Review[i]).Error
		if errDate != nil {
			return list, errDate
		}
	}
	return list, nil
}

func (db Database) GetReviewsByReviewerId(reviewerId uuid.UUID) (*models.ReviewList, error) {
	list := &models.ReviewList{}
	err := db.Conn.Where("reviewer_id = ?", reviewerId).Order("review_id desc").Find(&list.Review).Error
	if err != nil {
		return list, err
	}
	for i, review := range list.Review {
		errDate := db.Conn.Where("review_id = ?", review.ReviewID).Select("*").First(&list.Review[i]).Error
		if errDate != nil {
			return list, errDate
		}
	}
	return list, nil
}

func (db Database) GetReviewsByPostId(postId int) (*models.ReviewList, error) {
	list := &models.ReviewList{}
	err := db.Conn.Where("post_id = ?", postId).Order("review_id desc").Find(&list.Review).Error
	if err != nil {
		return list, err
	}
	for i, review := range list.Review {
		errDate := db.Conn.Where("review_id = ?", review.ReviewID).Select("*").First(&list.Review[i]).Error
		if errDate != nil {
			return list, errDate
		}
	}
	return list, nil
}
