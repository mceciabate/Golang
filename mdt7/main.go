package main

/*En la función main, definir una variable salary y asignarle un valor de tipo “int”. Luego, crear un error personalizado con un struct que implemente Error() con el mensaje “Error: el salario ingresado no alcanza el mínimo imponible” y lanzarlo en caso de que salary sea menor a 150.000. De lo contrario, imprimir por consola el mensaje “Debe pagar impuesto”.
 */
/*Repetir el proceso de la ejercitación realizada en clase, pero ahora implementando fmt.Errorf() para que el mensaje de error reciba por parámetro el valor de salary indicando que no alcanza el mínimo imponible. El mensaje mostrado por consola deberá decir: “Error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]” (siendo [salary] el valor de tipo int pasado por parámetro).
 */

import (
	"errors"
	"fmt"
)

var errSalario = errors.New("Salario muy bajo")

// var errSalary = fmt.Errorf("Salario es muy bajo %d", salary )
// var salary = 5000

func main() {
	var salario = 500
	err := validarSalario(salario)
	if errors.Is(err, errSalario) {
		fmt.Println(err)
	}
	var salary = 500
	err1 := validarSalarioII(salary)
	if err != nil {
		fmt.Println(err1)
	}
}

func validarSalario(salario int) error {
	if salario <= 10000 {
		return errSalario
	}
	return nil
}

func validarSalarioII(salary int) error {
	if salary <= 10000 {
		return fmt.Errorf("Error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	}
	return nil
}
