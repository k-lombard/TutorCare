package models

import (
	"fmt"
	"net/http"
)

type EmailVerification struct {
	Email string `json:"email"`
	Code  int    `json:"code"`
}

type EmailVerificationList struct {
	EmailVerifications []EmailVerification `json:"emailverifications"`
}

func (i *EmailVerification) Bind(r *http.Request) error {
	if i.Code == 0 || i.Email == "" {
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
