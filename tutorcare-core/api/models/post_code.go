package models

import (
	"fmt"
	"net/http"
)

type PostCode struct {
	PostID    int    `sql:",fk" json:"post_id" gorm:"default:null;"`
	CodeID    int    `sql:",pk" json:"code_id" gorm:"primaryKey;default:null;"`
	Code      int    `json:"code" gorm:"default:null;"`
	Verified  bool   `json:"verified" gorm:"default:null"`
	Timestamp string `json:"timestamp" gorm:"default:null"`
}

type PostCodeList struct {
	PostCodes []PostCode `json:"post_codes"`
}

func (i *PostCode) Bind(r *http.Request) error {
	if i.PostID == 0 {
		return fmt.Errorf("PostID is a required field.")
	}
	return nil
}
func (*PostCodeList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*PostCode) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
