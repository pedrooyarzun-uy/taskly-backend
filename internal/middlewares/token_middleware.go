package middlewares

import (
	"fmt"
	"os"

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
