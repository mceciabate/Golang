package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Soy Cecilia ejecutando Golang")
	//fmt.Scanf("Soy Cecilia ejecutando go")
	var cecilia string
	cecilia = "Cecilia"
	horas := 20
	const status string = "ok"
	var a [2]string
	a[0] = "Hello"
	a[1] = "Word"
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	myMap := make(map[string]string)
	myMap["Nombre"] = "Cecilia"
	myMap["Apellido"] = "Abate"
	myMap["Edad"] = "35"
	fmt.Println(cecilia, horas, status, a, s, myMap)
	delete(myMap, "Edad")
	fmt.Println(myMap["Apellido"], cap(s))
	c := sum(1, 1)
	d := resta(9, 2)
	fmt.Println(c, d)
}

func sum(a int, b int) int {
	suma := a + b
	return suma
}

func resta(a int, b int) int {
	return a - b
}
