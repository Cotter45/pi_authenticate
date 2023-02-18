package middleware

import (
	"github.com/gin-gonic/gin"

	"auth/db"
)

// DBMiddleware
func DBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db.DB)
		c.Next()
	}
}
