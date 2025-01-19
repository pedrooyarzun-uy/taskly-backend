package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"todo-app/internal/helpers"
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
			helpers.CreateTaskMenu(scanner)
		case "2":
			helpers.ChangeStatusOfTask(scanner)
		case "3":
			helpers.DeleteTask(scanner)
		default:
			helpers.ConsoleCleaner()
			fmt.Println("La opción no es valida. Será redirigido al menu para continuar")
			time.Sleep(2 * time.Second)
			helpers.ConsoleCleaner()
			continue
		}

	}
}
