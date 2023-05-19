package main

import (
	"errors"
	"fmt"
	"log"
)

/*Algunas tiendas de e-commerce necesitan realizar una funcionalidad en Go para administrar los diferentes métodos de pago que se ofrecen al momento de realizar una nueva compra. Actualmente, las empresas cuentan con 3 métodos de pago (pero se espera que sean muchos más): tarjeta de crédito, transferencia bancaria y efectivo. Cada método de pago, cuenta con su propia comisión por la utilización del servicio. Dichas comisiones son:
Tarjeta de crédito: cuenta con un 10% de recargo sobre el precio publicado.
Transferencia bancaria: no presenta recargo sobre el precio publicado.
Efectivo: cuenta con un 5% de descuento sobre el precio publicado.
Requerimientos:
Crear tres estructuras que representen los diferentes métodos de pagos permitidos.
Crear una estructura Purchase que tenga el precio y un arreglo de métodos de pagos disponibles para la compra. Deberemos crear una función que permita establecer el precio de la compra al momento de crear la misma.
Crear una interfaz PaymentMethod que defina el método Pay(purchase *purchase) a ser implementado por cada método de pago disponible.
Crear una función Process que nos permita procesar el pago de la compra recibiendo como parámetro el método de pago elegido.
Solicitar al usuario ingresar el método de pago seleccionado para la compra.
*/

const (
	c = "credit"
	t = "transfer"
	e = "cash"
)

type Purchase struct {
	precioCompra float64
	lista        [3]PaymentMethod
}

func Pay(valorCompra float64) (Purchase, error) {
	if valorCompra > 0 {
		return Purchase{
			precioCompra: valorCompra,
			lista:        ListaEstructural,
		}, nil
	}
	return Purchase{}, errors.New("El monto debe ser mayor a 0")
}

var ListaEstructural = [3]PaymentMethod{
	Credito{},
	Efectivo{},
	Transferencia{},
}

var listaPurchase, _ = Pay(300.0)

type PaymentMethod interface {
	comision(precio float64) float64
}

type Credito struct {
}

func (cred Credito) comision(precio float64) float64 {
	return precio * 1.10

}

type Efectivo struct {
}

func (efec Efectivo) comision(precio float64) float64 {
	return precio - precio*0.05

}

type Transferencia struct {
}

func (trans Transferencia) comision(precio float64) float64 {
	return precio

}

func Process(metodoPago string) (float64, error) {
	switch {
	case metodoPago == "credit":
		return listaPurchase.lista[0].comision(listaPurchase.precioCompra), nil
	case metodoPago == "cash":
		return listaPurchase.lista[1].comision(listaPurchase.precioCompra), nil

	case metodoPago == "transfer":
		return listaPurchase.lista[2].comision(listaPurchase.precioCompra), nil
	}
	return 0, errors.New("Ingrese un método de pago válido")
}

func main() {

	transaccion1, err := Process(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transaccion1)
	transaccion2, err := Process(t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transaccion2)
	transaccion3, err := Process(e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transaccion3)
	transaccion4, err := Process("mercadopago")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transaccion4)

}
