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
		Deleted:     false,
	}

	if len(Tasks) == 0 {
		newTask.Id = 1
	} else {
		newTask.Id = Tasks[len(Tasks)-1].Id + 1
	}

	Tasks = append(Tasks, newTask)
	//Save data into json
	storage.SaveData(Tasks)

}

func UpdateTask(id int) bool {
	for idx, val := range Tasks {
		if val.Id == id {
			Tasks[idx].Completed = true
			return true
		}
	}

	//Save data into json
	storage.SaveData(Tasks)
	return false
}

func DeleteTask(id int) bool {

	wasUpdated := false

	for idx, val := range Tasks {
		if val.Id == id {
			Tasks[idx].Deleted = true
			wasUpdated = true
		}
	}

	//Save data into json
	storage.SaveData(Tasks)
	return wasUpdated
}

func HasPendingTasks() bool {
	for _, val := range Tasks {
		if !val.Completed && !val.Deleted {
			return true
		}
	}

	return false
}
