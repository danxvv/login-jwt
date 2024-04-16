package middleware

import (
	"github.com/gin-gonic/gin"
	"login-user/pkg/jwtHelper"
	"strings"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := jwtHelper.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("token", token)
		claims := jwtHelper.GetTokenClaims(token)
		c.Set("userID", claims["user_id"])
		c.Next()
	}
}
