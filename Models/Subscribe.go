package Models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
)

type Subscribe struct {
	 ID		uint32			`gorm:"auto_increment;primary_key" json:"id"`
	 Email	string			`gorm:"type:varchar(100);not_null;unique" json:"email" validate:"required,email"`
}

func (u *Subscribe) Prepare() {
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

}

func (s *Subscribe) Validate() map[string]string {

	var errorMessages = make(map[string]string)
	var err error
	// check if the name empty
	if s.Email == "" {
		err = errors.New("Email is Required")
		errorMessages["Required_email"] = err.Error()
	}
	
	if s.Email != "" {
		if err = checkmail.ValidateFormat(s.Email); err != nil {
			err = errors.New("Invalid Email")
			errorMessages["Invalid_email"] = err.Error()
		}
	}
	return errorMessages
}

func (s *Subscribe) Subscribed(db *gorm.DB) (*Subscribe, error)  {
	var err error
	err = db.Debug().Create(&s).Error
	if err != nil {
		return &Subscribe{}, err
	}
	return s, nil

}