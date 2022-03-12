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

func (r routes) post_codes(rg *gin.RouterGroup) {
	rg.GET("/", getAllPostCodes)
	rg.GET("/:postid", getPostCodeByPostId)
	rg.GET("/user/:userid", getPostCodesByUserId)
	rg.POST("/", addPostCode)
	rg.PUT("/:postid", updatePostCode)
	rg.DELETE("/:postid", deletePostCode)
}

func addPostCode(c *gin.Context) {
	postcode := &models.PostCode{}
	r := c.Request
	if err := render.Bind(r, postcode); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	postcodeOut, err := dbInstance.AddPostCode(postcode)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	c.JSON(http.StatusOK, postcodeOut)
}

func getAllPostCodes(c *gin.Context) {
	postcodes, err := dbInstance.GetAllPostCodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, postcodes)
}

func getPostCodesByUserId(c *gin.Context) {
	userID := uuid.MustParse(c.Param("userid"))
	postcodes, err := dbInstance.GetPostCodesByUserId(userID)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, postcodes)
}

func getPostCodeByPostId(c *gin.Context) {
	postIDStr := c.Param("postid")
	postID, errConv := strconv.Atoi(postIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid PostID Integer")
		return
	}
	postcodes, err := dbInstance.GetPostCodeByPostId(postID)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, postcodes)
}

func deletePostCode(c *gin.Context) {
	postIDStr := c.Param("postid")
	postID, errConv := strconv.Atoi(postIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid PostID Integer")
		return
	}
	err := dbInstance.DeletePostCode(postID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}
	c.JSON(http.StatusOK, postID)
}
func updatePostCode(c *gin.Context) {
	r := c.Request
	postIDStr := c.Param("postid")
	postID, errConv := strconv.Atoi(postIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid PostID Integer")
		return
	}
	postData := models.PostCode{}
	if err := render.Bind(r, &postData); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	postcode, err := dbInstance.UpdatePostCode(postID, postData)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusInternalServerError, "Internal server error")
		}
		return
	}
	c.JSON(http.StatusOK, postcode)
}
