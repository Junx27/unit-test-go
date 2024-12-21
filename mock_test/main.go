package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ItemService interface {
	GetItemByID(id int) (Item, error)
}

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ItemServiceImpl struct{}

func (s *ItemServiceImpl) GetItemByID(id int) (Item, error) {

	if id == 1 {
		return Item{ID: 1, Name: "Item One"}, nil
	}
	return Item{}, http.ErrMissingFile
}

type ItemHandler struct {
	Service ItemService
}

func (h *ItemHandler) GetItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	item, err := h.Service.GetItemByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func SetupRouter(service ItemService) *gin.Engine {
	r := gin.Default()
	handler := &ItemHandler{Service: service}
	r.GET("/items/:id", handler.GetItem)
	return r
}

func main() {
	service := &ItemServiceImpl{}
	r := SetupRouter(service)
	r.Run(":8080")
}
