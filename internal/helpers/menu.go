package helpers

import (
	"bufio"
	"fmt"
	"time"
	"todo-app/internal/tasks"
)

func Menu() {
	fmt.Println(GREEN, "---Bienvenido a TODO-APP 😃---")
	fmt.Println("Seleccione una de las opciones para continuar:")
	fmt.Println("1. Crear una nueva tarea")
	fmt.Println("2. Modificar el estado de una tarea")
	fmt.Println("3. Eliminar una tarea")
	fmt.Println("4. Listar todas las tareas existentes", RESET)
	fmt.Println("------------------")
	fmt.Println("Ingrese su opcion: ")
}

func CreateTaskMenu(scanner *bufio.Reader) {
	ConsoleCleaner()
	fmt.Println("----Creacion de tareas----")

	fmt.Println("Ingrese un titulo para la tarea: ")
	title, _ := scanner.ReadString('\n')

	fmt.Println("Ingrese una descripción para la tarea: ")
	description, _ := scanner.ReadString('\n')

	//Task added
	tasks.CreateTask(title, description)

	ConsoleCleaner()

	fmt.Println(GREEN, "Tarea guardada exitosamente 😃", RESET)
	time.Sleep(2 * time.Second)
	ConsoleCleaner()
}

func ChangeStatusOfTask(scanner *bufio.Reader) {
	ConsoleCleaner()
	fmt.Println("----Modificar estado de la tarea----")

	if len(tasks.Tasks) == 0 {
		ConsoleCleaner()
		fmt.Println("No existen tares cargadas en el sistema")
		fmt.Println("Regresando al menú...")
		time.Sleep(2 * time.Second)
		return
	}

	for _, val := range tasks.Tasks {
		fmt.Println("Id: ", val.Id, ". Título: ", val.Title)
	}

	fmt.Println("--------------------")
	fmt.Println("Seleccione una opción: ")

	var option int

	_, err := fmt.Scanln(&option)

	if err != nil {
		ConsoleCleaner()
		fmt.Println("La opción ingresada no es correcta")
		return
	}

	done := tasks.UpdateTask(option)

	if done {
		ConsoleCleaner()
		fmt.Println("Se modificó la tarea exitosamente!")
		fmt.Println(GREEN, "Tarea modificada exitosamente! 😃", RESET)
		fmt.Println(tasks.Tasks)
		time.Sleep(2 * time.Second)
		ConsoleCleaner()
	} else {
		ConsoleCleaner()
		fmt.Println(RED, "La tarea ingresada no existe en el sistema", RESET)
		fmt.Println("Redirigiendo al menú principal")
		time.Sleep(2 * time.Second)
		ConsoleCleaner()
	}

}
