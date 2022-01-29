package models

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type EmailVerification struct {
	UserID uuid.UUID `sql:",fk" json:"user_id"`
	Code   int       `json:"application_id"`
}

type EmailVerificationList struct {
	EmailVerifications []EmailVerification `json:"emailverifications"`
}

func (i *EmailVerification) Bind(r *http.Request) error {
	if i.Code == 0 || (i.UserID).String() == "" {
		return fmt.Errorf("Code and user_id are required fields.")
	}
	return nil
}

func (*EmailVerificationList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*EmailVerification) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
