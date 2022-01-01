package handler

import (
	"main/database"
	"net/http"

	"main/models"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

var userIDKey = "userID"

func (r routes) users(rg *gin.RouterGroup) {

	rg.GET("/", getAllUsers)
	rg.POST("/", addUser)
	rg.GET("/:userid", getUserById)
	rg.PUT("/:userid", updateUser)
	rg.DELETE("/:userid", deleteUser)
}

func addUser(c *gin.Context) {
	user := &models.User{}
	r := c.Request
	if err := render.Bind(r, user); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	userOut, err := dbInstance.AddUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	c.JSON(http.StatusOK, userOut)
}

func getAllUsers(c *gin.Context) {
	users, err := dbInstance.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, users)
}

func getUserById(c *gin.Context) {
	userID := uuid.MustParse(c.Param("userid"))
	user, err := dbInstance.GetUserById(userID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, user)
}

func deleteUser(c *gin.Context) {
	userId := uuid.MustParse(c.Param("userid"))
	err := dbInstance.DeleteUser(userId)
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
func updateUser(c *gin.Context) {
	r := c.Request
	userId := uuid.MustParse(c.Param("userid"))
	userData := models.User{}
	if err := render.Bind(r, &userData); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	user, err := dbInstance.UpdateUser(userId, userData)
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
