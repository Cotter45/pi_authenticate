package model

import (
	"os"

	"gorm.io/gorm"

	"auth/pkg/crypto"
)

// User struct
type App struct {
	gorm.Model
	ID   uint   `json:"id"`
	Name string `gorm:"unique_index;not null" json:"name"` // unique app name
	Key  string `json:"key"`                               // app key - hashed
}

// Seed Apps
func SeedApps(db *gorm.DB) {
	var count int64
	db.Model(&App{}).Count(&count)
	if count == 0 {
		key, err := crypto.CreateHash(os.Getenv("ADMIN_KEY_STRING"))

		if err != nil {
			panic(err)
		}

		db.Create(&App{
			Name: "Admin",
			Key:  key,
		})
	}
}
