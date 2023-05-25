package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

type Animes struct {
	Anime  string
	Anime2 string
	Anime3 string
}
type Persona struct {
	Nombre   string `json:"nombre,omitempty"`
	Apellido string `json:"-"`
	Edad     int    `json:"edad"`
	Gustos   Animes
}

type Circulo struct {
	Radio float32
}

type Vehiculo struct {
	km     float64
	tiempo float64
}

type Auto struct {
	v Vehiculo
}

func main() {

	alumno1 := Alumno{
		Nombre:   "Cecilia",
		Apellido: "Abate",
		DNI:      "1234",
		Fecha:    "13/01/1998",
	}

	p1 := Persona{
		Nombre:   "Cecilia",
		Apellido: "Abate",
		Edad:     30,
		Gustos:   Animes{"Dororo", "FMA", "Bleach"},
	}

	c1 := Circulo{
		Radio: 30,
	}

	personita := p1
	myJson, err := json.Marshal(personita)

	fmt.Printf("La persona se llama: %v\n y su anime favorito es %v\n", p1.Nombre, p1.Gustos.Anime)
	fmt.Println("-------------------------------------")
	fmt.Println(string(myJson))
	fmt.Println("-------------------------------------")
	fmt.Println(err)
	fmt.Println("-------------------------------------")
	personita.descripcion()
	fmt.Println("-------------------------------------")
	fmt.Printf("El área del círculo es de : %v", c1.area())
	fmt.Println("-------------------------------------")
	alumno1.detalle()

}

func (p *Persona) descripcion() {
	fmt.Printf("La persona llamada %v tiene %v años", p.Nombre, p.Edad)
}

func (c Circulo) area() float32 {
	return math.Pi * c.Radio * c.Radio
}

func (a Alumno) detalle() {
	fmt.Printf("El alumno se apellida %v y su DNI es %v", a.Apellido, a.DNI)
}
