package middleware

import (
	"net/http"
	"strings"
	"urlshortner/internal/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenGenerator token.TokenGenerator) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")

		_, err := tokenGenerator.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}
		c.Next()
	}
}
