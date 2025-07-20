package middlewares

import (
	"os"
	"strconv"
	"todo-app/internal/helpers"

	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized request"})
			return
		}

		if token != os.Getenv("AUTHORIZATION_TOKEN") {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized request"})
			return
		}

		ctx.Next()
	}

}

func VerifyJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		claims, err := helpers.ParseJWT(token)

		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Session expired"})
			return
		}

		sub, ok := claims["sub"].(string)

		if !ok {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "There is something wrong"})
			return
		}

		userID, err := strconv.Atoi(sub)

		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Something went wrong"})
		}

		ctx.Set("userID", userID)

		ctx.Next()
	}
}
