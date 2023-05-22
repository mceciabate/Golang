package main

import "fmt"

//PAQUETES FMT

func main() {
	fmt.Println("Hola por la consola")
	nombre := "Cecilia"
	edad := 35
	fmt.Printf("La persona se llama %s, y tiene %d años \n", nombre, edad)
	var frase string = fmt.Sprintf("Hola %s %d", "mundo", 10)
	fmt.Println(frase)
	// var mensajeConsola string
	// fmt.Scan(&mensajeConsola)
	// fmt.Println("El mensaje recibido por consola es :", mensajeConsola)
	var nombre_completo, apellido string
	var telefono int
	fmt.Scanf("%s %s %d", &nombre_completo, &apellido, &telefono)
	fmt.Println(nombre_completo, apellido, telefono)

}
