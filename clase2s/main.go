package main

import (
	"errors"
	"fmt"
	"log"
)

/* Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
Categoría C: su salario es de $1.000 por hora.
Categoría B: su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A: su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
*/

func main() {

	// categoriaA := "A"
	// categoríaB := "B"
	// categoriaC := "C"

	haberes, e := calcularSalario("desempleado", 2400)
	if e != nil {
		log.Fatal("Fallo de en la categoría ingresada")
	}
	fmt.Println("El salario es de $: ", haberes)

}

func calcularSalario(categoria string, minutosTrabajados float32) (float32, error) {
	horas := minutosTrabajados / 60.0
	var salario float32
	switch categoria {
	case "A":
		salario = (horas * 3000.0)
		porcentaje := salario * 0.54
		return salario + porcentaje, nil
	case "B":
		salario = horas * 1500.0
		porcentaje := salario * 0.2
		return salario + porcentaje, nil
	case "C":
		salario = horas * 1000
		return salario, nil
	default:
		return 0, errors.New("Categoría inválida")
	}
}
