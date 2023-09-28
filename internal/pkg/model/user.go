package model

import (
	"github.com/marmotedu/component-base/pkg/auth"
	"gorm.io/gorm"
	"time"
)

type UserM struct {
	ID        int64     `gorm:"column:id;primary_key"`
	Uid       int64     `json:"uid" gorm:"column:uid;not null;default:0;unique;"`
	Nickname  string    `gorm:"column:nickname"`
	Email     string    `gorm:"column:email"`
	Avatar    string    `json:"avatar"`
	Password  string    `gorm:"column:password;not null"`
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
}

// TableName
func (u *UserM) TableName() string {
	return "user"
}

// BeforeCreate
func (u *UserM) BeforeCreate(tx *gorm.DB) (err error) {
	// Encrypt the user password.
	u.Uid = 1
	u.Password, err = auth.Encrypt(u.Password)
	if err != nil {
		return err
	}

	return nil
}
