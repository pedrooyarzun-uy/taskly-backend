package middlewares

import (
	"fmt"
	"os"
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
			fmt.Println("Tamo activo mi gente")
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized request"})
			return
		}

		ctx.Next()
	}

}

func VerifyJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		isValid := helpers.VerifyJWT(token)

		if !isValid {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Session expired"})
			return
		}

		ctx.Next()

	}
}
