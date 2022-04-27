## Gin Web 框架 3.0

------

`Gin Web 3.0` 版本在 `2.0` 版本基础上的最大改动是增强了 `Router` 的功能，支持动态路由解析

#### `Router` 设计

我们之前的 `Router` 是通过 `map` 来存储路由解析的相关数据，这种方法虽然效率很高，但是只能支持静态路由，也就是无法支持类似 `Restful API` 中类似 `/user/:ID` 获取指定用户信息这样的动态路由(`ID` 标识资源)，因此通过字典树(`trie` 树)这种数据结构来实现动态路由的功能。

由于我们 `http` 请求的 `path` 刚好是由 / 分隔的多段构成，因此可以将每一段作为 `trie` 树当中的节点。

- **`trie node` 设计**

  ```go
  // node trie 树节点
  type node struct {
  	pattern  string  // 待匹配路由，例如 /p/:lang
  	part     string  // 路由中的一部分，例如 :lang
  	children []*node // 子节点，例如 [doc, tutorial, intro]
  	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
  }
  ```

  - [x] 其中 `pattern` 表示完整路由匹配的 `path`，也就是说对于注册 `/user/home/address` 这样的路由的过程，会生成三个 `node` 节点，但是仅仅 `address` 所在的节点是代表一整个完整路由的，因此该节点上会保存完整 `pattern` 的信息，其他节点的 `pattern` 均为 `string` 零值空字符串
  - [x] `part` 存储的就是 `http` 请求 `path` 中具体某一段中的部分 
  - [x] 由于我们需要支持动态路由，原生的 `tire` 树本身并不能做到，因此我们添加了 `isWild` 字段，来表示当前节点是否需要特殊处理(如路由 `/user/:ID` 中 `ID` 所在的节点)

- **`trie` 树为 `Router` 提供的功能**

  - [x] 在 `tire` 树中插入完整路由的 `insert` 方法
  - [x] 在 `tire` 树中寻找具体路由的 `search` 方法
    - 在寻找最终 `node` 的过程中，还需要保存路由解析当中的参数 `params` (以 `map[string]string` 的形式存储)
    - 例：注册路由 `/user/:ID` ，实际请求 `/user/19`，那么保存的 `params` 为 `{"ID": "19"}`
  - [x] 支持参数匹配：`/p/:lang/doc`，可以匹配 `/p/c/doc` 和 `/p/go/doc`
  - [x] 支持通配符 `*`：`/static/*filepath`,可以匹配 `/static/a/b/c`

- **`router` 设计**

  ```go
  // router 提供路由注册与处理方法
  type router struct {
  	// roots 来存储每种请求方式的 Trie 树根节点
  	roots map[string]*node
  	// handlers 存储每种请求方式的处理方法 HandlerFunc
  	handlers map[string]HandlerFunc
  }
  ```

  - [x] 在原先 `router` 的基础之上增加了 `roots` 字典树 `map`，来存储各种请求所对应的路由树根节点，通俗来讲的话，`roots` 中存储的数据是 `["POST", &node{}]` 这样的 `tire` 树根节点
  - [x] `handlers` 依然存储路由及处理方法的映射关系

#### `Context` 调整

由于增加了动态路由的功能，因此会存在解析动态路由过程中产生的 `params` 数据，且这部分数据应归属于 `http` 请求上下文，因此拓展 `context` 当中的数据。

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
}
```

其中 `Params` 的值是在路由解析(`tire` 树查找路由)之后被赋值用于 `handleFunc` 中访问的，相应的 `context` 提供了 `Param` 方法用于获取当前请求 `URL` 当中 `param [key,value]` 数据

[`context` 本身的不断拓展，最终目的都是服务于框架的使用者]()

