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

func (r routes) posts(rg *gin.RouterGroup) {
	rg.GET("/", getAllPosts)
	rg.GET("/active", getActivePosts)
	rg.GET("/active-jobs/:userid", getActivePostsView)
	// rg.GET("/active-jobs/caregiver/:caregiverid", getActivePostsForCaregiverView)
	rg.GET("/applied-to/:caregiverid", getPostsAppliedTo)
	rg.POST("/", addPost)
	rg.GET("/user/:userid", getPostsByUserId)
	rg.GET("/user/completed/:userid", getPostsByUserIdCompleted)
	rg.GET("/:postid", getPostById)
	rg.PUT("/:postid", updatePost)
	rg.PUT("/completed/:postid", updatePostCompleted)
	rg.DELETE("/:postid", deletePost)
}

func addPost(c *gin.Context) {
	post := &models.Post{}
	r := c.Request
	if err := render.Bind(r, post); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	postOut, err := dbInstance.AddPost(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	postCode := &models.PostCode{}
	postCode.PostID = post.PostID
	_, err2 := dbInstance.AddPostCode(postCode)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad Request")
		return
	}
	c.JSON(http.StatusOK, postOut)
}

func getAllPosts(c *gin.Context) {
	posts, err := dbInstance.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, posts)
}

func getActivePosts(c *gin.Context) {
	posts, err := dbInstance.GetActivePosts()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, posts)
}

func getActivePostsView(c *gin.Context) {
	userID := uuid.MustParse(c.Param("userid"))
	posts, err := dbInstance.GetActivePostsView(userID)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, posts)
}

func getPostsAppliedTo(c *gin.Context) {
	caregiverID := uuid.MustParse(c.Param("caregiverid"))
	posts, err := dbInstance.GetPostsAppliedTo(caregiverID)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, posts)
}

func getPostById(c *gin.Context) {
	postIDStr := c.Param("postid")
	postID, errConv := strconv.Atoi(postIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid PostID Integer")
		return
	}
	post, err := dbInstance.GetPostById(postID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, post)
}

func getPostsByUserId(c *gin.Context) {
	userID := uuid.MustParse(c.Param("userid"))
	posts, err := dbInstance.GetPostsByUserId(userID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, posts)
}

func getPostsByUserIdCompleted(c *gin.Context) {
	userID := uuid.MustParse(c.Param("userid"))
	posts, err := dbInstance.GetPostsByUserIdCompleted(userID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, posts)
}

func deletePost(c *gin.Context) {
	postIDStr := c.Param("postid")
	postID, errConv := strconv.Atoi(postIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid PostID Integer")
		return
	}
	err := dbInstance.DeletePost(postID)
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
func updatePost(c *gin.Context) {
	r := c.Request
	postIDStr := c.Param("postid")
	postID, errConv := strconv.Atoi(postIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid PostID Integer")
		return
	}
	postData := models.Post{}
	if err := render.Bind(r, &postData); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	post, err := dbInstance.UpdatePost(postID, postData)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusInternalServerError, "Internal server error")
		}
		return
	}
	c.JSON(http.StatusOK, post)
}

func updatePostCompleted(c *gin.Context) {
	r := c.Request
	postIDStr := c.Param("postid")
	postID, errConv := strconv.Atoi(postIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid PostID Integer")
		return
	}
	postData := models.Post{}
	if err := render.Bind(r, &postData); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	post, err := dbInstance.UpdatePostCompleted(postID, postData)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusInternalServerError, "Internal server error")
		}
		return
	}
	c.JSON(http.StatusOK, post)
}
