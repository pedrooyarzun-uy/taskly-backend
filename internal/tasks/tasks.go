package tasks

import (
	"todo-app/internal/models"
	"todo-app/internal/storage"
)

var Tasks = []models.Task{}

// Loads data in json into array of the app
func Init() {
	Tasks = storage.LoadTasks()
}

func CreateTask(title, description string) {

	newTask := models.Task{
		Title:       title,
		Description: description,
		Completed:   false,
	}

	if len(Tasks) == 0 {
		newTask.Id = 1
	} else {
		newTask.Id = Tasks[len(Tasks)-1].Id + 1
	}

	Tasks = append(Tasks, newTask)

}

func UpdateTask(id int) bool {
	for idx, val := range Tasks {
		if val.Id == id {
			Tasks[idx].Completed = true
			return true
		}
	}

	return false
}

func DeleteTask(id int) bool {
	for idx, val := range Tasks {
		if val.Id == id {
			Tasks = append(Tasks[:idx], Tasks[idx+1:]...)
			return true
		}
	}

	return false
}

func HasPendingTasks() bool {
	for _, val := range Tasks {
		if !val.Completed {
			return true
		}
	}

	return false
}
