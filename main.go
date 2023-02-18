package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"auth/db"
	"auth/pkg/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	gin.SetMode(os.Getenv("GIN_MODE"))

	db.InitDB()
	r := routes.InitRoutes()
	r.Run(":8080")
}
