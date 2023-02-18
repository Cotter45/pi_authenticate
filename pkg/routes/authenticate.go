package routes

import (
	"time"
	"os"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"auth/model"
)

// Authenticate
func Authenticate(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	user := c.MustGet("user").(model.User)
	userApp := c.MustGet("userApp").(model.UserApp)

	// create a random api key uuid string
	uid := uuid.New().String()

	// update user app key to new key
	db.Model(&userApp).Update("api_key", uid)

	// convert user to safe user
	safeUser := model.SafeUser{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		ApiKey:    uid,
	}

	// create JWT token
	JWTManager := model.NewJWTManager(os.Getenv("JWT_SECRET"), 1*time.Hour)
	token,  err := JWTManager.Generate(&safeUser)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error logging in",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "success",
		"user":    safeUser,
		"token":   token,
	})
}
