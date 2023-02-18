package routes

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"auth/model"
)

// Logout
func Logout(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userApp := c.MustGet("userApp").(model.UserApp)

	// create a random api key uuid string
	uid := uuid.New().String()

	// update user app key to new key
	db.Model(&userApp).Update("api_key", uid)

	c.JSON(201, gin.H{
		"message": "success",
	})
}
