package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Chatroom struct {
	User1       uuid.UUID `sql:",fk" json:"user1"`
	User2       uuid.UUID `sql:",fk" json:"user2"`
	ChatroomID  int       `sql:",pk" json:"chatroom_id"`
	IsDeleted   bool      `json:"is_deleted"`
	DateCreated string    `json:"date_created"`
}

type ChatroomList struct {
	Chatrooms []Chatroom `json:"chatrooms"`
}

func (i *Chatroom) Bind(r *http.Request) error {
	if i.User1.String() == "" || i.ChatroomID == 0 || (i.User2).String() == "" {
		return fmt.Errorf("User1, user2, and chatroom_id are required fields.")
	}
	return nil
}

func (*ChatroomList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Chatroom) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
