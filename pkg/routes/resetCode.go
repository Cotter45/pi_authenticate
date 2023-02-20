package routes

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"auth/model"
)

// Reset password code
func ResetCode(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	user := c.MustGet("user").(model.User)

	// create a random api key uuid string
	uid := uuid.New().String()

	// update user app key to new key
	db.Model(&user).Update("reset_token", uid)

	c.JSON(201, gin.H{
		"message": "success",
		"reset_code": uid,
	})
}