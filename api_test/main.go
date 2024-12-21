package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItem(c *gin.Context) {
	id := c.Param("id")

	if id == "1" {
		c.JSON(http.StatusOK, gin.H{"id": 1, "name": "Item One"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/items/:id", GetItem)
	return r
}

func main() {
	r := SetupRouter()
	r.Run(":8080")
}
