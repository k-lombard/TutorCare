package handler

import (
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
