package main

import (
	"Gin/gin-v3"
	"net/http"
)

func main() {
	var filterBuilder gin_v3.FilterBuilder = gin_v3.MetricFilterBuilder
	r := gin_v3.New(filterBuilder)
	r.GET("/", func(c *gin_v3.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gin_v3.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gin_v3.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gin_v3.Context) {
		c.JSON(http.StatusOK, gin_v3.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
