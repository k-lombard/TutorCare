package database

import (
	"database/sql"
	"fmt"

	"main/models"

	"github.com/google/uuid"
)

func (db Database) GetAllChatrooms() (*models.ChatroomList, error) {
	list := &models.ChatroomList{}
	rows, err := db.Conn.Query("SELECT * FROM chatrooms ORDER BY chatroom_id DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var chatroom models.Chatroom
		err := rows.Scan(&chatroom.User1ID, &chatroom.User2ID, &chatroom.ChatroomID, &chatroom.IsDeleted, &chatroom.DateCreated)
		if err != nil {
			return list, err
		}
		list.Chatrooms = append(list.Chatrooms, chatroom)
	}
	return list, nil
}

func (db Database) AddChatroom(chatroom *models.Chatroom) (models.Chatroom, error) {
	sqlStatement := `INSERT INTO chatrooms (user1, user2) VALUES ($1, $2) RETURNING chatroom_id, is_deleted, date_created;`
	var chatroom_id int
	var is_deleted bool
	var date_created string
	chatroomOut := models.Chatroom{}

	err := db.Conn.QueryRow(sqlStatement, &chatroom.User1ID, &chatroom.User2ID).Scan(&chatroom_id, &is_deleted, &date_created)

	if err != nil {
		return chatroomOut, err
	}
	err2 := db.Conn.QueryRow(`SELECT * FROM chatrooms WHERE chatroom_id = $1;`, chatroom_id).Scan(&chatroomOut.User1ID, &chatroomOut.User2ID, &chatroomOut.ChatroomID, &chatroomOut.IsDeleted, &chatroomOut.DateCreated)
	if err2 != nil {
		return chatroomOut, err2
	}
	fmt.Println("New chatroom record created with chatroomID and timestamp: ", chatroom_id, date_created)
	return chatroomOut, nil
}

func (db Database) GetChatroomById(chatroomId int) (models.Chatroom, error) {
	chatroomOut := models.Chatroom{}
	query := `SELECT * FROM chatrooms WHERE chatroom_id = $1;`
	row := db.Conn.QueryRow(query, chatroomId)
	switch err := row.Scan(&chatroomOut.User1ID, &chatroomOut.User2ID, &chatroomOut.ChatroomID, &chatroomOut.IsDeleted, &chatroomOut.DateCreated); err {
	case sql.ErrNoRows:
		return chatroomOut, ErrNoMatch
	default:
		userOut := models.User{}
		query2 := `SELECT * FROM users WHERE user_id = $1;`
		row2 := db.Conn.QueryRow(query2, chatroomOut.User1ID)
		err3 := row2.Scan(&userOut.UserID, &userOut.FirstName, &userOut.LastName, &userOut.Email, &userOut.Password, &userOut.DateJoined, &userOut.Status, &userOut.UserCategory, &userOut.Experience, &userOut.Bio)
		if err3 != nil {
			return chatroomOut, err3
		}
		chatroomOut.User1 = userOut
		userOut2 := models.User{}
		query3 := `SELECT * FROM users WHERE user_id = $1;`
		row3 := db.Conn.QueryRow(query3, chatroomOut.User2ID)
		err4 := row3.Scan(&userOut2.UserID, &userOut2.FirstName, &userOut2.LastName, &userOut2.Email, &userOut2.Password, &userOut2.DateJoined, &userOut2.Status, &userOut2.UserCategory, &userOut2.Experience, &userOut2.Bio)
		if err4 != nil {
			return chatroomOut, err4
		}
		chatroomOut.User2 = userOut2
		return chatroomOut, nil
	}
}

func (db Database) GetChatroomsByUserId(userId uuid.UUID) (*models.ChatroomList, error) {
	list := &models.ChatroomList{}
	rows, err := db.Conn.Query("SELECT * FROM chatrooms WHERE user1 = $1 OR user2 = $1 ORDER BY chatroom_id DESC", userId)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var chatroomOut models.Chatroom
		err2 := rows.Scan(&chatroomOut.User1ID, &chatroomOut.User2ID, &chatroomOut.ChatroomID, &chatroomOut.IsDeleted, &chatroomOut.DateCreated)
		if err2 != nil {
			return list, err2
		}
		userOut := models.User{}
		query2 := `SELECT * FROM users WHERE user_id = $1;`
		row2 := db.Conn.QueryRow(query2, chatroomOut.User1ID)
		err3 := row2.Scan(&userOut.UserID, &userOut.FirstName, &userOut.LastName, &userOut.Email, &userOut.Password, &userOut.DateJoined, &userOut.Status, &userOut.UserCategory, &userOut.Experience, &userOut.Bio)
		if err3 != nil {
			return list, err3
		}
		chatroomOut.User1 = userOut
		userOut2 := models.User{}
		query3 := `SELECT * FROM users WHERE user_id = $1;`
		row3 := db.Conn.QueryRow(query3, chatroomOut.User2ID)
		err4 := row3.Scan(&userOut2.UserID, &userOut2.FirstName, &userOut2.LastName, &userOut2.Email, &userOut2.Password, &userOut2.DateJoined, &userOut2.Status, &userOut2.UserCategory, &userOut2.Experience, &userOut2.Bio)
		if err4 != nil {
			return list, err4
		}
		chatroomOut.User2 = userOut2
		list.Chatrooms = append(list.Chatrooms, chatroomOut)
	}
	return list, nil
}

func (db Database) DeleteChatroom(chatroomId int) error {
	var id int
	query := `DELETE FROM chatrooms WHERE chatroom_id = $1 RETURNING chatroom_id;`
	err := db.Conn.QueryRow(query, chatroomId).Scan(&id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNoMatch
		default:
			return err
		}
	}
	fmt.Println("Chatroom deleted with ChatroomID: ", id)
	return nil
}

// func (db Database) UpdateApplication(applicationId int, appData models.Application) (models.Application, error) {
// 	app := models.Application{}
// 	query := `UPDATE applications SET message=$1, accepted=$2 WHERE application_id=$3 RETURNING user_id, post_id, application_id, message, accepted, date_created;`

// 	query2 := `SELECT * FROM applications WHERE application_id=$1;`
// 	app2 := models.Application{}
// 	errTwo := db.Conn.QueryRow(query2, applicationId).Scan(&app2.UserID, &app2.ApplicationID, &app2.PostID, &app2.Message, &app2.Accepted, &app2.DateCreated)
// 	if errTwo != nil {
// 		if errTwo == sql.ErrNoRows {
// 			return app, ErrNoMatch
// 		}
// 		return app, errTwo
// 	}
// 	err := db.Conn.QueryRow(query, appData.Message, appData.Accepted, applicationId).Scan(&app.UserID, &app.ApplicationID, &app.PostID, &app.Message, &app.Accepted, &app.DateCreated)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return app, ErrNoMatch
// 		}
// 		return app, err
// 	}
// 	return app, nil
// }
