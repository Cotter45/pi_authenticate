package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"

	"auth/model"
)

// AppMiddleware is a middleware that finds the app by the given app name and adds it to the context
func AppMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)
		// header app name
		appName := c.GetHeader("AppName")
		if appName == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "No registered app requested",
			})
			return
		}

		// find app
		var app model.App
		if err := db.Where("name = ?", appName).First(&app).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "No registered app found",
			})
			return
		}

		// add app to context
		c.Set("app", app)
		c.Next()
	}
}
