package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"todo-app/internal/models"
)

const FILE_NAME = "data.json"

func LoadTasks() []models.Task {

	basePath := "../../internal/storage/"

	fullPath := filepath.Join(basePath, FILE_NAME)

	var tasks []models.Task

	data, err := os.ReadFile(fullPath)
	if err != nil {
		dir, _ := os.Getwd()
		fmt.Print("VERGAAAAAAAAA", dir)
		panic(err)
	}

	err = json.Unmarshal(data, &tasks)

	if err != nil {
		fmt.Println("No se pudo parsear la data")
		panic(err)
	}

	return tasks
}
