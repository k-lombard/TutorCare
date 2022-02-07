package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Chatroom struct {
	User1ID     uuid.UUID `sql:",fk" json:"user1_id"`
	User2ID     uuid.UUID `sql:",fk" json:"user2_id"`
	ChatroomID  int       `sql:",pk" json:"chatroom_id"`
	IsDeleted   bool      `json:"is_deleted"`
	DateCreated string    `json:"date_created"`
	User1       User      `json:"user1"`
	User2       User      `json:"user2"`
	Messages    []Message `json:"messages"`
}

type ChatroomList struct {
	Chatrooms []Chatroom `json:"chatrooms"`
}

func (i *Chatroom) Bind(r *http.Request) error {
	if i.User1ID.String() == "" || (i.User2ID).String() == "" {
		return fmt.Errorf("User1ID, User2ID are required fields.")
	}
	return nil
}

func (*ChatroomList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Chatroom) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
