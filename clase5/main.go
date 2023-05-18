package main

/*Algunas tiendas de e-commerce necesitan realizar una funcionalidad en Go para
administrar productos y retornar el valor del precio total. La empresa tiene tres tipos de
productos: Pequeño, Mediano y Grande. Pero se espera que sean muchos más.
Los costos adicionales para cada uno son:
● Pequeño: solo tiene el costo del producto.
● Mediano: el precio del producto + un 3% por mantenerlo en la tienda.
● Grande: el precio del producto + un 6% por mantenerlo en la tienda y un adicional
de $2500 de costo de envío.
El porcentaje de mantenerlo en la tienda se calcula sobre el precio del producto.
Se requiere una función factory que reciba el tipo de producto y el precio, y retorne una
interfaz Producto que tenga el método Precio.
Se debe poder ejecutar el método Precio y que el método devuelva el precio total
basándose en el costo del producto y los adicionales, en caso de que los tenga.
*/

import "fmt"

const (
	smallType  = "small"
	mediumType = "medium"
	largeType  = "large"
)

type Large struct {
	PriceBase float32
}

func (l Large) calcular() float32 {
	return l.PriceBase*0.6 + l.PriceBase + 2500.0
}

type Medium struct {
	PriceBase float32
}

func (m Medium) calcular() float32 {
	return m.PriceBase*0.3 + m.PriceBase
}

type Small struct {
	PriceBase float32
}

func (s Small) calcular() float32 {
	return s.PriceBase
}

type Producto interface {
	calcular() float32
}

func Factory(productoType string, precioB float32) Producto {
	switch {
	case productoType == "small":
		return Small{PriceBase: precioB}
	case productoType == "medium":
		return Medium{PriceBase: precioB}
	case productoType == "large":
		return Large{PriceBase: precioB}
	}
	return nil
}

func main() {

	p1 := Factory("small", 200.0)
	p2 := Factory("medium", 200.0)
	p3 := Factory("large", 200.0)
	fmt.Println(p1)
	fmt.Println("Costo final del producto 1: ", p1.calcular())
	fmt.Println("--------------------")
	fmt.Println("Costo final del producto 2: ", p2.calcular())
	fmt.Println("--------------------")
	fmt.Println("Costo final del producto 3: ", p3.calcular())

}
