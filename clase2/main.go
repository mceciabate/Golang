package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(calcularSalario(3000.0))
	fmt.Println("---------------------")
	fmt.Println(calcularSalario(85000.0))
	fmt.Println("---------------------")
	fmt.Println(calcularSalario(200000.0))
	fmt.Println("---------------------")
	fmt.Println(calcularPromedio(3.0, 4.5, 10.0))
	fmt.Println("---------------------")
	fmt.Println(comprobarNumeros(0.5, 0.7, -0.2))
	fmt.Println("---------------------")
	fmt.Println(calcularPromedioII(5.0, 3.0, -4.5))

}

func calcularSalario(salario float32) (impuesto float32) {
	switch {
	case salario >= 50000.0:
		impuesto = salario * 0.17
	case salario >= 150000:
		impuesto = salario * 0.27
	default:
		impuesto = 0.0
	}
	return impuesto
}

/* func calcularPromedio (valores ...float32) (float32, error){
var longArray int = len(valores)
longitudDecimal := float32(longArray)
var promedio float32
var acumulador float32

for _, valor := range valores{
	if valor < 0 {
		return 0, errors.New("No se puede ingresar valores negativos")
	} else {
		acumulador += valor
		promedio = acumulador/longitudDecimal
	}
		} return promedio, nil
}
*/

func calcularPromedio(valores ...float32) (float32, error) {
	var promedio float32
	longArray := len(valores)
	logDecimal := float32(longArray)
	var acumulador float32
	for _, valor := range valores {
		acumulador += valor
	}
	promedio = acumulador / logDecimal
	return promedio, nil
}

func comprobarNumeros(numeros ...float32) bool {
	for _, numero := range numeros {
		if numero < 0 {
			return true
		}
	}
	return false
}

func calcularPromedioII(valores ...float32) (float32, error) {
	var promedio float32
	longArray := len(valores)
	logDecimal := float32(longArray)
	var acumulador float32
	switch {
	case comprobarNumeros(valores...) == false:
		{
			for _, valor := range valores {
				acumulador += valor
			}
			promedio = acumulador / logDecimal
			return promedio, nil
		}
	case comprobarNumeros(valores...) == true:
		{
			return 0.0, errors.New("No se pueden ingresar valores negativos")
		}
	}
	return 0, nil
}
