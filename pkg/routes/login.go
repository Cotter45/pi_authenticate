package routes

import (
	"time"
	"os"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"auth/model"
	"auth/pkg/crypto"
)

// LoginUser is a struct that represents the login user
type LoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	app := c.MustGet("app").(model.App)

	var loginUser LoginUser
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// check if user exists
	var user model.User
	if err := db.Where("email = ?", loginUser.Email).First(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Incorrect credentials",
		})
		return
	}

	// check if user - app relationship exists
	var userApp model.UserApp
	if err := db.Where("user_id = ? AND app_id = ?", user.ID, app.ID).First(&userApp).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Incorrect credentials",
		})
		return
	}

	// check if password is correct
	if status := crypto.CheckHash(loginUser.Password, user.Password); !status {
		c.JSON(400, gin.H{
			"error": "Incorrect credentials",
		})
		return
	}

	// create a random api key uuid string
	uid := uuid.New().String()

	// update user app key to new key
	db.Model(&userApp).Update("api_key", uid)

	// convert to safe user
	safeUser := model.SafeUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		ApiKey:   userApp.ApiKey,
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
