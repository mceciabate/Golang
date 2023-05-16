package main

import (
	"errors"
	"fmt"
	"log"
)

/*Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products, instanciado con valores.
Dos métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
Una función getById() a la cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
Ejecutar al menos una vez cada método y función definidos desde main().*/

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
	Category    string
}

var Products []Product

func (p Product) save() {
	Products = append(Products, p)
}

func (p Product) getAll() {
	for _, p := range Products {
		fmt.Println(p)
	}
}

func getById(id int) (Product, error) {
	for _, p := range Products {
		if p.ID == id {
			return p, nil
		}

	}
	return Product{}, errors.New("No se encuentra el producto solicitado")
}

func main() {

	p1 := Product{
		1,
		"p1",
		2.5,
		"un producto",
		"categoria 1",
	}
	p2 := Product{
		2,
		"p2",
		3.5,
		"otro producto",
		"categoria 1",
	}

	p1.save()
	p2.save()
	p1.getAll()
	encontrarp, err := getById(7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(encontrarp)

}
