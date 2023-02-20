package routes

import (
	"github.com/gin-gonic/gin"

	"auth/pkg/middleware"
)

// InitRoutes initializes all routes
func InitRoutes() *gin.Engine {
	app := gin.Default()

	r := app.Group("/api/auth/v1")
	r.Use(middleware.DBMiddleware())
	r.Use(middleware.AppMiddleware())

	r.POST("/signup", Signup)
	r.POST("/login", Login)

	r.Use(middleware.AuthMiddleware())

	r.POST("/authenticate", Authenticate)
	r.POST("/refresh", Refresh)
	r.POST("/logout", Logout)

	r.POST("/resetCode", ResetCode)
	r.POST("/resetPassword", ResetPassword)

	r.Use(middleware.AdminMiddleware())
	r.POST("/createApp", CreateApp)

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error": "Not found",
		})
	})

	app.NoMethod(func(c *gin.Context) {
		c.JSON(405, gin.H{
			"error": "Method not allowed",
		})
	})

	return app
}
