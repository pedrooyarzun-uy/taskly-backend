package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
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
	api := r.Group("/api")
	routes.RegisterUserRoutes(api, s)

	r.Run()

}
