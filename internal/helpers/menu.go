package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"todo-app/internal/tasks"
)

func Menu() {
	tasks.Init()
	fmt.Println(YELLOW, "---Bienvenido a TODO-APP 😃---", RESET)
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

	fmt.Println(GREEN, "Tarea guardada exitosamente! 😃", RESET)
	time.Sleep(2 * time.Second)
	ConsoleCleaner()
}

func ChangeStatusOfTask(scanner *bufio.Reader) {
	ConsoleCleaner()

	if !tasks.HasPendingTasks() {
		fmt.Println(YELLOW, "No existen tares cargadas en el sistema", RESET)
		fmt.Println("Regresando al menú...")
		time.Sleep(2500 * time.Millisecond)
		return
	}

	fmt.Println("----Modificar estado de la tarea----")

	for _, val := range tasks.Tasks {

		if !val.Completed {
			fmt.Printf("ID: %d. Título: %s \n", val.Id, strings.TrimSpace(val.Title))
		}
	}

	fmt.Println("--------------------")
	fmt.Println("Ingrese el ID de la tarea: ")

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
		fmt.Println(GREEN, "Tarea pasada a finalizada con éxito! 😃", RESET)
		time.Sleep(2 * time.Second)
		ConsoleCleaner()
	} else {
		ConsoleCleaner()
		fmt.Println(RED, "La tarea ingresada no existe en el sistema ❌", RESET)
		fmt.Println("Redirigiendo al menú principal...")
		time.Sleep(2 * time.Second)
		ConsoleCleaner()
	}

}

func DeleteTask(scanner *bufio.Reader) {
	ConsoleCleaner()
	fmt.Println("----Eliminar tarea----")

	if len(tasks.Tasks) == 0 {
		ConsoleCleaner()
		fmt.Println(YELLOW, "No existen tares cargadas en el sistema ⚠️", RESET)
		fmt.Println("Redirigiendo al menú principal...")
		time.Sleep(2500 * time.Millisecond)
		return
	}

	for _, val := range tasks.Tasks {
		var option string

		if val.Completed {
			option = "Done"
		} else {
			option = "Doing"
		}

		fmt.Printf("ID: %d. Título: %s. Status: %s \n", val.Id, strings.TrimSpace(val.Title), option)
	}

	fmt.Println("--------------------")
	fmt.Println("Ingrese el ID de la tarea: ")

	var option int

	_, err := fmt.Scanln(&option)

	if err != nil {
		ConsoleCleaner()
		fmt.Println("La opción ingresada no es correcta")
		return
	}

	deleted := tasks.DeleteTask(option)

	if deleted {
		ConsoleCleaner()
		fmt.Println(GREEN, "Se eliminó la tarea exitosamente! 😃", RESET)
		time.Sleep(2 * time.Second)
		ConsoleCleaner()
	} else {
		ConsoleCleaner()
		fmt.Println(RED, "La tarea ingresada no existe en el sistema", RESET)
		fmt.Println("Redirigiendo al menú principal...")
		time.Sleep(2 * time.Second)
		ConsoleCleaner()
	}

}

func GetAllTasks() {
	ConsoleCleaner()

	if len(tasks.Tasks) == 0 {
		fmt.Println(YELLOW, "No existen tareas cargadas en el sistema ⚠️", RESET)
		fmt.Println("Redirigiendo al menú principal...")
		time.Sleep(2500 * time.Millisecond)
		return
	}

	fmt.Println("----Listar tareas----")

	for idx, val := range tasks.Tasks {

		var option string
		if val.Completed {
			option = "Done"
		} else {
			option = "Doing"
		}

		var color string
		if idx%2 == 0 {
			color = MAGENTA
		} else {
			color = BLUE
		}

		fmt.Printf("%s ID: %d. Título: %s. Descripción: %s. Status: %s %s \n", color, val.Id, strings.TrimSpace(val.Title), strings.TrimSpace(val.Description), option, RESET)
	}

	fmt.Println("--------------------")
	fmt.Println("Presione enter para continuar: ")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
