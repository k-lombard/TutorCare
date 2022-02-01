package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Message struct {
	Sender     uuid.UUID `sql:",fk" json:"sender"`
	MessageID  int       `sql:",pk" json:"message_id"`
	ChatroomID int       `sql:",fk" json:"chatroom_id"`
	IsDeleted  bool      `json:"is_deleted"`
	Timestamp  string    `json:"timestamp"`
}

type MessageList struct {
	Messages []Message `json:"messages"`
}

func (i *Message) Bind(r *http.Request) error {
	if i.Sender.String() == "" || i.ChatroomID == 0 || i.MessageID == 0 {
		return fmt.Errorf("User1, user2, and chatroom_id are required fields.")
	}
	return nil
}

func (*MessageList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Message) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
