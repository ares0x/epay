package model

import (
	"gorm.io/gorm"
)

type CollectionM struct {
}

// TableName
func (u *CollectionM) TableName() string {
	return "collection"
}

// BeforeCreate
func (u *CollectionM) BeforeCreate(tx *gorm.DB) (err error) {

	return nil
}
