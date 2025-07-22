package routes

import (
	"todo-app/internal/dto"
	"todo-app/internal/middlewares"
	"todo-app/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterCategoriesRoutes(r *gin.RouterGroup, s service.CategoryService) {
	r.Use(middlewares.VerifyJWT())

	r.POST("/create-category", func(ctx *gin.Context) {
		var req dto.CreateCategoryRequest

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

		err := s.CreateCategory(req, userID)

		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"message": "ok"})
	})
}
