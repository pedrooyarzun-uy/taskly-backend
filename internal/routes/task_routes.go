package routes

import (
	"todo-app/internal/dto"
	"todo-app/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(r *gin.RouterGroup, s service.TaskService) {

	r.POST("/create-task", func(ctx *gin.Context) {
		var req dto.CreateTaskRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		userID := 51
		err := s.CreateTask(req, userID)

		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"message": "ok"})
	})
}
