package gin_v6

import (
	"fmt"
	"strings"
)

// node trie 树节点
type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// String node 字符串输出方法
func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

// matchChild 当前 node 中第一个匹配 part 成功的子节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 当前 node 中所有匹配 part 成功的子节点集合，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert 向 trie 树中插入节点
// pattern	[完整路由]
// parts	[完整路由拆分后的 string 数组]
// height	[当前节点位于 trie 树中的层数]
func (n *node) insert(pattern string, parts []string, height int) {
	// 若当前节点层数恰好与 parts 总数相同(即当前节点命中完整路由)
	if len(parts) == height {
		// 为当前节点 node 设置 pattern
		n.pattern = pattern
		return
	}

	// 获取下一层需要匹配 pattern 中的 part
	part := parts[height]
	// 在子节点列表中查找有无匹配 child
	child := n.matchChild(part)
	// 若没有匹配的 child 则新建并追加至 children 子节点列表中
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	// 若存在匹配 child 则交由 child 完成后续插入工作
	child.insert(pattern, parts, height+1)
}

// search 在 tire 树中查找路由
func (n *node) search(parts []string, height int) *node {
	// 若当前节点命中完整路由，或当前节点匹配 part 以 '*' 开头
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		// 若该节点 pattern 为空则返回 nil 表示未命中
		if n.pattern == "" {
			return nil
		}
		// 否则返回当前节点表示命中
		return n
	}

	// 获取下一层需要匹配 pattern 中的 part
	part := parts[height]
	// 在子节点列表中查找匹配 part 的节点集合
	children := n.matchChildren(part)

	// 遍历该集合
	for _, child := range children {
		// 交由匹配的 child 节点继续完成后续查找工作
		result := child.search(parts, height+1)
		// 如果存在匹配结果则返回(也就是说 search 方法会返回第一个匹配的节点)
		if result != nil {
			return result
		}
	}
	return nil
}

// travel 将 trie 树中的有效节点均追加至 list 中
func (n *node) travel(list *[]*node) {
	// 若当前节点存在 pattern (即有效路由尾节点) 则将该节点追加至 list
	if n.pattern != "" {
		*list = append(*list, n)
	}
	// 遍历所有子节点，重复上述操作
	for _, child := range n.children {
		child.travel(list)
	}
}
