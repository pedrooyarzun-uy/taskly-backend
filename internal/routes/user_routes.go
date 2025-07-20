package routes

import (
	"todo-app/internal/dto"
	"todo-app/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, s service.UserService) {
	r.POST("/sign-up", func(c *gin.Context) {

		var req dto.CreateUserRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err := s.CreateUser(req)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "ok"})
	})

	r.GET("/verify-user", func(c *gin.Context) {
		token := c.Query("token")

		err := s.VerifyUser(token)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "ok"})
	})

	r.POST("/sign-in", func(c *gin.Context) {
		var req dto.SignInRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		token, err := s.SignIn(req)

		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "ok",
			"token":   token,
		})

	})
}
