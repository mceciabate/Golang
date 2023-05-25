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
	Insertar al menos cuatro productos.*/
//PARTE II
/*Generar otro archivo CSV, llamado estimaciones.csv. Este tendrá los resultados de la suma de todos los productos de cada una de las categorías.
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
	CantidadActual: 2,
}
var p2 = Producto{
	Codigo:         2,
	Nombre:         "Lavarropas",
	PrecioActual:   200000.00,
	CantidadActual: 1,
}
var p3 = Producto{
	Codigo:         3,
	Nombre:         "Mesa",
	PrecioActual:   35000.00,
	CantidadActual: 2,
}

var p4 = Producto{
	Codigo:         4,
	Nombre:         "Sillon",
	PrecioActual:   75000.00,
	CantidadActual: 1,
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

func (lp ListaProductos) calculadoraEstimaciones() (float64, error) {
	var acumulador float64
	for _, v := range lp {
		acumulador += float64(v.CantidadActual) * (v.PrecioActual)
	}
	return acumulador, nil
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
	//ESTIMACIONES
	estimacionMuebles, err := ListaMuebles.calculadoraEstimaciones()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(estimacionMuebles)
	estimacionHerramientas, err := ListaHerramientas.calculadoraEstimaciones()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(estimacionHerramientas)
	estimacionElectro, err := ListaElectro.calculadoraEstimaciones()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(estimacionElectro)

	//CONVERTIR A STRING LOS FLOAT RETORNADOS
	stringElectro := fmt.Sprintf("%v", estimacionElectro)
	stringMuebles := fmt.Sprintf("%v", estimacionMuebles)
	stringHerramientas := fmt.Sprintf("%v", estimacionHerramientas)
	fmt.Println("----------ESTIMACIONES AS STRING-------------")
	fmt.Println(stringElectro, stringMuebles, stringHerramientas)

	//ARMO Y CONVIERTO LAS LISTAS EN STRINGS
	listadoCompleto := []string{"001", "Electro", string(lista1), ";", "002", "Muebles", string(lista2), ";", "003", "Herram", string(lista3)}
	listadoAsString := strings.Join(listadoCompleto, " ")
	listadoEstimaciones := []string{"Categoria", "Estimados", ";", "Electro", stringElectro, ";", "Muebles", stringMuebles, ";", "Herram", stringHerramientas}
	listadoEstimacionesAsString := strings.Join(listadoEstimaciones, " ")

	fmt.Println("----------LISTADO AS STRING-------------")
	fmt.Println(listadoAsString)
	fmt.Println(listadoEstimacionesAsString)

	//CREO EL ARCHIVO PARA CATEGORIAS
	f, err := os.Create("categorias.csv")
	if err != nil {
		log.Fatal(err)
	}

	//CREO EL ARCHIVO PARA ESTIMACIONES
	f2, err := os.Create("estimaciones.csv")
	if err != nil {
		log.Fatal(err)
	}

	//ESCRIBO EL ARCHIVO CATEGORIAS
	if _, err = f.WriteString(listadoAsString); err != nil {
		log.Fatal(err)
	}

	//ESCRIBO EL ARCHIVO ESTIMACIONES
	if _, err = f2.WriteString(listadoEstimacionesAsString); err != nil {
		log.Fatal(err)
	}

	//LEO EL ARCHIVO CATEGORIAS
	rawData, err := os.ReadFile("./categorias.csv")
	if err != nil {
		log.Fatal(err)
	}

	//LEO EL ARCHIVO ESTIMACIONES
	rawDataEstimacion, err := os.ReadFile("./estimaciones.csv")
	if err != nil {
		log.Fatal(err)
	}

	//PASO A STRING LAS LECTURAS OBTENIDAS
	data := strings.Split(string(rawData), "; ")
	dataEstimacion := strings.Split(string(rawDataEstimacion), "; ")
	fmt.Println("--------DATOS OBTENIDOS DE LAS LECTURAS--------------")
	fmt.Println(data, "\n", dataEstimacion)

	fmt.Println("-----------------FOR---------------------------")

	fmt.Println("--------TABLA CATEGORIAS---------")
	for _, v := range data {
		line := strings.Split(v, " ; ")
		line2 := strings.Split(line[0], " ")
		fmt.Printf("%s\t%s\t%s\t\n", line2[0], line2[1], line2[2])
	}

	fmt.Println("--------TABLA ESTIMACIONES---------")
	for _, v := range dataEstimacion {
		line := strings.Split(v, " ; ")
		line2 := strings.Split(line[0], " ")
		fmt.Printf("%s\t\t\t%s\n", line2[0], line2[1])

	}

}
