package handler

import (
	"fmt"
	"main/database"
	"net/http"
	"strconv"

	"main/models"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func (r routes) reviews(rg *gin.RouterGroup) {
	rg.POST("/", addReview)
	rg.GET("/:userid", getReviewsByUserId)
	rg.GET("/reviewer/:reviewerid", getReviewsByReviewerId)
	rg.GET("/post/:postid", getReviewsByPostId)
}

func addReview(c *gin.Context) {
	review := &models.Review{}
	r := c.Request
	if err := render.Bind(r, review); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	postOut, err := dbInstance.AddReview(review)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	profOut, err2 := dbInstance.GetUserProfileById(review.UserID)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, "Error fetching user profile: bad request")
		return
	}
	reviewsOut, err3 := dbInstance.GetReviewsByUserId(review.UserID)
	if err3 != nil {
		c.JSON(http.StatusBadRequest, "Error fetching total user reviews: bad request")
		return
	}
	newAverage := ((float64(profOut.Rating) * float64(len(reviewsOut.Review))) + float64(review.Rating)) / float64((len(reviewsOut.Review) + 1))
	profNew := models.Profile{}
	profNew.Rating = newAverage
	finalUpdate, err4 := dbInstance.UpdateUserProfileRating(review.UserID, profNew)
	if err4 != nil {
		c.JSON(http.StatusBadRequest, "Error updating user profile: bad request")
		return
	}
	fmt.Printf(finalUpdate.UserID.String())
	c.JSON(http.StatusOK, postOut)
}

func getReviewsByUserId(c *gin.Context) {
	userID := uuid.MustParse(c.Param("reviewerid"))
	reviews, err := dbInstance.GetReviewsByUserId(userID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, reviews)
}

func getReviewsByReviewerId(c *gin.Context) {
	userID := uuid.MustParse(c.Param("userid"))
	reviews, err := dbInstance.GetReviewsByReviewerId(userID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, reviews)
}

func getReviewsByPostId(c *gin.Context) {
	postIDStr := c.Param("postid")
	postID, errConv := strconv.Atoi(postIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid PostID Integer")
		return
	}
	reviews, err := dbInstance.GetReviewsByPostId(postID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, reviews)
}
