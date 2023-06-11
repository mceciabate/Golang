package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
}

var Users = []User{
	{1, "Juan Pablo"},
	{2, "José Perez"},
	{3, "Jaime Pogo"},
	{4, "Jenaro Pinto"},
}

func main() {
	r := gin.Default()

	r.PUT("/users/:id", func(c *gin.Context) {
		var user User
		c.BindJSON(&user)

		idReceived := c.Param("id")
		idReceivedInt, err := strconv.Atoi(idReceived)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for k, u := range Users {
			if u.ID == idReceivedInt {
				user.ID = idReceivedInt
				Users[k] = user
				c.JSON(http.StatusOK, user)
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user with id %d not found", idReceivedInt)})
	})

	r.Run(":8080")
}
