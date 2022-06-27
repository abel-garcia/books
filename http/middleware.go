package http

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wackGarcia/books/utils/cryptography"
)

func Authorization(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "API token required"})
		return
	}

	token = strings.Split(token, " ")[1]
	jwt, claims, err := cryptography.DecodeJWTSigned(token)
	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
		return
	}

	if !jwt.Valid {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
		return
	}

	ctx.Set("X-USER-ID", claims.Subject)
	ctx.Next()
}
