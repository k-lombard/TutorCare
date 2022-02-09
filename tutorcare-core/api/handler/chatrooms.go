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

func (r routes) chatrooms(rg *gin.RouterGroup) {
	rg.GET("/", getAllChatrooms)
	rg.POST("/", addChatroom)
	rg.GET("/user/:userid", getChatroomsByUserId)
	rg.GET("/:chatroomid", getChatroomById)
	// rg.PUT("/:chatroomid", updateChatroom)
	rg.DELETE("/:chatroomid", deleteChatroom)
	rg.GET("/users/:userid1/:userid2", getChatroomByTwoUsers)
	rg.GET("/websocket/:userid", GetWebsocketToken)
}

func addChatroom(c *gin.Context) {
	chatroom := &models.Chatroom{}
	r := c.Request
	if err := render.Bind(r, chatroom); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	chatroomOut, err := dbInstance.AddChatroom(chatroom)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	c.JSON(http.StatusOK, chatroomOut)
}

func getAllChatrooms(c *gin.Context) {
	chatrooms, err := dbInstance.GetAllChatrooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, chatrooms)
}

func getChatroomById(c *gin.Context) {
	chatroomIDStr := c.Param("chatroomid")
	chatroomID, errConv := strconv.Atoi(chatroomIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid ChatroomID Integer")
		return
	}
	chatroom, err := dbInstance.GetChatroomById(chatroomID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, chatroom)
}

func getChatroomsByUserId(c *gin.Context) {
	userID := uuid.MustParse(c.Param("userid"))
	chatrooms, err := dbInstance.GetChatroomsByUserId(userID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, chatrooms)
}

func getChatroomByTwoUsers(c *gin.Context) {
	userID1 := uuid.MustParse(c.Param("userid1"))
	userID2 := uuid.MustParse(c.Param("userid2"))
	chatroom, err := dbInstance.GetChatroomByTwoUsers(userID1, userID2)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, chatroom)
}

func deleteChatroom(c *gin.Context) {
	chatroomIDStr := c.Param("chatroomid")
	chatroomID, errConv := strconv.Atoi(chatroomIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid ApplicationID Integer")
		return
	}
	err := dbInstance.DeleteChatroom(chatroomID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}
	c.JSON(http.StatusOK, chatroomID)
}

// func updateChatroom(c *gin.Context) {
// 	r := c.Request
// 	chatroomIDStr := c.Param("chatroomid")
// 	chatroomID, errConv := strconv.Atoi(chatroomIDStr)
// 	if errConv != nil {
// 		c.JSON(http.StatusBadRequest, "Error: Invalid ApplicationID Integer")
// 		return
// 	}
// 	chatroomData := models.Chatroom{}
// 	if err := render.Bind(r, &chatroomData); err != nil {
// 		c.JSON(http.StatusBadRequest, "Error: Bad request")
// 		return
// 	}
// 	chatroom, err := dbInstance.UpdateChatroom(chatroomID, chatroomData)
// 	if err != nil {
// 		if err == database.ErrNoMatch {
// 			c.JSON(http.StatusNotFound, "Error: Resource not found")
// 		} else {
// 			c.JSON(http.StatusInternalServerError, "Internal Server Error")
// 		}
// 		return
// 	}
// 	c.JSON(http.StatusOK, chatroom)
// }
