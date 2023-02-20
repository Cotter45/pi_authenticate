package routes

import (
	"github.com/gin-gonic/gin"

	"auth/model"
)

// Authenticate
func Authenticate(c *gin.Context) {
	user := c.MustGet("user").(model.User)
	userApp := c.MustGet("userApp").(model.UserApp)
	// convert user to safe user
	safeUser := model.SafeUser{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		ApiKey: 	 userApp.ApiKey,
	}

	c.JSON(201, gin.H{
		"message": "success",
		"user":    safeUser,
	})
}
