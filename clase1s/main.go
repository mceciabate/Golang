package main

import "fmt"

func main() {

	var alumnos = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Println(alumnos)
	alumnos = append(alumnos, "Gabriela")
	fmt.Println("-------------------------")
	fmt.Println(alumnos)
	fmt.Println("-------------------------")
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var empleadosMayoresDeEdad int = funcContador(employees)
	fmt.Println(empleadosMayoresDeEdad)
	employees["Federico"] = 25
	fmt.Println("-------------------------")
	fmt.Print("Lista empleados actualizada: ", employees, "\n")
	delete(employees, "Pedro")
	fmt.Println("-------------------------")
	fmt.Print("Eliminando empleado: ", employees, "\n")

}

func funcContador(lista map[string]int) int {

	acumulador := 0
	for _, v := range lista {
		if v > 21 {
			acumulador++
		}

	}
	return acumulador
}
