package models

import (
	"fmt"
	"net/http"
)

type WebsocketToken struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}

func (i *WebsocketToken) Bind(r *http.Request) error {
	if i.Token == "" || i.UserID == "" {
		return fmt.Errorf("UserID and token are required fields.")
	}
	return nil
}
