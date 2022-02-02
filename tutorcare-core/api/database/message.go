package database

import (
	"database/sql"
	"fmt"

	"main/models"

	"github.com/google/uuid"
)

func (db Database) GetAllMessages() (*models.MessageList, error) {
	list := &models.MessageList{}
	rows, err := db.Conn.Query("SELECT *, TO_CHAR(timestamp, 'FMDay, FMDD  HH12:MI') FROM messages ORDER BY message_id DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var message models.Message
		err := rows.Scan(&message.SenderID, &message.MessageID, &message.ChatroomID, &message.Message, &message.IsDeleted, &message.Timestamp, &message.Timestamp)
		if err != nil {
			return list, err
		}
		list.Messages = append(list.Messages, message)
	}
	return list, nil
}

func (db Database) AddMessage(message *models.Message) (models.Message, error) {
	sqlStatement := `INSERT INTO messages (sender, chatroom_id, message) VALUES ($1, $2, $3) RETURNING message_id, is_deleted, timestamp;`
	var message_id int
	var is_deleted bool
	var timestamp string
	messageOut := models.Message{}

	err := db.Conn.QueryRow(sqlStatement, &message.SenderID, &message.ChatroomID, &message.Message).Scan(&message_id, &is_deleted, &timestamp)

	if err != nil {
		return messageOut, err
	}
	err2 := db.Conn.QueryRow(`SELECT * FROM messages, TO_CHAR(timestamp, 'FMDay, FMDD  HH12:MI') WHERE message_id = $1;`, message_id).Scan(&messageOut.SenderID, &messageOut.MessageID, &messageOut.ChatroomID, &messageOut.Message, &messageOut.IsDeleted, &messageOut.Timestamp, &messageOut.Timestamp)
	if err2 != nil {
		return messageOut, err2
	}
	fmt.Println("New message record created with messageID and timestamp: ", message_id, timestamp)
	return messageOut, nil
}

func (db Database) GetMessageById(messageId int) (models.Message, error) {
	messageOut := models.Message{}
	query := `SELECT * FROM messages, TO_CHAR(timestamp, 'FMDay, FMDD  HH12:MI') WHERE message_id = $1;`
	row := db.Conn.QueryRow(query, messageId)
	switch err := row.Scan(&messageOut.SenderID, &messageOut.MessageID, &messageOut.ChatroomID, &messageOut.Message, &messageOut.IsDeleted, &messageOut.Timestamp, &messageOut.Timestamp); err {
	case sql.ErrNoRows:
		return messageOut, ErrNoMatch
	default:
		userOut := models.User{}
		query2 := `SELECT * FROM users WHERE user_id = $1;`
		row2 := db.Conn.QueryRow(query2, messageOut.SenderID)
		err3 := row2.Scan(&userOut.UserID, &userOut.FirstName, &userOut.LastName, &userOut.Email, &userOut.Password, &userOut.DateJoined, &userOut.Status, &userOut.UserCategory, &userOut.Experience, &userOut.Bio)
		if err3 != nil {
			return messageOut, err3
		}
		messageOut.Sender = userOut
		return messageOut, nil
	}
}

func (db Database) GetMessagesByUserId(userId uuid.UUID) (*models.MessageList, error) {
	list := &models.MessageList{}
	rows, err := db.Conn.Query("SELECT * FROM messages, TO_CHAR(timestamp, 'FMDay, FMDD  HH12:MI') WHERE sender = $1 ORDER BY timestamp DESC", userId)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var messageOut models.Message
		err2 := rows.Scan(&messageOut.SenderID, &messageOut.MessageID, &messageOut.ChatroomID, &messageOut.Message, &messageOut.IsDeleted, &messageOut.Timestamp, &messageOut.Timestamp)
		if err2 != nil {
			return list, err2
		}
		userOut := models.User{}
		query2 := `SELECT * FROM users WHERE user_id = $1;`
		row2 := db.Conn.QueryRow(query2, messageOut.SenderID)
		err3 := row2.Scan(&userOut.UserID, &userOut.FirstName, &userOut.LastName, &userOut.Email, &userOut.Password, &userOut.DateJoined, &userOut.Status, &userOut.UserCategory, &userOut.Experience, &userOut.Bio)
		if err3 != nil {
			return list, err3
		}
		messageOut.Sender = userOut
		list.Messages = append(list.Messages, messageOut)
	}
	return list, nil
}

func (db Database) GetMessagesByChatroomId(chatroomId int) (*models.MessageList, error) {
	list := &models.MessageList{}
	rows, err := db.Conn.Query("SELECT *, TO_CHAR(timestamp, 'FMDay, FMDD  HH12:MI') FROM messages WHERE chatroom_id = $1 ORDER BY timestamp DESC", chatroomId)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var messageOut models.Message
		err2 := rows.Scan(&messageOut.SenderID, &messageOut.MessageID, &messageOut.ChatroomID, &messageOut.Message, &messageOut.IsDeleted, &messageOut.Timestamp, &messageOut.Timestamp)
		if err2 != nil {
			return list, err2
		}
		userOut := models.User{}
		query2 := `SELECT * FROM users WHERE user_id = $1;`
		row2 := db.Conn.QueryRow(query2, messageOut.SenderID)
		err3 := row2.Scan(&userOut.UserID, &userOut.FirstName, &userOut.LastName, &userOut.Email, &userOut.Password, &userOut.DateJoined, &userOut.Status, &userOut.UserCategory, &userOut.Experience, &userOut.Bio)
		if err3 != nil {
			return list, err3
		}
		messageOut.Sender = userOut
		list.Messages = append(list.Messages, messageOut)
	}
	return list, nil
}

func (db Database) DeleteMessage(messageId int) error {
	var id int
	query := `DELETE FROM messages WHERE message_id = $1 RETURNING message_id;`
	err := db.Conn.QueryRow(query, messageId).Scan(&id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNoMatch
		default:
			return err
		}
	}
	fmt.Println("Message deleted with MessageID: ", id)
	return nil
}

func (db Database) UpdateMessage(messageId int, messageData models.Message) (models.Message, error) {
	msg := models.Message{}
	query := `UPDATE messages SET message=$1, is_deleted=$2 WHERE message_id=$3 RETURNING sender, message_id, chatroom_id, message, is_deleted, timestamp;`

	query2 := `SELECT *, TO_CHAR(timestamp, 'FMDay, FMDD  HH12:MI') FROM messages WHERE message_id=$1;`
	messageOut := models.Message{}
	errTwo := db.Conn.QueryRow(query2, messageId).Scan(&messageOut.SenderID, &messageOut.MessageID, &messageOut.ChatroomID, &messageOut.Message, &messageOut.IsDeleted, &messageOut.Timestamp, &messageOut.Timestamp)
	if errTwo != nil {
		if errTwo == sql.ErrNoRows {
			return msg, ErrNoMatch
		}
		return msg, errTwo
	}
	err := db.Conn.QueryRow(query, messageData.Message, messageData.IsDeleted, messageId).Scan(&msg.SenderID, &msg.MessageID, &msg.ChatroomID, &msg.Message, &msg.IsDeleted, &msg.Timestamp)
	if err != nil {
		if err == sql.ErrNoRows {
			return msg, ErrNoMatch
		}
		return msg, err
	}
	return msg, nil
}
