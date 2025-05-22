package routes

import (
	"github.com/gin-gonic/gin"
	"todo-app/internal/dto"
	"todo-app/internal/service"
)

func RegisterUserRoutes(r *gin.RouterGroup, s service.UserService) {
	r.POST("/sign-up", func(c *gin.Context) {

		var req dto.CreateUserRequest

		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err = s.CreateUser()

	})
}
