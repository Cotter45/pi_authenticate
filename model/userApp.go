package model

import "gorm.io/gorm"

// User - App relationship
type UserApp struct {
	gorm.Model
	UserID uint   `gorm:"not null" json:"userId"`
	AppID  uint   `gorm:"not null" json:"appId"`
	ApiKey string `gorm:"not null" json:"apiKey"` // random string generated at signup, changed on password reset, login and authenticate
	User	 User `gorm:"foreignKey:UserID"`
	App		 App `gorm:"foreignKey:AppID"`
}