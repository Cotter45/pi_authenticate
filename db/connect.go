package db

import (
	"fmt"
	"auth/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB connect to db (sqlite)
func InitDB() {
	var err error
	if err != nil {
		fmt.Println(err)
	}

	DB, err = gorm.Open(
		sqlite.Open("db/auth.db"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent), PrepareStmt: true,
		})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.App{})
	DB.AutoMigrate(&model.UserApp{})

	model.SeedApps(DB)
	fmt.Println("Database Migrated")
}
