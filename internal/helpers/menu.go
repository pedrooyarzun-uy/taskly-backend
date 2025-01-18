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
