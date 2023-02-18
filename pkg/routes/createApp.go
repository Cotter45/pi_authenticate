package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"auth/model"
)

// CreateApp creates a new app
func CreateApp(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var app model.App
	if err := c.ShouldBindJSON(&app); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if app.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "App name is required",
		})
		return
	}

	db.Create(&app)
	c.JSON(http.StatusCreated, gin.H{
		"message": "App created",
		"app":     app,
	})
}
