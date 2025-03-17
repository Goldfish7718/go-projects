package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("Request logged")
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(LoggerMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello gin",
		})
	})

	r.GET("/param/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": name,
		})
	})

	r.POST("/json", func(c *gin.Context) {
		var data struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid JSON",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Recieved",
			"name":    data.Name,
			"email":   data.Email,
		})
	})

	r.Run()
}
