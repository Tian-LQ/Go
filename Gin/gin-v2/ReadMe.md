## Gin Web 框架 2.0

------

`Gin Web 2.0` 版本在 `1.0` 版本的基础上引入了请求上下文 `Context` 的概念，并且将之前的路由 `router` 抽离出来，便于后续进行拓展。 

#### `Context` 设计

- **原因**

首先先来看看没有引入 `context` 之前，此时我们每次 `http` 请求的上下文信息是通过 `(ResponseWriter, *Request)` 来获取到的，我们通过 `*Request` 获取请求数据，结合自身逻辑处理，再通过 `ResponseWriter` 响应请求。在这种模式下，我们在处理请求的 `handleFunc` 中会经常性的使用 `*Request` 当中的方法来获取一些基本的请求数据信息，同时也会重复性性的通过 `ResponseWriter` 来返回响应信息，这样一来对于我们的使用者(也就是注册路由用户)来说，是会需要频繁的写大量重复性代码的，[因此我们将每次的 http 请求上下文信息用一个 `context` 的概念来封装起来，并且提供一些常用的方法，便于我们的使用者使用。]()

- **功能**
  - [x] `context` 中除了包含原生的 `http` 包的上下文信息 `writer` 和 `request` 之外，还包括了请求及相应相关信息
  - [x] 提供了从请求报文 `form` 表单中获取 `[key, value]` 数据的方法
  - [x] 提供了从请求 `URL query parm` 中获取 `[key, value]` 数据的方法
  - [x] 提供了设置响应报文状态码的方法
  - [x] 提供了向响应报文当中的响应首部添加 `[key,value]` 值对的方法
  - [x] 提供了将 `String` 类型数据写入响应报文 `body` 中的方法
  - [x] 提供了将 `Json` 类型数据写入响应报文 `body` 中的方法
  - [x] 提供了将 `Html` 类型数据写入响应报文 `body` 中的方法

#### `Router` 抽离

之前我们是 `router` 是通过一个 `map` 来存储我们的路由与处理方法映射关系的，这次我们将 `router` 的概念单独抽离并封装起来，同时我们处理路由的 `handle` 方法参数也变成了上面的 `context`，但是实现方式还是基于原先的逻辑，这样一来是为了便于后续对 `router` 部分进行拓展更新。

------

**这里其实 `router` 的抽离应该考虑通过 `interface` 的方式去抽离出来，而不是采用一个具体的 `struct`**