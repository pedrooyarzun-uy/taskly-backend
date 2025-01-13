package helpers

import (
	"fmt"
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
