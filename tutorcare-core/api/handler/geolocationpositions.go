package handler

import (
	"fmt"
	"main/database"
	"net/http"

	"main/models"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func (r routes) geolocationpositions(rg *gin.RouterGroup) {
	rg.GET("/", getAllGeolocationPositions)
	rg.GET("/caregivers", getCaregiverGeolocationPositions)
	rg.POST("/", addGeolocationPosition)
	rg.GET("/:userid", getGeolocationPositionByUserId)
	rg.PUT("/:userid", updateGeolocationPosition)
	rg.DELETE("/:userid", deleteGeolocationPosition)
}

func addGeolocationPosition(c *gin.Context) {
	loc := &models.GeolocationPosition{}
	r := c.Request
	if err := render.Bind(r, loc); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	geolocationPositionOut, err := dbInstance.AddGeolocationPosition(loc)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	c.JSON(http.StatusOK, geolocationPositionOut)
}

func getAllGeolocationPositions(c *gin.Context) {
	locs, err := dbInstance.GetAllLocations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, locs)
}

func getCaregiverGeolocationPositions(c *gin.Context) {
	locs, err := dbInstance.GetCaregiverLocations()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, locs)
}

func getGeolocationPositionByUserId(c *gin.Context) {
	userID := uuid.MustParse(c.Param("userid"))
	loc, err := dbInstance.GetGeolocationPositionByUserId(userID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, loc)
}

func deleteGeolocationPosition(c *gin.Context) {
	userId := uuid.MustParse(c.Param("userid"))
	err := dbInstance.DeleteGeolocationPosition(userId)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}
	c.JSON(http.StatusOK, userId)
}
func updateGeolocationPosition(c *gin.Context) {
	r := c.Request
	userId := uuid.MustParse(c.Param("userid"))
	locData := models.GeolocationPosition{}
	if err := render.Bind(r, &locData); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	loc, err := dbInstance.UpdateGeolocationPosition(userId, locData)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}
	c.JSON(http.StatusOK, loc)
}
