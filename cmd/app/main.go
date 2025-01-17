package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"todo-app/internal/helpers"
	"todo-app/internal/tasks"
)

func main() {
	for {

		//Scanner works for reading sentences not only one word
		scanner := bufio.NewReader(os.Stdin)

		helpers.ConsoleCleaner()
		helpers.Menu()

		var option string

		_, err := fmt.Scanln(&option)

		if err != nil {
			helpers.ConsoleCleaner()
			fmt.Println("No se pudo leer la opcion. Ingrese otra para continuar")
			time.Sleep(2 * time.Second)
			helpers.ConsoleCleaner()
			continue
		}

		switch option {
		case "1":
			helpers.ConsoleCleaner()

			fmt.Println("----Creacion de tareas----")

			fmt.Println("Ingrese un titulo para la tarea: ")
			title, _ := scanner.ReadString('\n')

			fmt.Println("Ingrese una descripción para la tarea: ")
			description, _ := scanner.ReadString('\n')

			//Task added
			tasks.CreateTask(title, description)

			helpers.ConsoleCleaner()

			fmt.Println(helpers.GREEN, "Tarea guardada exitosamente 😃", helpers.RESET)
			time.Sleep(2 * time.Second)
			helpers.ConsoleCleaner()
		case "2":
			helpers.ConsoleCleaner()
		default:
			helpers.ConsoleCleaner()
			fmt.Println("La opción no es valida. Será redirigido al menu para continuar")
			time.Sleep(2 * time.Second)
			helpers.ConsoleCleaner()
			continue
		}

	}
}
