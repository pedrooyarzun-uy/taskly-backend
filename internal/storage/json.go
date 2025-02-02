package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"todo-app/internal/models"
)

const FILE_NAME = "data.json"
const PATH = "../../internal/storage/" + FILE_NAME

func CheckFileExistance() error {
	file, err := os.OpenFile(PATH, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	status, err := file.Stat()

	if err != nil {
		return err
	}

	if status.Size() == 0 {
		_, err := file.Write([]byte("[]"))

		if err != nil {
			return err
		}
	}

	return nil

}

func LoadTasks() []models.Task {

	var tasks []models.Task

	//Checks if file exists
	existance := CheckFileExistance()

	if existance != nil {
		panic(existance)
	}

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
