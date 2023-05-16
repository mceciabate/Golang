package main

import (
	"errors"
	"fmt"
	"log"
)

type Producto struct {
	Id        int
	Nombre    string
	Precio    float32
	Categoria string
}

var p1 = Producto{
	1,
	"primer producto",
	2.5,
	"una categoria",
}

var p2 = Producto{
	2,
	"segundo producto",
	2.5,
	"una categoria",
}
var p3 = Producto{
	3,
	"tercer producto",
	2.5,
	"una categoria",
}

var p4 = Producto{
	4,
	"cuarto producto",
	7.8,
	"una categoria",
}

type Productos []Producto

var ListadoProductos = Productos{
	p1,
	p2,
	p3,
}

func (p Productos) listar() {
	for _, v := range p {
		fmt.Println(v)
	}
}

func (p Productos) guardar(a Producto) {
	p = append(p, a)
}

func (p Productos) buscarPorId(id int) (Producto, error) {
	for _, v := range p {
		if v.Id == id {
			return v, nil
		}
	}
	return Producto{}, errors.New("No se encuentra el producto especificado")
}

func main() {

	ListadoProductos.listar()
	ListadoProductos.guardar(p4)
	fmt.Println("---------------------")
	ListadoProductos.listar()
	encontrarProducto, e := ListadoProductos.buscarPorId(9)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("---------------------")
	fmt.Println(encontrarProducto)

}
