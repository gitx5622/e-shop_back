package Models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

type Address struct {
	ID         uint32    `gorm:"auto_increment;primary_key" json:"id"`
	FirstName  string    `gorm:"type:varchar(100);not_null;" json:"first_name"`
	LastName   string    `gorm:"type:varchar(100);not_null;" json:"last_name"`
	Address    string    `gorm:"type:varchar(100);not_null;unique" json:"address"`
	Author     User      `json:"author"`
	AuthorID   uint32    `gorm:"not null" json:"author_id"`
	City       string    `gorm:"type:varchar(100);not_null" json:"city"`
	Country    string    `json:"country" gorm:"default:'Kenya'"`
	PostalCode string    `json:"postal_code" gorm:"default:'00200'"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *Address) Prepare() {
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Address = html.EscapeString(strings.TrimSpace(u.Address))
	u.Author = User{}
	u.City = html.EscapeString(strings.TrimSpace(u.City))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (s *Address) Validate() map[string]string {

	var errorMessages = make(map[string]string)
	var err error
	// check if the name empty
	if s.FirstName == "" {
		err = errors.New("First Name is Required")
		errorMessages["Required_fistname"] = err.Error()
	}
	if s.LastName == "" {
		err = errors.New("Last Name is Required")
		errorMessages["Required_lastname"] = err.Error()
	}
	if s.Address == "" {
		err = errors.New("Address is Required")
		errorMessages["Required_address"] = err.Error()
	}
	if s.City == "" {
		err = errors.New("City is Required")
		errorMessages["Required_city"] = err.Error()
	}
	if s.Country == "" {
		err = errors.New("Country is Required")
		errorMessages["Required_country"] = err.Error()
	}
	if s.PostalCode == "" {
		err = errors.New("Postal Code is Required")
		errorMessages["Required_postalcode"] = err.Error()
	}
	if s.AuthorID < 1 {
		err = errors.New("Required Author")
		errorMessages["Required_author"] = err.Error()
	}

	return errorMessages
}

func (s *Address) SaveAddress(db *gorm.DB) (*Address, error) {
	var err error
	err = db.Debug().Create(&s).Error
	if err != nil {
		return &Address{}, err
	}
	if s.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", s.AuthorID).Take(&s.Author).Error
		if err != nil {
			return &Address{}, err
		}
	}
	return s, nil

}

func (p *Address) FindUserAddress(db *gorm.DB, uid uint32) (*[]Address, error) {

	var err error
	address := []Address{}
	err = db.Debug().Model(&Address{}).Where("author_id = ?", uid).Limit(100).Order("created_at desc").Find(&address).Error
	if err != nil {
		return &[]Address{}, err
	}
	if len(address) > 0 {
		for i, _ := range address {
			err := db.Debug().Model(&User{}).Where("id = ?", address[i].AuthorID).Take(&address[i].Author).Error
			if err != nil {
				return &[]Address{}, err
			}
		}
	}
	return &address, nil
}
