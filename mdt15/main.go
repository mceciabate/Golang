package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

/*Lista de productos
Vamos a crear una aplicación web con el framework Gin que nos permita consultar datos sobre los usuarios y sus posts en nuestro sitio.
Para esto, necesitaremos descargar, desde este link, los dos archivos JSON que contienen dicha información.
Crear una estructura Usuario y otra Post que representen los datos que encontrarán dentro de cada archivo JSON.
Leer los archivos users.json y posts.json:
Crear un método GET que permite imprimir la lista completa de usuarios en formato JSON.
Crear un método GET que permite recuperar los datos de un único usuario, en formato JSON, recibiendo el ID del mismo por parámetro en la URL.
Crear un método GET que permite recuperar todos los posts realizados por un usuario en particular. Imprimirlo en formato JSON.
Crear una ruta /search que permite buscar por parámetro los usuarios cuyo nombre contenga lo ingresado en el parámetro name (total o parcialmente).*/

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type User struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}
type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Business    string `json:"bs"`
}

var post []Post
var users []User

func main() {
	loadPost("./data/posts.json")
	loadUsers("./data/users.json")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })
	usersGroup := r.Group("/users")
	{
		usersGroup.GET("/getAll", GetAll)
		usersGroup.GET("/:id", SearchById)
		usersGroup.GET("/:id/posts", SearchPostById)
		usersGroup.GET("/search/:text", SearchUser)

	}
	r.Run(":8080")

}

func SearchUser(c *gin.Context) {
	frase := c.Param("text")
	var search []User
	for _, v := range users {
		if strings.Contains(v.Name, frase) {
			search = append(search, v)
		}
	}
	c.JSON(http.StatusOK, search)
}

func SearchPostById(c *gin.Context) {
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Id",
		})
		return
	}
	var postUser []Post
	for _, v := range post {
		if v.UserId == Id {
			postUser = append(postUser, v)
		}

	}
	c.JSON(http.StatusOK, postUser)

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
	var encontrado User
	for _, v := range users {
		if v.Id == Id {
			encontrado = v
		}
	}

	c.JSON(http.StatusOK, encontrado)

}

func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func loadPost(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(file, &post); err != nil {
		panic(err)
	}
}

func loadUsers(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(file, &users); err != nil {
		panic(err)
	}
}
