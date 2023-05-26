package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*Imprimir datos
Un estudio contable necesita que generemos un archivo customers.txt con información de sus clientes. Ahora que el archivo sí existe, el panic no debe ser lanzado.
Creamos el archivo customers.txt y le agregamos la información de los clientes.
Extendemos el código del punto uno para que podamos leer este archivo e imprimir los datos que contenga. En el caso de no poder leerlo, se debe lanzar un panic.
Recordemos que siempre que termina la ejecución, independientemente del resultado, debemos tener un defer que nos indique que la ejecución finalizó. También recordemos cerrar los archivos al finalizar su uso.
¡Éxitos!
e*/

type Cliente struct {
	Nombre   string `json:"Nombre,omitempty"`
	Apellido string `json:"Apellido,omitempty"`
	Telefono string `json:"Telefono,omitempty"`
	Mail     string `json:"Mail,omitempty"`
}

type ListaClientes []Cliente

var cliente1 = Cliente{
	Nombre:   "Serena",
	Apellido: "Tsukino",
	Telefono: "123456",
	Mail:     "sailormoon@mail.com",
}

var cliente2 = Cliente{
	Nombre:   "Hichigo",
	Apellido: "Kurosaki",
	Telefono: "123456",
	Mail:     "bleach@mail.com",
}

var ListaClientes1 = ListaClientes{
	cliente1,
	cliente2,
}

func main() {
	lista1, err := json.Marshal(ListaClientes1)
	if err != nil {
		log.Fatal(err)
	}
	listaJson := (string(lista1))
	fmt.Println(listaJson)
	f, err := os.Create("customers.txt")
	if err != nil {
		log.Fatal(err)
	}

	if _, err = f.WriteString(listaJson); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover de este panic:", r)
			fmt.Println("¡Todo OK! Sigo funcionando...")
		}
	}()
	rawData, err := os.ReadFile("./customers.txt")
	if err != nil {
		panic(err)
	}

	dataRecover := string(rawData)
	fmt.Println(dataRecover)
	if err = f.Close(); err != nil {
		log.Fatal(err)
	}

}
