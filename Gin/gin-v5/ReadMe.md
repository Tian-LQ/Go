## Gin Web 框架 5.0

------

`Gin Web 5.0` 版本在 `4.0` 版本基础上的增加 `middleware` 的部分

#### 中间件 `middleware`

中间件可以理解为与业务无关的技术组件，作为 `Web` 框架需要存在一个可扩展的入口，让用户能够根据自身需求，将各种中间件嵌入到框架之中，仿佛中间件的功能是框架原生支持的一样。

#### `middleware` 设计

在本项目中，关于中间件的设计，采用与路由映射的 `Handler` 一致，输入参数为 `context` 请求上下文对象。而插入到框架之中，是通过向 `context` 引入中间件列表，并且提供 `Next` 方法来完整的处理当前请求。

**`Context` 变动：**

```go
// Context HTTP 请求上下文
type Context struct {
	// 原生的上下文对象
	Writer http.ResponseWriter
	Req    *http.Request
	// 请求相关信息
	Path   string
	Method string
	Params map[string]string
	// 响应相关信息
	StatusCode int
	// 中间件
	handlers []HandlerFunc
	index    int
}

// newContext 创建一个HTTP请求上下文对象
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

// Next 处理中间件
func (c *Context) Next() {
	c.index++
	n := len(c.handlers)
	for ; c.index < n; c.index++ {
		c.handlers[c.index](c)
	}
}
```

**`middleware` 形式：**

```go
// 中间件
func Middleware() HandlerFunc {
    return func(c *Context) {
        // do something...
        c.Next()
        // do something...
    }
}
```

做到这样并不够，因为还没有将真正的业务处理与中间件处理结合在一起，其本质是通过将中间件以及路由 `handleFunc` 插入 `context` 的 `handlers` 列表中：

```go
// handle 根据请求上下文信息路由到相应处理方法
func (r *router) handle(c *Context) {
	// 获取有效路由尾节点以及请求 params 数据
	n, params := r.getRoute(c.Method, c.Path)
	// 若路由命中则保存 params 数据至 context 中并调用相应路由处理方法
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		// 将真正的路由处理逻辑加入 context 的 handler 列表中
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(context *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
```

也就是说真正的 `handleFunc` 会被放置到 `handlers` 列表的末尾，在执行 `context` 的 `Next()` 方法时会将各个中间件的处理，叠加包裹在业务处理 `handleFunc` 的外围，最终执行我们的业务处理。

我们之前说了中间件应该是作用于路由分组层面的，这样才能便于我们做请求处理的多样化，因此创建路由分组后，应该能够为当前分组设置中间件。

```go
// Use 为当前分组添加中间件
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}
```

那么我们应该如何为真实的一条 `http request` 请求的 `context` 上绑定所属于它的中间件呢？这里我们通过保存所有的路由分组 `RouterGroup` 信息在 `Engine` 之上，而 `RouterGroup` 本身之上包含其专属的 `middleware` 列表，在处理请求的最开始，简单的通过 `request` 的 `URL` 路径，是否包含路由分组 `prefix` 前缀，来为 `context` 绑定。

```go
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
```

[到这里我们的 `5.0` 版本便完成了，可以自由的在不同路由分组层添加各自的 `middleware`，让我们的框架更加的灵活起来]()

