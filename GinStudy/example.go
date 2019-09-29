package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		//c.JSON(200, gin.H{"message": "pong"})
		c.SecureJSON(200, names)
	})
	r.Run()
}
