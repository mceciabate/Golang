package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

/*Ejercicio - Mayorista
Se necesita un software desarrollado en Go que imprima las estimaciones de lo que generarían, los productos de cada categoría en un mayorista, en ventas de productos para el hogar. Para eso se detallarán los pasos para conseguirlo:
Funcionalidad para generar un archivo CSV, llamado categorías.csv.
Cargar el archivo con las categorías. Ejemplo:
		001	Electrodomésticos	ListaProductos
		002	Muebles		ListaProductos
		003	Herramientas		ListaProductos
		004	Pinturas		ListaProductos
		005	Aberturas		ListaProductos
		006	Construcción		ListaProductos
		007	Automotor 		ListaProductos
		Etcétera…
	Elegir al menos tres categorías.

Generar para cada una de estas categorías los productos. Estos tendrán como información:
Código
Nombre
PrecioActual
CantidadActual
	Insertar al menos cuatro productos.
Generar otro archivo CSV, llamado estimaciones.csv. Este tendrá los resultados de la suma de todos los productos de cada una de las categorías.
Imprimir todos los estimativos por consola. Ejemplo:
Categoría			EstimativoPorCategoría
Construcción				60.700
Pinturas 				40.500
Aberturas				55.300
TotalEstimativo 			156.500
¡Éxitos!
*/

type Producto struct {
	Codigo         int     `json:"Cod,omitempty"`
	Nombre         string  `json:"Nombre,omitempty"`
	PrecioActual   float64 `json:"Precio,omitempty"`
	CantidadActual int     `json:"Stock,omitempty"`
}

type ListaProductos []Producto

var ListaElectro = ListaProductos{
	p1,
	p2,
}

var ListaMuebles = ListaProductos{
	p3,
	p4,
}

var ListaHerramientas = ListaProductos{
	p5,
	p6,
}

var p1 = Producto{
	Codigo:         1,
	Nombre:         "Pava Electrica",
	PrecioActual:   10000.00,
	CantidadActual: 20,
}
var p2 = Producto{
	Codigo:         2,
	Nombre:         "Lavarropas",
	PrecioActual:   200000.00,
	CantidadActual: 15,
}
var p3 = Producto{
	Codigo:         3,
	Nombre:         "Mesa",
	PrecioActual:   35000.00,
	CantidadActual: 35,
}

var p4 = Producto{
	Codigo:         4,
	Nombre:         "Sillon",
	PrecioActual:   75000.00,
	CantidadActual: 4,
}

var p5 = Producto{
	Codigo:         5,
	Nombre:         "Taladro",
	PrecioActual:   15000.00,
	CantidadActual: 30,
}
var p6 = Producto{
	Codigo:         6,
	Nombre:         "Martillo",
	PrecioActual:   3000.00,
	CantidadActual: 12,
}

func main() {
	//JSON
	lista1, err := json.Marshal(ListaElectro)
	if err != nil {
		log.Fatal(err)
	}
	lista2, err := json.Marshal(ListaMuebles)
	if err != nil {
		log.Fatal(err)
	}
	lista3, err := json.Marshal(ListaHerramientas)
	if err != nil {
		log.Fatal(err)
	}

	listadoCompleto := []string{"001", "Electrodomesticos", string(lista1), ";", "002", "Muebles", string(lista2), ";", "003", "Herramientas", string(lista3)}
	listadoAsString := strings.Join(listadoCompleto, " ")
	fmt.Println("----------LISTADO AS STRING-------------")

	fmt.Println(listadoAsString)

	//CREO EL ARCHIVO
	f, err := os.Create("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	//ESCRIBO EL ARCHIVO CON EL JSON
	if _, err = f.WriteString(listadoAsString); err != nil {
		log.Fatal(err)
	}

	//LEO EL ARCHIVO
	rawData, err := os.ReadFile("./data.csv")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(string(rawData), "; ")
	fmt.Println("---------SALIDA SPLIT--------------")
	fmt.Println(data)
	fmt.Println("-----------------FOR---------------------------")
	for _, v := range data {
		line := strings.Split(v, " ; ")
		fmt.Printf("%s\n", line[0])

	}
}
