package model

import (
	"gorm.io/gorm"
)

type OAuth2M struct {
	Id           int64  `json:"id" gorm:"column:id;primary"`
	Uid          int64  `json:"uid" gorm:"column:uid;type:"`
	OAuthType    int    `json:"oAuthType" grom:"column:o_auth_type"`
	OauthId      string `json:"oauthId"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// TableName
func (u *OAuth2M) TableName() string {
	return "oauth2"
}

// BeforeCreate
func (u *OAuth2M) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}
