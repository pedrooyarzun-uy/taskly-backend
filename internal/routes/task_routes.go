package routes

import (
	"todo-app/internal/dto"
	"todo-app/internal/middlewares"
	"todo-app/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(r *gin.RouterGroup, s service.TaskService) {

	r.Use(middlewares.VerifyJWT())

	r.POST("/create-task", func(ctx *gin.Context) {
		var req dto.CreateTaskRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		val, ok := ctx.Get("userID")

		if !ok {
			ctx.JSON(400, gin.H{"error": "Something went wrong"})
			return
		}

		userID, ok := val.(int)

		if !ok {
			ctx.JSON(400, gin.H{"error": "Something went wrong"})
			return
		}

		err := s.CreateTask(req, userID)

		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"message": "ok"})
	})

	r.POST("/complete-task", func(ctx *gin.Context) {
		var req dto.CompleteTaskRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		val, ok := ctx.Get("userID")

		if !ok {
			ctx.JSON(400, gin.H{"error": "Something went wrong"})
			return
		}

		userID, ok := val.(int)

		if !ok {
			ctx.JSON(400, gin.H{"error": "Something went wrong"})
			return
		}

		err := s.CompleteTask(req, userID)

		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"message": "ok"})
	})

	r.POST("/delete-task", func(ctx *gin.Context) {
		var req dto.DeleteTaskRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		val, _ := ctx.Get("userID")
		userID, ok := val.(int)

		if !ok {
			ctx.JSON(400, gin.H{"error": "something went wrong"})
			return
		}

		err := s.DeleteTask(req, userID)

		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"message": "ok"})

	})

	r.PUT("/modify-task", func(ctx *gin.Context) {
		var req dto.ModifyTaskRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.Category == 0 && req.Description == "" && req.Title == "" {
			ctx.JSON(400, gin.H{"error": "At least one of the fields must be modified"})
			return
		}

		val, _ := ctx.Get("userID")
		userID, ok := val.(int)

		if !ok {
			ctx.JSON(400, gin.H{"error": "something went wrong"})
			return
		}

		err := s.ModifyTask(req, userID)

		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"message": "ok"})
	})

	r.GET("/get-tasks", func(ctx *gin.Context) {
		val, _ := ctx.Get("userID")
		userID, ok := val.(int)

		if !ok {
			ctx.JSON(400, gin.H{"error": "something went wrong"})
			return
		}

		tasks, err := s.GetPendingTasks(userID)

		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "ok",
			"tasks":   tasks,
		})
	})

	r.GET("get-all-tasks", func(ctx *gin.Context) {
		val, _ := ctx.Get("userID")
		userID, ok := val.(int)

		if !ok {
			ctx.JSON(400, gin.H{"error": "Something went wrong"})
			return
		}

		tasks, err := s.GetAllTasks(userID)

		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "ok",
			"tasks":   tasks,
		})
	})
}
