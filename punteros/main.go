package main

import "fmt"

func main() {
	var x int = 42
	var p *int // Declaramos un puntero p que apunta a un entero

	p = &x // Asignamos la dirección de memoria de x al puntero p

	fmt.Println("Valor de x:", x)
	fmt.Println("Dirección de memoria de x:", &x)
	fmt.Println("Valor de p:", p)
	fmt.Println("Valor al que apunta p:", *p)

	num := 5
	fmt.Println("Valor original de num:", num)

	increment(&num)
	fmt.Println("Valor de num después de incrementar:", num)

	j := 10
	y := 20

	fmt.Println("Valores antes de intercambiar:")
	fmt.Println("j:", j)
	fmt.Println("y:", y)

	swap(&j, &y)

	fmt.Println("Valores después de intercambiar:")
	fmt.Println("j:", j)
	fmt.Println("y:", y)

	numbers := [3]int{10, 20, 30}
	q := &numbers[0]

	fmt.Println("Array original:", numbers)

	*q = 42 // Modificamos el primer elemento del array usando el puntero p
	fmt.Println("Array después de modificar el primer elemento:", numbers)

	numbers1 := [3]int{1, 2, 3}
	fmt.Println("Array original:", numbers1)

	incrementElements(&numbers1)
	fmt.Println("Array después de incrementar los elementos:", numbers1)
}

func increment(x *int) {
	*x = *x + 1
}
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func incrementElements(arr *[3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i]++
	}
}
