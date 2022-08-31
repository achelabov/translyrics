package middlewares

import (
	"github.com/achelabov/translyrics/auth"
	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")

	if tokenString == "" {
		ctx.JSON(401, gin.H{"error": "request does not contain an access token"})
		ctx.Abort()
		return
	}

	err := auth.ValidateToken(tokenString)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.Next()
}
