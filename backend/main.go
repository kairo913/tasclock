package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock/internal/utility"
)

func main() {
	utility.Init()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:5000")
}
