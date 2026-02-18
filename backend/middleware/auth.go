package middleware

import (
	"strings"

	"github.com/ApexPlayground/Linkkit/util"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"error": "no token provided"})
		c.Abort()
		return
	}

	// Strip "Bearer " prefix
	token := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := util.ParseJWT(token)
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}

	c.Set("user_id", claims.UserID)
	c.Set("is_admin", claims.IsAdmin)
	c.Next()
}
