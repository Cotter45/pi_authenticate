package db

import (
	"fmt"
	"os"
	"auth/model"

	// "gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

// InitDB connect to db (sqlite)
func InitDB() {
	var err error
	if err != nil {
		fmt.Println(err)
	}

	// DB, err = gorm.Open(
	// 	sqlite.Open("db/auth.db"),
	// 	&gorm.Config{
	// 		Logger: logger.Default.LogMode(logger.Silent), PrepareStmt: true,
	// 	})

	// if err != nil {
	// 	fmt.Println(err)
	// 	panic("failed to connect database")
	// }

	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=Asia/Shanghai"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
