package gin_v5

import (
	"net/http"
	"strings"
)

// router 提供路由注册与处理方法
type router struct {
	// roots 来存储每种请求方式的 Trie 树根节点
	roots map[string]*node
	// handlers 存储每种请求方式的处理方法 HandlerFunc
	handlers map[string]HandlerFunc
}

// newRouter 创建一个 router 对象
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// parsePattern 解析路由 pattern(只允许第一个 * 匹配)
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// addRoute	注册路由及相应处理方法
// method 	[http 请求方法]
// pattern 	[路由 path]
// handler	[处理方法]
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	// 解析 pattern
	parts := parsePattern(pattern)

	// roots 中插入各种请求方式的 Trie 树根节点
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	// 在具体某个方法的 Trie 树根节点下插入路由全链路节点
	r.roots[method].insert(pattern, parts, 0)
	// handlers 中插入路由对应处理方法HandlerFunc
	r.handlers[key] = handler
}

// getRoute 获取有效路由尾节点以及请求 params 数据
// method	[请求方法]
// path		[路由全路径]
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	// 解析 pattern
	searchParts := parsePattern(path)
	params := make(map[string]string)

	// 查找是否存在当前请求方法 tire 树根节点
	root, ok := r.roots[method]
	// 若不存在直接返回
	if !ok {
		return nil, nil
	}

	// 在当前请求方法 tire 树根节点下搜索有效路由尾节点
	n := root.search(searchParts, 0)
	// 若命中路由
	if n != nil {
		parts := parsePattern(n.pattern)
		// 遍历注册路由 pattern 中的 parts
		for index, part := range parts {
			// 若当前 part 以 ':' 开头(则说明该part为uri)
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			// 若当前 part 以 '*' 开头则直接将后续 searchParts 整合
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	// 若路由未命中
	return nil, nil
}

// getRoutes 获取当前 method 下全部有效路由尾节点
func (r *router) getRoutes(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes
}

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
