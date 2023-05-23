package main

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
