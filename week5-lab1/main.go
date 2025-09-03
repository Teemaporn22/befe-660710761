package main

import (
	"github.com/gin-gonic/gin"
)

// struct
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		user := []User{{ID: "1", Name: "Teemaporn"}}
		c.JSON(200, user)
	})

	r.Run(":8080")

}
