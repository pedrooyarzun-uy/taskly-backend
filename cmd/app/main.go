package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
	"todo-app/internal/db"
	"todo-app/internal/middlewares"
	"todo-app/internal/repository"
	"todo-app/internal/routes"
	"todo-app/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	envPath := filepath.Join("../../", ".env")
	var err error

	if os.Getenv("RAILWAYS_ENVIRONMENT") == "" {
		err = godotenv.Load(envPath)
	}

	if err != nil {
		log.Fatal(".env variables could't load", err)
	}

	db.Init()

	ur := repository.NewUserRepository(db.DB)
	vr := repository.NewVerificationRepository(db.DB)
	s := service.NewUserService(ur, vr)

	r := gin.Default()

	r.Use(cors.New(cors.Config{

		AllowOrigins:     []string{os.Getenv("ALLOWED_ORIGINS")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(middlewares.VerifyToken())

	r.SetTrustedProxies(nil)

	api := r.Group("/api")
	routes.RegisterUserRoutes(api, s)

	r.Run(":" + os.Getenv("GIN_PORT"))

}
