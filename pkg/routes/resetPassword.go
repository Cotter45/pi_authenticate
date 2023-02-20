package routes

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"

	"auth/model"
	"auth/pkg/crypto"
)

type ResetPasswordRequest struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
	ResetCode   string `json:"resetCode" binding:"required"`
}

// Reset password
func ResetPassword(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	
	var request ResetPasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
		return
	}
	
	// find user
	var user model.User
	if err := db.Where("reset_token = ? AND reset_token IS NOT NULL", request.ResetCode).First(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// check if old password is correct
	if !crypto.CheckHash(request.OldPassword, user.Password) {
		c.JSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}

	hash, err := crypto.CreateHash(request.NewPassword)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error creating hash",
		})
		return
	}

	// update user app key to new key
	db.Model(&user).Updates(map[string]interface{}{"password": hash, "reset_token": gorm.Expr("NULL")})

	// return success
	c.JSON(201, gin.H{
		"message": "success",
	})
}
