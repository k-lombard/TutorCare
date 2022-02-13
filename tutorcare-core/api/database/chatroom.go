package database

import (
	"database/sql"
	"fmt"

	"main/models"

	"github.com/google/uuid"
)

func (db Database) GetAllChatrooms() (*models.ChatroomList, error) {
	list := &models.ChatroomList{}
	err := db.Conn.Order("chatroom_id desc").Find(&list.Chatrooms).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func (db Database) AddChatroom(chatroom *models.Chatroom) (models.Chatroom, error) {
	checkExists := models.Chatroom{}
	chatroomOut := models.Chatroom{}
	switch errLatest := db.Conn.Where("(user1 = ? AND user2 = ?) OR (user1 = ? AND user2 = ?)", chatroom.User1ID, chatroom.User2ID, chatroom.User2ID, chatroom.User1ID).First(&checkExists).Error; errLatest {
	case sql.ErrNoRows:
		err := db.Conn.Create(&chatroom).Error
		if err != nil {
			return chatroomOut, err
		}
		err2 := db.Conn.First(&chatroomOut, "user1 = ? AND user2 = ?", chatroom.User1ID, chatroom.User2ID).Error
		if err2 != nil {
			return chatroomOut, err2
		}
		fmt.Println("New chatroom record created with chatroomID and timestamp: ", chatroomOut.ChatroomID, chatroomOut.DateCreated)
		return chatroomOut, nil
	default:
		return chatroomOut, ErrDuplicate
	}
}

func (db Database) GetChatroomById(chatroomId int) (models.Chatroom, error) {
	chatroomOut := models.Chatroom{}
	switch err := db.Conn.First(&chatroomOut, "chatroom_id = ?", chatroomId).Error; err {
	case sql.ErrNoRows:
		return chatroomOut, ErrNoMatch
	default:
		userOut := models.User{}
		err3 := db.Conn.First(&userOut, "user_id = ?", chatroomOut.User1ID).Error
		if err3 != nil {
			return chatroomOut, err3
		}
		chatroomOut.User1 = userOut
		userOut2 := models.User{}
		err4 := db.Conn.First(&userOut2, "user_id = ?", chatroomOut.User2ID).Error
		if err4 != nil {
			return chatroomOut, err4
		}
		chatroomOut.User2 = userOut2
		return chatroomOut, nil
	}
}

func (db Database) GetChatroomByTwoUsers(userid1 uuid.UUID, userid2 uuid.UUID) (models.Chatroom, error) {
	chatroomOut := models.Chatroom{}
	switch err := db.Conn.Where("user1 = ? AND user2 = ? OR user1 = ? AND user2 = ?", userid1, userid2, userid2, userid1).First(&chatroomOut).Error; err {
	case sql.ErrNoRows:
		return chatroomOut, ErrNoMatch
	default:
		userOut := models.User{}
		err3 := db.Conn.First(&userOut, "user_id = ?", chatroomOut.User1ID).Error
		if err3 != nil {
			return chatroomOut, err3
		}
		chatroomOut.User1 = userOut
		userOut2 := models.User{}
		err4 := db.Conn.First(&userOut2, "user_id = ?", chatroomOut.User2ID).Error
		if err4 != nil {
			return chatroomOut, err4
		}
		chatroomOut.User2 = userOut2
		return chatroomOut, nil
	}
}

func (db Database) GetChatroomsByUserId(userId uuid.UUID) (*models.ChatroomList, error) {
	list := &models.ChatroomList{}
	err := db.Conn.Order("chatroom_id desc").Find(&list.Chatrooms, "user1 = ? OR user2 = ?", userId, userId).Error
	if err != nil {
		return list, err
	}
	for i, chatroom := range list.Chatrooms {
		userOut := models.User{}
		err3 := db.Conn.First(&userOut, "user_id = ?", chatroom.User1ID).Error
		if err3 != nil {
			return list, err3
		}
		list.Chatrooms[i].User1 = userOut
		userOut2 := models.User{}
		err4 := db.Conn.First(&userOut2, "user_id = ?", chatroom.User2ID).Error
		if err4 != nil {
			return list, err4
		}
		list.Chatrooms[i].User2 = userOut2
	}
	return list, nil
}

func (db Database) DeleteChatroom(chatroomId int) error {
	err := db.Conn.Delete(&models.Chatroom{}, chatroomId).Error
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNoMatch
		default:
			return err
		}
	}
	fmt.Println("Chatroom deleted with ChatroomID: ", chatroomId)
	return nil
}
