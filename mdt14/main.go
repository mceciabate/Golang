package main

/*Consigna
Un supermercado necesita un sistema para gestionar los productos frescos que tienen publicados en su página web.
Para poder hacer esto, necesitan un servidor que ejecute una API que les permita manipular los productos cargados
desde distintos clientes. Los campos que conforman un producto son:
Con la misma API que veníamos trabajando en clase, vamos a resolver los siguientes ejercicios.

Crear una ruta /productparams que tome todos los datos de la estructura de un producto por parámetro y
lo devuelva en forma de JSON. Podemos seguir el siguiente ejemplo:

Insertar este último producto a la lista de productos y verificar si lo podemos tomar con la ruta /products/:id
(más adelante veremos otro método para insertar datos en nuestras listas o bases de datos).

Se necesita un endpoint que devuelva una lista de productos que estén entre ciertas cantidades de stock. Por ejemplo:
los productos que tengan entre 300 y 400 unidades. La ruta se llamará /searchbyquantity y debemos pasar los límites de las cantidades por parámetro.

Se necesita un endpoint que brinde el detalle de una compra. Por parámetro se deberá pasar el code_value del producto y la cantidad de unidades a comprar.
El detalle de la compra deberá ser: nombre del producto, cantidad y precio total. La ruta se deberá llamar /buy.

*/

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var products []Product

func main() {
	loadProducts("products.json")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })
	productsGroup := r.Group("/products")
	{
		//productsGroup.GET("/:id", getByID)            // path param
		productsGroup.GET("/search", Search) // query param: products/search?priceGt=100&isPublished=true
		productsGroup.GET("/productparams", GetByParam)
		productsGroup.GET("/:id", SearchById)
		productsGroup.GET("/searchbyquantity", GetMinMax)
		productsGroup.GET("/buy", GetBuy)

	}

	r.Run(":8080")
}

func Search(c *gin.Context) {
	query := c.Query("priceGt")
	priceGt, err := strconv.ParseFloat(query, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Price",
		})
		return
	}

	list := []Product{}
	for _, product := range products {
		if product.Price > priceGt {
			list = append(list, product)
		}
	}

	c.JSON(http.StatusOK, list)
}

func GetByParam(c *gin.Context) {
	Id := c.Query("id")
	id, err := strconv.Atoi(Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Id",
		})
		return
	}
	Name := c.Query("name")
	Quantity := c.Query("quantity")
	quantity, err := strconv.Atoi(Quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Quantity",
		})
		return
	}
	CodeValue := c.Query("code_value")
	IsPublished := c.Query("is_published")
	p, err3 := strconv.ParseBool(IsPublished)
	if err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Bool",
		})
		return
	}
	Expiration := c.Query("expiration")
	Price := c.Query("price")
	price, err2 := strconv.ParseFloat(Price, 64)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Quantity",
		})
		return
	}
	productoObtenido := Product{
		ID:          id,
		Name:        Name,
		Quantity:    quantity,
		CodeValue:   CodeValue,
		IsPublished: p,
		Expiration:  Expiration,
		Price:       price,
	}
	products = append(products, productoObtenido)
	escribir()
	c.JSON(http.StatusOK, productoObtenido)

}

func escribir() {
	listaActualizada, err := json.Marshal(products)
	if err != nil {
		panic(err)
	}
	file, err2 := os.OpenFile("products.json", os.O_RDWR, 0666)
	if err2 != nil {
		panic(err2)
	}
	_, e := file.Write(listaActualizada)
	if e != nil {
		panic(e)
	}
	file.Close()
}

func SearchById(c *gin.Context) {
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Id",
		})
		return
	}
	var productoEncontrado Product
	for _, v := range products {
		if v.ID == Id {
			productoEncontrado = v
		}
	}
	if productoEncontrado == (Product{}) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No se encuentra el producto",
		})
	} else {
		c.JSON(http.StatusOK, productoEncontrado)
	}

}

func GetMinMax(c *gin.Context) {
	min := c.Query("min")
	Min, err := strconv.Atoi(min)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Min",
		})
		return
	}
	max := c.Query("max")
	Max, err := strconv.Atoi(max)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Max",
		})
		return
	}
	var listaMinMax []Product
	for _, v := range products {
		if v.Quantity > Min && v.Quantity < Max {
			listaMinMax = append(listaMinMax, v)
		}
	}
	if len(listaMinMax) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No se encuentran coincidencias",
		})
	} else {
		c.JSON(http.StatusOK, listaMinMax)
	}

}

func GetBuy(c *gin.Context) {
	codeValue := c.Query("code_value")
	cantidad := c.Query("cantidad")
	Cantidad, err := strconv.Atoi(cantidad)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Cantidad",
		})
		return
	}
	var productoEncontrado Product
	for _, v := range products {
		if v.CodeValue == codeValue {
			productoEncontrado = v
		}
	}
	if productoEncontrado == (Product{}) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Producto No encontrado",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Nombre del producto": productoEncontrado.Name,
			"Cantidad":            Cantidad,
			"Precio Total":        float64(Cantidad) * productoEncontrado.Price,
		})

	}

}

func loadProducts(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(file, &products); err != nil {
		panic(err)
	}
}
