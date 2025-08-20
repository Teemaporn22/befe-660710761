// main.go
package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Product struct
type Products struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Price int    `json:"price"`
}

// In-memory database (
var products = []Products{
	{ID: "1", Name: "Keyboardv1", Color: "black", Price: 1490},
	{ID: "2", Name: "Keyboardv1", Color: "White", Price: 1390},
	{ID: "3", Name: "Keyboardv2", Color: "black", Price: 1850},
	{ID: "4", Name: "Keyboardv2", Color: "White", Price: 1490},
}

func getPro(c *gin.Context) {
	colorQuery := c.Query("color")

	if colorQuery != "" {
		filter := []Products{}
		for _, pro := range products {
			if strings.EqualFold(pro.Color, colorQuery) {
				filter = append(filter, pro)
			}
		}
		c.JSON(http.StatusOK, filter)
		return
	}

	// Return all Produce
	c.JSON(http.StatusOK, products)
}

func main() {
	r := gin.Default()

	r.GET("/pro", func(c *gin.Context) {
		c.JSON(200, gin.H{"Stock": "Good"})
	})

	api := r.Group("/api/pro")
	{
		api.GET("/products", getPro)
	}

	r.Run(":8080")
}
