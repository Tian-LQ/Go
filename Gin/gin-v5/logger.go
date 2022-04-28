package gin_v5

import (
	"log"
	"time"
)

// Logger 中间件 Logger 打印请求相关信息
func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
