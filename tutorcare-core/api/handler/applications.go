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

func (r routes) applications(rg *gin.RouterGroup) {
	rg.GET("/", getAllApplications)
	rg.POST("/", addApplication)
	rg.GET("/user/:userid", getApplicationsByUserId)
	rg.GET("/post/:postid", getApplicationsByPostId)
	rg.GET("/:applicationid", getApplicationById)
	rg.PUT("/:applicationid", updateApplication)
	rg.DELETE("/:applicationid", deleteApplication)
	rg.PUT("/accept/:applicationid", acceptApplication)
}

func addApplication(c *gin.Context) {
	app := &models.Application{}
	r := c.Request
	if err := render.Bind(r, app); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	appOut, err2 := dbInstance.AddApplication(app)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	c.JSON(http.StatusOK, appOut)
}

func getAllApplications(c *gin.Context) {
	apps, err := dbInstance.GetAllApplications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, apps)
}

func getApplicationById(c *gin.Context) {
	applicationIDStr := c.Param("applicationid")
	appID, errConv := strconv.Atoi(applicationIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid ApplicationID Integer")
		return
	}
	app, err := dbInstance.GetApplicationById(appID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, app)
}

func getApplicationsByUserId(c *gin.Context) {
	userID := uuid.MustParse(c.Param("userid"))
	apps, err := dbInstance.GetApplicationsByUserId(userID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, apps)
}

func getApplicationsByPostId(c *gin.Context) {
	postIDStr := c.Param("postid")
	postID, errConv := strconv.Atoi(postIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid PostID Integer")
		return
	}
	apps, err := dbInstance.GetApplicationsByPostId(postID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, apps)
}

func deleteApplication(c *gin.Context) {
	applicationIDStr := c.Param("applicationid")
	appID, errConv := strconv.Atoi(applicationIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid ApplicationID Integer")
		return
	}
	err := dbInstance.DeleteApplication(appID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}
	c.JSON(http.StatusOK, appID)
}
func updateApplication(c *gin.Context) {
	r := c.Request
	applicationIDStr := c.Param("applicationid")
	appID, errConv := strconv.Atoi(applicationIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid ApplicationID Integer")
		return
	}
	appData := models.Application{}
	if err := render.Bind(r, &appData); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	app, err := dbInstance.UpdateApplication(appID, appData)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}
	c.JSON(http.StatusOK, app)
}

func acceptApplication(c *gin.Context) {
	r := c.Request
	applicationIDStr := c.Param("applicationid")
	appID, errConv := strconv.Atoi(applicationIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid ApplicationID Integer")
		return
	}
	appData := models.Application{}
	if err := render.Bind(r, &appData); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	appData.Accepted = true
	app, err := dbInstance.UpdateApplication(appID, appData)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err)
		}
		return
	}
	postData := models.Post{}
	postData.PostID = appData.PostID
	postData.CaregiverID = app.UserID
	postData.CaregiverID = appData.UserID
	post, err2 := dbInstance.AddApplicationToPost(appData.PostID, postData, appData.UserID)
	if err2 != nil {
		if err2 == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			fmt.Println(err2)
			c.JSON(http.StatusInternalServerError, err2)
		}
		return
	}
	c.JSON(http.StatusOK, post)
}
