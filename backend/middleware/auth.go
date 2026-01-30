package middleware

import (
	"github.com/ApexPlayground/Linkkit/util"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{"error": "no token provided"})
		c.Abort()
		return
	}

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
