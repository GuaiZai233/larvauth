package middleware

import (
	"github.com/GuaiZai233/larvauth/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get Authorization header
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "请求头中缺少 Authorization Token",
			})
			c.Abort()
			return
		}

		//separate string and auth format
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization Token 格式错误",
			})
			c.Abort()
			return
		}

		//get token string and resolve
		tokenString := parts[1]
		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token 无效或已过期",
			})
			c.Abort()
			return
		}

		//set user info to context
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}
