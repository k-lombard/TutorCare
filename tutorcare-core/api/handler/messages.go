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

func (r routes) messages(rg *gin.RouterGroup) {
	rg.GET("/", getAllMessages)
	rg.POST("/", addMessage)
	rg.GET("/user/:userid", getMessagesByUserId)
	rg.GET("/chatroom/:chatroomid", getMessagesByChatroomId)
	rg.GET("/:messageid", getMessageById)
	rg.PUT("/:messageid", updateMessage)
	rg.DELETE("/:messageid", deleteMessage)
}

func addMessage(c *gin.Context) {
	msg := &models.Message{}
	r := c.Request
	if err := render.Bind(r, msg); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	msgOut, err := dbInstance.AddMessage(msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	c.JSON(http.StatusOK, msgOut)
}

func getAllMessages(c *gin.Context) {
	messages, err := dbInstance.GetAllMessages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, messages)
}

func getMessageById(c *gin.Context) {
	msgIDStr := c.Param("messageid")
	msgID, errConv := strconv.Atoi(msgIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid MessageID Integer")
		return
	}
	msg, err := dbInstance.GetMessageById(msgID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, msg)
}

func getMessagesByUserId(c *gin.Context) {
	userID := uuid.MustParse(c.Param("userid"))
	messages, err := dbInstance.GetMessagesByUserId(userID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, messages)
}

func getMessagesByChatroomId(c *gin.Context) {
	chatroomIDStr := c.Param("chatroomid")
	chatroomID, errConv := strconv.Atoi(chatroomIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid ChatroomID Integer")
		return
	}
	messages, err := dbInstance.GetMessagesByChatroomId(chatroomID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusBadRequest, "Error: Bad request")
		}
		return
	}
	c.JSON(http.StatusOK, messages)
}

func deleteMessage(c *gin.Context) {
	messageIDStr := c.Param("messageid")
	messageID, errConv := strconv.Atoi(messageIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid MessageID Integer")
		return
	}
	err := dbInstance.DeleteMessage(messageID)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}
	c.JSON(http.StatusOK, messageID)
}

func updateMessage(c *gin.Context) {
	r := c.Request
	messageIDStr := c.Param("messageid")
	msgID, errConv := strconv.Atoi(messageIDStr)
	if errConv != nil {
		c.JSON(http.StatusBadRequest, "Error: Invalid ApplicationID Integer")
		return
	}
	msgData := models.Message{}
	if err := render.Bind(r, &msgData); err != nil {
		c.JSON(http.StatusBadRequest, "Error: Bad request")
		return
	}
	message, err := dbInstance.UpdateMessage(msgID, msgData)
	if err != nil {
		if err == database.ErrNoMatch {
			c.JSON(http.StatusNotFound, "Error: Resource not found")
		} else {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}
	c.JSON(http.StatusOK, message)
}
