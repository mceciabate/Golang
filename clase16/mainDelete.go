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

	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, Users)
		return
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		idReceived := c.Param("id")
		idReceivedInt, err := strconv.Atoi(idReceived)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for k, u := range Users {
			if u.ID == idReceivedInt {
				Users = append(Users[:k], Users[k+1:]...)
				c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("user %d deleted", idReceivedInt)})
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user %d not found", idReceivedInt)})

	})

	r.Run(":8080")
}
