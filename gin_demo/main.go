package main

import "github.com/gin-gonic/gin"

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello golang",
	})
}

func main() {
	// 返回默认路由引擎
	r := gin.Default()
	// 注册路由
	r.GET("/hello", sayHello)
	// 启动服务
	r.Run(":9090")
}
