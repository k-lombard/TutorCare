package database

import (
	"errors"
	"fmt"

	"main/models"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

func (db Database) GetAllMessages() (*models.MessageList, error) {
	list := &models.MessageList{}
	err := db.Conn.Order("message_id desc").Find(&list.Messages).Error
	if err != nil {
		return list, err
	}
	for i, message := range list.Messages {
		errThree := db.Conn.Where("message_id = ?", message.MessageID).Select("TO_CHAR(timestamp at time zone 'US/Eastern', 'FMDay, FMDD  HH12:MI AM') as timestamp").First(&list.Messages[i]).Error
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
	err2 := db.Conn.Where("message_id = ?", message.MessageID).Select("*", "TO_CHAR(timestamp at time zone 'US/Eastern', 'FMDay, FMDD  HH12:MI AM') as timestamp").First(&messageOut, &messageOut).Error
	if err2 != nil {
		return messageOut, err2
	}
	fmt.Println("New message record created with messageID and timestamp: ", messageOut.MessageID, messageOut.Timestamp)
	return messageOut, nil
}

func (db Database) GetMessageById(messageId int) (models.Message, error) {
	messageOut := models.Message{}
	err2 := db.Conn.Select("*", "TO_CHAR(timestamp at time zone 'US/Eastern', 'FMDay, FMDD  HH12:MI AM') as timestamp").First(&messageOut, &messageOut).Error
	if err2 != nil {
		if errors.Is(err2, gorm.ErrRecordNotFound) {
			return messageOut, ErrNoMatch
		}
		return messageOut, err2
	} else {
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
	for i, message := range list.Messages {
		errThree := db.Conn.Where("message_id = ?", list.Messages[i].MessageID).Select("TO_CHAR(timestamp at time zone 'US/Eastern', 'FMDay, FMDD  HH12:MI AM') as timestamp").First(&list.Messages[i].MessageID).Error
		if errThree != nil {
			return list, errThree
		}
		userOut := models.User{}
		err3 := db.Conn.First(&userOut, "user_id = ?", message.SenderID).Error
		if err3 != nil {
			return list, err3
		}
		list.Messages[i].Sender = userOut
	}
	return list, nil
}

func (db Database) GetMessagesByChatroomId(chatroomId int) (*models.MessageList, error) {
	list := &models.MessageList{}
	err := db.Conn.Where("chatroom_id = ?", chatroomId).Order("message_id asc").Find(&list.Messages).Error
	if err != nil {
		return list, err
	}
	for i, message := range list.Messages {
		errThree := db.Conn.Where("message_id = ?", message.MessageID).Select("TO_CHAR(timestamp at time zone 'US/Eastern', 'FMDay, FMDD  HH12:MI AM') as timestamp").First(&list.Messages[i]).Error
		if errThree != nil {
			return list, errThree
		}
		userOut := models.User{}
		err3 := db.Conn.First(&userOut, "user_id = ?", message.SenderID).Error
		if err3 != nil {
			return list, err3
		}
		list.Messages[i].Sender = userOut
	}
	return list, nil
}

func (db Database) DeleteMessage(messageId int) error {
	err := db.Conn.Delete(&models.Message{}, messageId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNoMatch
		} else {
			return err
		}
	}
	fmt.Println("Message deleted with MessageID: ", messageId)
	return nil
}

func (db Database) UpdateMessage(messageId int, messageData models.Message) (models.Message, error) {
	msg := models.Message{}
	messageOut := models.Message{}
	errTwo := db.Conn.Where("message_id = ?", messageId).Select("*", "TO_CHAR(timestamp at time zone 'US/Eastern', 'FMDay, FMDD  HH12:MI AM') as timestamp").First(&messageOut, &messageOut).Error
	if errTwo != nil {
		if errors.Is(errTwo, gorm.ErrRecordNotFound) {
			return msg, ErrNoMatch
		}
		return msg, errTwo
	}
	err := db.Conn.Model(&msg).Where("message_id = ?", messageId).Updates(models.Message{Message: messageData.Message, IsDeleted: messageData.IsDeleted}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return msg, ErrNoMatch
		}
		return msg, err
	}
	errOut := db.Conn.First(&msg, "message_id = ?", messageId).Error
	if errOut != nil {
		if errors.Is(errOut, gorm.ErrRecordNotFound) {
			return msg, ErrNoMatch
		}
		return msg, errOut
	}
	return msg, nil
}
