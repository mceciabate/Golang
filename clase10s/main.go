package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

		for _, v := range slice {
			num := v
			go func() {
				fmt.Println(num)
			}()
		}
	*/

	/*
		imprimirInicio()

		go func() {
			for {
				fmt.Println("hola")
				time.Sleep(2 * time.Second)
			}
		}()

		go func() {
			for {
				fmt.Println("mundo")
				time.Sleep(3 * time.Second)
			}
		}()

		fmt.Scanln()
	*/

	/*
		Una empresa de finanzas procesa órdenes de préstamos y órdenes de devoluciones en una proporción de dos a uno (es decir, por cada devolución dan dos préstamos). Aproximadamente, se procesa una devolución por segundo, y nos han solicitado que programemos un modelo de procesamiento concurrente que denote este modelo.
		Se requiere que corran al menos dos goroutines de forma concurrente que procesen estas órdenes. Donde además se muestre por la consola una devolución por cada dos préstamos hasta que se apriete cualquier botón y se acabe con la ejecución del programa. El output esperado debería ser similar a este:
		devolucion procesada
		prestamo procesado
		prestamo procesado
		devolucion procesada
		prestamo procesado
		prestamo procesado
	*/

	orden := make(chan string)
	devolucion := make(chan string)

	// Proceso las devoluciones
	go func() {
		for {
			time.Sleep(time.Second)
			// Proceso la devolución
			devolucion <- "devolución procesada"
		}
	}()

	// Proceso las ordenes
	go func() {
		for {
			time.Sleep(time.Second / 2)
			// Proceso la orden
			orden <- "orden procesada"
		}
	}()

	go func() {
		for ordenProcesada := range orden {
			fmt.Println(ordenProcesada)
		}
	}()

	go func() {
		for devoProcesada := range devolucion {
			fmt.Println(devoProcesada)
		}
	}()

	/*
		go func() {
			for {
				select {
				case msgOrden := <-orden:
					fmt.Println(msgOrden)
				case msgDevo := <-devolucion:
					fmt.Println(msgDevo)
				}
			}
		}()
	*/

	fmt.Scanln()
}

func procesarDevolucion(devolucion chan string) {
	for {
		time.Sleep(1 * time.Second)
		// Proceso la devolución
		devolucion <- "devolución procesada"
	}
}

func imprimirInicio() {
	fmt.Println("iniciando el programa")
	time.Sleep(1 * time.Second)
}
