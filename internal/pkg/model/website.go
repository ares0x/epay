package model

import (
	"gorm.io/gorm"
)

type WebsiteM struct {
	Id   int64  `json:"id" gorm:"column:id;primary"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Tags string `json:"tags"`
}

// TableName
func (u *WebsiteM) TableName() string {
	return "website"
}

// BeforeCreate
func (u *WebsiteM) BeforeCreate(tx *gorm.DB) (err error) {

	return nil
}
