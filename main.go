package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":4000") // listen and serve on 0.0.0.0:8080
}
