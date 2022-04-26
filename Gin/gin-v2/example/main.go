package main

import (
	"Gin/gin-v2"
	"net/http"
)

func main() {
	r := gin_v2.New()
	r.GET("/", func(c *gin_v2.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gin_v2.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gin_v2.Context) {
		c.JSON(http.StatusOK, gin_v2.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
