package gin_v2

import (
	"log"
	"net/http"
)

// router 提供路由注册与处理方法
type router struct {
	handlers map[string]HandlerFunc
}

// newRouter 创建一个 router 对象
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// addRoute	注册路由及相应处理方法
// method 	[http 请求方法]
// pattern 	[路由 path]
// handler	[处理方法]
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// handle 根据请求上下文信息路由到相应处理方法
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
