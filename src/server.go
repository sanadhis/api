package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	routes := []string{"ping", "pong", "pang"}
	r := gin.Default()

	for _, route := range routes {
		r.GET("/"+route, func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	r.Run()
}
