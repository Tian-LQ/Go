package main

import (
	"Gin/gin-v6"
	"net/http"
)

func main() {
	r := gin_v6.Default()
	r.GET("/", func(c *gin_v6.Context) {
		c.String(http.StatusOK, "Hello tianlq\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gin_v6.Context) {
		names := []string{"tianlq"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
