package Models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
)

type MpesaStkPush struct {
	ID       int32  `gorm:"auto_increment;primary_key"`
	PartyA   string `gorm:"not_null"  json:"party_a"`
	Amount   string `gorm:"not_null" json:"amount"`
	Author   User   `json:"author"`
	AuthorID uint32 `gorm:"not null" json:"author_id"`
}

func (u *MpesaStkPush) Prepare() {
	u.PartyA = html.EscapeString(strings.TrimSpace(u.PartyA))
	u.Amount = html.EscapeString(strings.TrimSpace(u.Amount))
}

func (s *MpesaStkPush) Validate() map[string]string {

	var errorMessages = make(map[string]string)
	var err error
	// check if the name empty
	if s.PartyA== "" {
		err = errors.New("PartyA is Required")
		errorMessages["Required_partya"] = err.Error()
	}
	if s.Amount == "" {
		err = errors.New("Amount is Required")
		errorMessages["Required_amount"] = err.Error()
	}


	return errorMessages
}


func (s *MpesaStkPush) SaveMpesaTrasaction(db *gorm.DB) (*MpesaStkPush, error) {
	var err error
	err = db.Debug().Create(&s).Error
	if err != nil {
		return &MpesaStkPush{}, err
	}
	if s.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", s.AuthorID).Take(&s.Author).Error
		if err != nil {
			return &MpesaStkPush{}, err
		}
	}
	return s, nil

}
