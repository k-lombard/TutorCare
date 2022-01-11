package handler

import (
	"main/database"
	"net/http"

	"main/models"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func (r routes) profile(rg *gin.RouterGroup) {
	rg.PUT("/:userid", updateUserProfile)
}

func updateUserProfile(c *gin.Context) {
	r := c.Request
	userID := uuid.MustParse(c.Param("userid"))
	userData := models.User{}
	if err := render.Bind(r, &userData); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	user, err := dbInstance.UpdateUserProfile(userID, userData)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}
	c.JSON(http.StatusOK, user)
}
