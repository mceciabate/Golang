package main

/*Lista de productos
Vamos a crear una aplicación web con el framework Gin que tenga un endpoint /productos que responda con una lista de productos.
Crear una estructura Productos con los valores:
Id
Nombre
Precio
Stock
Codigo
Publicado
FechaDeCreacion
Crear un archivo productos.json con al menos seis productos (deben seguir la misma estructura ya mencionada).
Leer el archivo productos.json.
Imprimir la lista de productos por consola.
Imprimir la lista de productos a través del endpoint en formato JSON. El endpoint deberá ser de método GET.
*/

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id              int     `json:"Id"`
	Nombre          string  `json:"Nombre"`
	Precio          float64 `json:"Precio"`
	Stock           int     `json:"Stock"`
	Codigo          string  `json:"Codigo"`
	Publicado       bool    `json:"Publicado"`
	FechaDeCreacion string  `json:"FechaCreacion"`
}

func ObtenerDatos(categoria string) ([]Producto, error) {
	var lista []Producto
	switch categoria {
	case "electro":
		data, e1 := os.ReadFile("./data/electro.json")
		if e1 != nil {
			log.Fatal(e1)
		}
		if err := json.Unmarshal([]byte(data), &lista); err != nil {
			log.Fatal(err)
		}
		return lista, nil
	case "indumentaria":
		data, e1 := os.ReadFile("./data/indumentaria.json")
		if e1 != nil {
			log.Fatal(e1)
		}
		if err := json.Unmarshal([]byte(data), &lista); err != nil {
			log.Fatal(err)
		}
		return lista, nil
	case "juguetes":
		data, e1 := os.ReadFile("./data/juguetes.json")
		if e1 != nil {
			log.Fatal(e1)
		}
		if err := json.Unmarshal([]byte(data), &lista); err != nil {
			log.Fatal(err)
		}
		return lista, nil
	case "instrumentos":
		data, e1 := os.ReadFile("./data/instrumentos.json")
		if e1 != nil {
			log.Fatal(e1)
		}
		if err := json.Unmarshal([]byte(data), &lista); err != nil {
			log.Fatal(err)
		}
		return lista, nil
	default:
		return nil, errors.New("No se encuentra la categoría especificada")
	}
}

func main() {
	listaElectro, e := ObtenerDatos("electro")
	if e != nil {
		log.Fatal(e)
	}
	listaIndumentaria, e2 := ObtenerDatos("indumentaria")
	if e2 != nil {
		log.Fatal(e2)
	}
	listaJuguetes, e3 := ObtenerDatos("juguetes")
	if e3 != nil {
		log.Fatal(e3)
	}
	listaInstrumentos, e4 := ObtenerDatos("instrumentos")
	if e4 != nil {
		log.Fatal(e4)
	}
	// data, e := os.ReadFile("productos.json")
	// if e != nil {
	// 	log.Fatal(e)
	// }
	// data1, e1 := os.ReadFile("categorias.json")
	// if e1 != nil {
	// 	log.Fatal(e1)
	// }

	// var lista []Producto
	// if err := json.Unmarshal([]byte(data), &lista); err != nil {
	// 	log.Fatal(err)
	// }
	// if err1 := json.Unmarshal([]byte(data1), &categorias); err1 != nil {
	// 	log.Fatal(err1)
	// }
	// fmt.Println(lista)
	router := gin.Default()
	router.GET("/productos", func(c *gin.Context) {
		c.JSON(200, gin.H{"Juguetes": listaJuguetes,
			"Electrodomésticos":      listaElectro,
			"Indumentaria":           listaIndumentaria,
			"Instrumentos Musicales": listaInstrumentos})
	})
	router.Run()

}
