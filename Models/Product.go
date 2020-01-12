package Models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

type Product struct {
	ID			uint32		`gorm:"primary_key;auto_increment" json:"id"`
	Title		string		`gorm:"type:varchar(100);not_null;unique" json:"title"`
	Price		string		`gorm:"type:varchar(100);not_null;" json:"price"`
	Description	string		`gorm:"type:varchar(1000);not_null" json:"description"`
	Author    	User      	`json:"author"`
	AuthorID  	uint32    	`gorm:"not null" json:"author_id"`
	DiscoutPrice int32		`gorm:"nullable" json:"discout_price"`
	ImageUrl1	string		`gorm:"type:varchar(1000);not_null"json:"image_url_1"`
	ImageUrl2	string		`gorm:"type:varchar(1000)" json:"image_url_2"`
	CreatedAt 	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

}

func (s *Product) Prepare()  {
	s.Title = html.EscapeString(strings.TrimSpace(s.Title))
	s.Price = html.EscapeString(strings.TrimSpace(s.Price))
	s.Description = html.EscapeString(strings.TrimSpace(s.Description))
	s.Author = User{}
	s.ImageUrl1 = html.EscapeString(strings.TrimSpace(s.ImageUrl1))
	s.ImageUrl2 = html.EscapeString(strings.TrimSpace(s.ImageUrl2))
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()

}

func (s *Product) Validate() map[string]string  {
	var err error

	var errorMessages = make(map[string]string)

	if s.Title == "" {
		err = errors.New("Required Title")
		errorMessages["Required_title"] = err.Error()
	}
	if s.Price == "" {
		err = errors.New("Required Price")
		errorMessages["Required_price"] = err.Error()

	}
	if s.Description == "" {
		err = errors.New("Required Description")
		errorMessages["Required_description"] = err.Error()
	}
	if s.ImageUrl1 == "" {
		err = errors.New("Required ImageUrl1")
		errorMessages["Required_imageurl1"] = err.Error()

	}
	if s.ImageUrl2 == "" {
		err = errors.New("Required ImageUrl2")
		errorMessages["Required_imageurl2"] = err.Error()
	}
	if s.AuthorID < 1 {
		err = errors.New("Required Author")
		errorMessages["Required_author"] = err.Error()
	}
	return errorMessages
}

func (s *Product) SaveProduct(db *gorm.DB) (*Product, error)  {
	var err error
	err = db.Debug().Model(&Product{}).Create(&s).Error
	if err != nil {
		return &Product{}, err
	}
	if s.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", s.AuthorID).Take(&s.Author).Error
		if err != nil {
			return &Product{}, err
		}
	}
	return s, nil
}

func (p *Product) FindAllProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Limit(100).Order("created_at desc").Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}
	if len(products) > 0 {
		for i, _ := range products {
			err := db.Debug().Model(&User{}).Where("id = ?", products[i].AuthorID).Take(&products[i].Author).Error
			if err != nil {
				return &[]Product{}, err
			}
		}
	}
	return &products, nil
}


func (p *Product) FindProductByID(db *gorm.DB, pid uint64) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Product{}, err
		}
	}
	return p, nil
}


func (p *Product) UpdateAProduct(db *gorm.DB) (*Product, error) {

	var err error

	err = db.Debug().Model(&Product{}).Where("id = ?", p.ID).Updates(Product{Title: p.Title,
		Price: p.Price, Description: p.Description, ImageUrl1: p.ImageUrl1, ImageUrl2: p.ImageUrl2, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Product{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Product{}, err
		}
	}
	return p, nil
}
