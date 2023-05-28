package main

import "fmt"

/*Calcular precio
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de
Productos, Servicios y Mantenimientos. Debido a la fuerte demanda, y para optimizar la velocidad,
requieren que el cálculo de la sumatoria se realice en paralelo mediante tres goroutines.
Se requieren tres estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.
Y se requieren tres funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio por cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio por media hora trabajada.
	Si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.
Las tres se deben ejecutar concurrentemente y, al final, se debe mostrar por pantalla el monto final (sumando el total de las tres).
*/

type Productos struct {
	nombre   string
	precio   float64
	cantidad int
}

func sumarProductos(a []Productos, ch *chan float64) {
	var acumulador float64
	for _, v := range a {
		acumulador += v.precio * float64(v.cantidad)
	}
	*ch <- acumulador
}

type Servicios struct {
	nombre           string
	precio           float64
	minuosTrabajados int
}

func sumarServicios(a []Servicios, ch chan float64) {
	var acumulador float64
	for _, v := range a {
		switch {
		case (v.minuosTrabajados > 0 && v.minuosTrabajados <= 30):
			acumulador += v.precio

		case (v.minuosTrabajados > 30):
			acumulador += v.precio * float64(v.minuosTrabajados) / 30
		}
		ch <- acumulador
	}
	ch <- 0
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func sumarMantenimiento(a []Mantenimiento, ch chan float64) {
	var acumulador float64
	for _, v := range a {
		acumulador += v.precio
	}
	ch <- acumulador
}

//INSTANCIO LOS PRODUCTOS

var p1 = Productos{
	nombre:   "producto 1",
	precio:   1000.0,
	cantidad: 10,
}
var p2 = Productos{
	nombre:   "producto 2",
	precio:   1500.0,
	cantidad: 2,
}
var m1 = Mantenimiento{
	nombre: "mantenimiento 1",
	precio: 10000.0,
}

var m2 = Mantenimiento{
	nombre: "mantenimiento 2",
	precio: 5000.0,
}
var s1 = Servicios{
	nombre:           "servicio 1",
	precio:           7000.0,
	minuosTrabajados: 60,
}

//INSTANCIO LAS LISTAS
var listadoProductos = []Productos{
	p1,
	p2,
}
var listadoServicios = []Servicios{
	s1,
}
var listadoMantenimiento = []Mantenimiento{
	m1,
	m2,
}

func main() {

	ch := make(chan float64)
	ch1 := make(chan float64)
	ch2 := make(chan float64)

	go sumarProductos(listadoProductos, &ch)
	go sumarMantenimiento(listadoMantenimiento, ch1)
	go sumarServicios(listadoServicios, ch2)

	totalProductos, totalMantenimiento, totalServicios := <-ch, <-ch1, <-ch2
	total := totalProductos + totalMantenimiento + totalServicios
	fmt.Println(total)
}

//ch <- v    // Send v to channel ch.
//v := <-ch  // Receive from ch, and assign value to v.
