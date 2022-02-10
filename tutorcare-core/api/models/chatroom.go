package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chatroom struct {
	gorm.Model
	User1ID     uuid.UUID `sql:",fk" json:"user1_id" gorm:"column:user1"`
	User2ID     uuid.UUID `sql:",fk" json:"user2_id" gorm:"column:user2"`
	ChatroomID  int       `sql:",pk" json:"chatroom_id" gorm:"primaryKey"`
	IsDeleted   bool      `json:"is_deleted"`
	DateCreated string    `json:"date_created" gorm:"autoCreateTime"`
	User1       User      `json:"user1" gorm:"-"`
	User2       User      `json:"user2" gorm:"-"`
	Messages    []Message `json:"messages" gorm:"-"`
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
