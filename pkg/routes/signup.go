package routes

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"auth/model"
	"auth/pkg/crypto"
)

// Sing up user struct
type SignupUser struct {
	Username string `json:"username"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Signup
func Signup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	app := c.MustGet("app").(model.App)

	var signupUser SignupUser
	if err := c.ShouldBindJSON(&signupUser); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// check if user exists
	var user model.User
	if err := db.Where("email = ?", signupUser.Email).First(&user).Error; err == nil {
		c.JSON(400, gin.H{
			"error": "Error signing up, please try again or try different credentials",
		})
		return
	}

	// hash password
	hashedPassword, err := crypto.CreateHash(signupUser.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error signing up, please try again or try different credentials",
		})
		return
	}

	// create user
	user = model.User{
		Username: signupUser.Username,
		Email:    signupUser.Email,
		Password: hashedPassword,
	}

	db.Create(&user)

	// create a random api key uuid string
	uid := uuid.New().String()

	// create user app relationship
	userApp := model.UserApp{
		UserID: user.ID,
		AppID:  app.ID,
		ApiKey: uid,
	}

	db.Create(&userApp)

	// convert to safe user
	safeUser := model.SafeUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		ApiKey:   userApp.ApiKey,
		AppId:    app.ID,
	}

	// create JWT token
	JWTManager := model.NewJWTManager(os.Getenv("JWT_SECRET"), 1*time.Hour)
	token,  err := JWTManager.Generate(&safeUser)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error creating token",
		})
		return
	}

	// remove api key from response
	safeUser.ApiKey = ""

	c.JSON(201, gin.H{
		"message": "success",
		"user":    safeUser,
		"token":   token,
	})
}
