## Gin Web 框架 4.0

------

`Gin Web 4.0` 版本在 `3.0` 版本基础上的增加了路由分组的部分

#### 路由分组

分组控制`(Group Control)`是 Web 框架应提供的基础功能之一。所谓分组，是指路由的分组。之所以对路由分组，是因为我们在业务当中非常可能会对不同分组路由做不同的处理

- 以 `/admin` 开头的路由需要鉴权
- 以 `/api/v1/business` 中 `/api/v1` 开头的路由需要使用 `v1` 版本的 `API` 方法调用

在分组上我们不仅可以使用中间件`(middleware)`，也可以实现代码中的 `AOP`，其中分组还需要支持嵌套，也就是说对于分组 `/group`，`/group/A` 和 `/group/B` 是该分组下的子分组，那么作用于 `/group` 分组上的中间件也需要作用在子分组上。

下面给出路由分组的定义：

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
```

- 我们抽象出了 `RouterGroup` 的概念，用于描述一组路由，其中包括了分组前缀，当前分组内使用的中间件，并且由于我们要根据分组对 `handleFunc` 做不同处理，因此还需要拥有 `Router` 的能力，因此需要共享 `Engine` 对象。

- 在原先 `Engine` 的基础之上，将其抽象为路由的顶层分组，并且保存了所有路由分组列表用于 `handle` 时根据具体路由来将各层路由分组上的 `middleware` 附着在此次请求中。

在添加了 `RouterGroup` 之后，之前所属于 `Engine` 的路由注册方法便交由 `RouterGroup` 去处理，而在Engine 中，以抽象顶层路由分组的形式，嵌套着 `RouterGroup` ，以内嵌的方式同样拥有了路由注册的能力

[这一个版本仅仅只是添加了路由分组，为了后续拓展 middleware 做出入口]()