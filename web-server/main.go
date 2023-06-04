package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.New()
	s.GET("/saludo", func(c *gin.Context) {
		c.String(http.StatusOK, "Hola Bootcampers!")
	})
	s.Run()
}
