package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"

	"auth/model"
	"auth/pkg/crypto"
)

// AdminMiddleware is a middleware that verifies the admin token
func AdminMiddleware() gin.HandlerFunc {
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

		// header admin token
		adminToken := c.GetHeader("AdminToken")
		if adminToken == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Unauthorized",
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

		// verify admin token
		if status := crypto.CheckHash(adminToken, app.Key); !status {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		c.Next()
	}
}
