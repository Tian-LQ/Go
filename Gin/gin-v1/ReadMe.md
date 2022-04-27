## Gin Web 框架 1.0

------

基于 Go 中内置的 net/http 库，封装出 Server 的概念，实现了 http.Handler 接口

#### http.Handler 接口描述：

在描述 http.Handler 接口之前，需要先看一下 net/http 中提供的 ListenAndServe 方法

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```

其中调用 ListenAndServe 方法后会启动一个 http server，并且将 http 请求交由 handler 处理，这里我们再来看 http.Handler 接口，其中定义了一个方法 ServeHTTP，其中入参即为 http 请求处理方法的基本入参，也就是说我们只要给 ListenAndServe 方法提供一个实现了 http.Handler 接口的实例，那么该 server 在每次收到请求后，都会通过该实例的 ServeHTTP 去处理该请求。其中 ResponseWriter 用于向请求端发送响应报文，*Request 中保存了本次请求相关的元数据信息。

因此 1.0 版本初步实现了 http.Handler 接口，并且封装了注册路由，启动服务的相关方法，并在 ServeHTTP 中通过路由映射执行相应处理方法。

##### 路由 router：

- 通过 map[string]HandlerFunc 这样的结构来存储路由 pattern 与相应处理方法的关系
- 注册路由时将请求方法 method 以及 路由 pattern 作为 key，将处理方法作为 value 插入到路由映射 map 中
- ServeHTTP 处理请求时，根据 request 中信息生成 key，在路由 map 中查找对应 handleFunc 进行处理响应

------

#### 实现功能

- [x] 静态路由注册及处理
- [x] 服务启动