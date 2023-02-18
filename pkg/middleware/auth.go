package middleware

import (
	"os"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"auth/model"
)

// AuthMiddleware grabs the Authorization header, verifies the token, verifies the user & app with the db, and adds the user to the context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)

		// header token
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		// parse out token from Bearer: token
		token = token[8:]

		// ensure valid jwt
		JWTManager := model.NewJWTManager(os.Getenv("JWT_SECRET"), 1*time.Hour)
		claims, err := JWTManager.Verify([]byte(token))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		// find app user with api key claims
		var userApp model.UserApp
		if err := db.Where("api_key = ?", claims.ApiKey).First(&userApp).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "No registered app found",
			})
			return
		}

		// find user
		var user model.User
		if err := db.Where("id = ?", claims.UserID).First(&user).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "No registered user found",
			})
			return
		}

		// add user to context
		c.Set("user", user)
		c.Set("userApp", userApp)
		c.Next()
	}
}