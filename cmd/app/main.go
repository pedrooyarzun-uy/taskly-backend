package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"todo-app/internal/db"
	"todo-app/internal/repository"
	"todo-app/internal/routes"
	"todo-app/internal/service"
)

func main() {

	envPath := filepath.Join("../../", ".env")
	err := godotenv.Load(envPath)

	if err != nil {
		log.Fatal(".env variables could't load", err)
	}

	db.Init()

	repo := repository.NewUserRepository(db.DB)
	s := service.NewUserService(repo)

	r := gin.Default()

	origins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")

	r.Use(cors.New(cors.Config{

		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	routes.RegisterUserRoutes(api, s)

	r.Run()

}
