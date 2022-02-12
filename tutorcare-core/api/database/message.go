package database

import (
	"database/sql"
	"fmt"

	"main/models"

	"github.com/google/uuid"
)

func (db Database) GetAllMessages() (*models.MessageList, error) {
	list := &models.MessageList{}
	err := db.Conn.Order("message_id desc").Find(&list.Messages).Error
	if err != nil {
		return list, err
	}
	for _, message := range list.Messages {
		errThree := db.Conn.Select("TO_CHAR(timestamp, 'FMDay, FMDD  HH12:MI')").First(&message.Timestamp, "message_id = ?", message.MessageID).Error
		if errThree != nil {
			return list, errThree
		}
	}
	return list, nil
}

func (db Database) AddMessage(message *models.Message) (models.Message, error) {
	messageOut := models.Message{}

	err := db.Conn.Create(&message).Error
	if err != nil {
		return messageOut, err
	}
	err2 := db.Conn.Select("*", "TO_CHAR(timestamp, 'FMDay, FMDD  HH12:MI')").First(&messageOut, &messageOut.Timestamp).Error
	if err2 != nil {
		return messageOut, err2
	}
	fmt.Println("New message record created with messageID and timestamp: ", messageOut.MessageID, messageOut.Timestamp)
	return messageOut, nil
}

func (db Database) GetMessageById(messageId int) (models.Message, error) {
	messageOut := models.Message{}
	switch err2 := db.Conn.Select("*", "TO_CHAR(timestamp, 'FMDay, FMDD  HH12:MI')").First(&messageOut, &messageOut.Timestamp).Error; err2 {
	case sql.ErrNoRows:
		return messageOut, ErrNoMatch
	default:
		userOut := models.User{}
		err3 := db.Conn.First(&userOut, "user_id = ?", messageOut.SenderID).Error
		if err3 != nil {
			return messageOut, err3
		}
		messageOut.Sender = userOut
		return messageOut, nil
	}
}

func (db Database) GetMessagesByUserId(userId uuid.UUID) (*models.MessageList, error) {
	list := &models.MessageList{}
	err := db.Conn.Where("user_id = ?", userId).Order("message_id desc").Find(&list.Messages).Error
	if err != nil {
		return list, err
	}
	for _, message := range list.Messages {
		errThree := db.Conn.Select("TO_CHAR(timestamp, 'FMDay, FMDD  HH12:MI')").First(&message.Timestamp, "message_id = ?", message.MessageID).Error
		if errThree != nil {
			return list, errThree
		}
		userOut := models.User{}
		err3 := db.Conn.First(&userOut, "user_id = ?", message.SenderID).Error
		if err3 != nil {
			return list, err3
		}
		message.Sender = userOut
	}
	return list, nil
}

func (db Database) GetMessagesByChatroomId(chatroomId int) (*models.MessageList, error) {
	list := &models.MessageList{}
	err := db.Conn.Where("chatroom_id = ?", chatroomId).Order("message_id desc").Find(&list.Messages).Error
	if err != nil {
		return list, err
	}
	for _, message := range list.Messages {
		errThree := db.Conn.Select("TO_CHAR(timestamp, 'FMDay, FMDD  HH12:MI')").First(&message.Timestamp, "message_id = ?", message.MessageID).Error
		if errThree != nil {
			return list, errThree
		}
		userOut := models.User{}
		err3 := db.Conn.First(&userOut, "user_id = ?", message.SenderID).Error
		if err3 != nil {
			return list, err3
		}
		message.Sender = userOut
	}
	return list, nil
}

func (db Database) DeleteMessage(messageId int) error {
	err := db.Conn.Delete(&models.Message{}, messageId).Error
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNoMatch
		default:
			return err
		}
	}
	fmt.Println("Message deleted with MessageID: ", messageId)
	return nil
}

func (db Database) UpdateMessage(messageId int, messageData models.Message) (models.Message, error) {
	msg := models.Message{}
	messageOut := models.Message{}
	errTwo := db.Conn.Select("*", "TO_CHAR(timestamp, 'FMDay, FMDD  HH12:MI')").First(&messageOut, &messageOut.Timestamp).Error
	if errTwo != nil {
		if errTwo == sql.ErrNoRows {
			return msg, ErrNoMatch
		}
		return msg, errTwo
	}
	err := db.Conn.Model(&msg).Updates(models.Message{MessageID: messageId, Message: messageData.Message, IsDeleted: messageData.IsDeleted}).Error
	if err != nil {
		if err == sql.ErrNoRows {
			return msg, ErrNoMatch
		}
		return msg, err
	}
	return msg, nil
}
