package routes

import (
	"github.com/gin-gonic/gin"

	"auth/pkg/middleware"
)

// InitRoutes initializes all routes
func InitRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.DBMiddleware())
	r.Use(middleware.AppMiddleware())

	r.POST("/signup", Signup)
	r.POST("/login", Login)

	r.Use(middleware.AuthMiddleware())

	r.POST("/authenticate", Authenticate)
	r.POST("/logout", Logout)

	r.Use(middleware.AdminMiddleware())
	r.POST("/createApp", CreateApp)

	return r
}
