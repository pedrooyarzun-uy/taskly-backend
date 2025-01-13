package main

import (
	"fmt"
	"time"
	"todo-app/helpers"
)

func main() {
	for {
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
			fmt.Print("Prueba")
			time.Sleep(2 * time.Second)

		default:
			helpers.ConsoleCleaner()
			fmt.Println("La opción no es valida. Será redirigido al menu para continuar")
			time.Sleep(2 * time.Second)
			helpers.ConsoleCleaner()
			continue
		}

	}
}
