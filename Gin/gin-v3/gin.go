package gin_v3

import (
	"log"
	"net/http"
)

// HandlerFunc 定义请求处理方法
type HandlerFunc func(*Context)

// Engine http.Handler 接口的实例化对象
type Engine struct {
	router *router
	root   Filter
}

// New 创建并初始化 Engine 实例对象
func New(builders ...FilterBuilder) *Engine {
	router := newRouter()
	// 根 filter
	var root = func(c *Context) {
		router.handle(c)
	}
	// 包装根 filter(从后往前)
	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}
	return &Engine{router: router, root: root}
}

// addRoute 注册路由及相应处理方法
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router.addRoute(method, pattern, handler)
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
	c := newContext(w, req)
	engine.root(c)
}
