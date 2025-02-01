package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"todo-app/internal/models"
)

const FILE_NAME = "data.json"
const PATH = "../../internal/storage/" + FILE_NAME

func LoadTasks() []models.Task {

	var tasks []models.Task

	data, err := os.ReadFile(PATH)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &tasks)

	if err != nil {
		fmt.Println("No se pudo parsear la data")
		panic(err)
	}

	return tasks
}

func SaveData(tasks []models.Task) bool {
	updatedJSON, _ := json.MarshalIndent(tasks, "", "	")

	err := os.WriteFile(PATH, updatedJSON, 0644)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true

}
