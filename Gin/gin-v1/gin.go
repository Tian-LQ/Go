package gin_v1

import (
	"fmt"
	"net/http"
)

// HandlerFunc 定义请求处理方法
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine http.Handler 接口的实例化对象
type Engine struct {
	router map[string]HandlerFunc
}

// New 创建并初始化 Engine 实例对象
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute 注册路由及相应处理方法
// method 	[http 请求方法]
// pattern 	[路由 path]
// handler	[处理方法]
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET 注册 GET 请求路由及相应处理方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 注册 POST 请求路由及相应处理方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 启动 HTTP Server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 实现了 http.Handler 接口中定义的 ServeHTTP 方法
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
