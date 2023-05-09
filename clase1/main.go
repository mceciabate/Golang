package main

import "fmt"

func main() {
	descuento := calcularDescuento(100.0, 0.25)
	fmt.Println("El descuento del producto es de $:", descuento)
	// caso := calculadoraPrestamos(23, false, 2, 20)
	// fmt.Println(caso)
	// sueldo := calcularInteres(200000)
	// fmt.Print(sueldo)
	caso1 := calculadoraPrestamosII(evaluarAptitud(22, true, 2), 10000)
	fmt.Println("------------------------")
	fmt.Println(caso1)

}

func calcularDescuento(a float32, b float32) float32 {
	return a * b
}

/* func calculadoraPrestamos(edad int, empleado bool, antiguedad int, ingreso int) (aptitud string) {
	if edad > 22 && empleado == true && antiguedad > 1 {
		aptitud = "Apto para recibir un prestamo"
	} else {
		aptitud = "No cumple los requisitos para el prestamo"
	}
	return aptitud
}
*/
/* func calcularInteres(ingreso int) (interes string) {
	if ingreso >= 100000 {
		interes = "No abonará interés"
	} else {
		interes = "Debe abonar intereses"
	}
	return interes
} */

func evaluarAptitud(edad int, empleado bool, antiguedad int) (aptitud bool) {
	if edad >= 22 && empleado == true && antiguedad > 1 {
		aptitud = true
	} else {
		aptitud = false
	}
	return aptitud
}

func calculadoraPrestamosII(datos bool, ingreso int) (apto string) {
	switch {
	case datos == true && ingreso >= 100000:
		apto = "Puede recibir el préstamo y no abonará intereses"
	case datos == true && ingreso < 100000:
		apto = "Puede recibir el préstamo y abonará con intereses"
	case datos == false:
		apto = "No cumple los requisitos para acceder a un préstamo"
	}
	return apto
}
