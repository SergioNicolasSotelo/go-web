package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.GET("/hola", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hola",
		})
	})

	router.Run()
}
