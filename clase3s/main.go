package main

import (
	"errors"
	"fmt"
	"log"
)

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	alimentoNecesario, err := animal("conejo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("La cantidad necesaria de alimento es de: %+v\n gramos", alimentoNecesario(20))

}

func animal(mascota string) (func(int) int, error) {
	switch mascota {
	case perro:
		return comidaPerros, nil
	case gato:
		return comidaGatos, nil
	case hamster:
		return comidaHamsters, nil
	case tarantula:
		return comidaTaratula, nil
	default:
		return nil, errors.New("El animal no se encuentra en la lista de mascotas")
	}
}

func comidaPerros(cantidad int) int {
	return cantidad * 10000
}

func comidaGatos(cantidad int) int {
	return cantidad * 5000
}

func comidaHamsters(cantidad int) int {
	return cantidad * 250
}

func comidaTaratula(cantidad int) int {
	return cantidad * 150
}
