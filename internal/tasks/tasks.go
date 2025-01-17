package tasks

var Tasks = []Task{}

func CreateTask(title, description string) {

	newTask := Task{
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

func UpdateTask(id int) {
	for idx, val := range Tasks {
		if val.Id == id {
			Tasks[idx].Completed = true
		}
	}
}

func DeleteTask(id int) {
	for idx, val := range Tasks {
		if val.Id == id {
			Tasks = append(Tasks[:idx], Tasks[idx+1:]...)
			break
		}
	}
}
