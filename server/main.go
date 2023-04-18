package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		fmt.Println("roottttt")
		c.JSON(http.StatusOK, gin.H{
			"message": "roott",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("hellooo")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run("0.0.0.0:5000")
}
