package model

import "gorm.io/gorm"

// User struct
type User struct {
	gorm.Model
	ID         uint   `json:"id"`
	Username   string `gorm:"unique_index;" json:"username"` // app may not use usernames
	Email      string `gorm:"unique_index;not null" json:"email"`
	Password   string `gorm:"not null" json:"password"` // hashed
	ResetToken string `json:"resetToken"`               // random string for password reset
}

// Safe user struct
type SafeUser struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	ApiKey     string `json:"ApiKey"` // random string for password reset or jwt refresh
}
