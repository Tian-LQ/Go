package gin_v5

import (
	"log"
	"net/http"
	"strings"
)

// HandlerFunc 定义请求处理方法
type HandlerFunc func(*Context)

type (
	// RouterGroup 路由分组
	RouterGroup struct {
		prefix      string        // 路由分组前缀
		middlewares []HandlerFunc // 中间件
		engine      *Engine       // 共享 Engine
	}

	// Engine 路由顶层分组
	Engine struct {
		*RouterGroup
		router *router
		groups []*RouterGroup
	}
)

// New 创建并初始化 Engine 实例对象
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group 创建 Group 方法
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// Use 为当前分组添加中间件
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// addRoute 注册路由及相应处理方法
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET 注册 GET 请求路由及相应处理方法
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST 注册 POST 请求路由及相应处理方法
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Run 启动 HTTP Server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 实现了 http.Handler 接口中定义的 ServeHTTP 方法
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	engine.router.handle(c)
}
