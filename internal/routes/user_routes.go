package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todo-app/internal/service"
)

func RegisterUserRoutes(r *gin.RouterGroup, s service.UserService) {
	r.POST("/sign-up", func(c *gin.Context) {
		fmt.Println(c.ShouldBind())
	})
}
