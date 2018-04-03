package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"text": "Event Detected",
		})
	})
	r.Run()
}
