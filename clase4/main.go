package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Animes struct {
		Anime  string
		Anime2 string
		Anime3 string
	}

	type Persona struct {
		Nombre   string `json: "nombre", omitempty`
		Apellido string `json: "-"`
		Edad     int    `json: "edad"`
		Gustos   Animes
	}

	p1 := Persona{
		Nombre:   "Cecilia",
		Apellido: "Abate",
		Edad:     30,
		Gustos:   Animes{"Dororo", "FMA", "Bleach"},
	}

	personita := p1
	myJson, err := json.Marshal(personita)

	fmt.Printf("La persona se llama: %v\n y su anime favorito es %v\n", p1.Nombre, p1.Gustos.Anime)
	fmt.Println("-------------------------------------")
	fmt.Println(string(myJson))
	fmt.Println("-------------------------------------")
	fmt.Println(err)

}
