package main

import (
	"Gin/gin-v4"
	"net/http"
)

func main() {
	r := gin_v4.New()
	r.GET("/index", func(c *gin_v4.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin_v4.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gin</h1>")
		})

		v1.GET("/hello", func(c *gin_v4.Context) {
			// expect /hello?name=tianlq
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gin_v4.Context) {
			// expect /hello/tianlq
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gin_v4.Context) {
			c.JSON(http.StatusOK, gin_v4.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
