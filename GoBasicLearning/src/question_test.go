package main

import (
	"math"
	"sort"
	"testing"
)

func function(s []int) int {
	i := 0
	for ; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			break
		}
	}
	if i == len(s)-1 {
		return 0
	}
	return i + 1
}

func TestFunctionName(t *testing.T) {
	s := []int{1, 2, 3}
	t.Log(function(s))
}

func function2(s []int) bool {
	if s == nil || len(s) <= 2 {
		return true
	}
	m := make(map[int]bool)
	min, max := s[0], s[0]
	for _, val := range s {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
		m[val] = true
	}
	x := (max - min) % (len(s) - 1)
	if x != 0 {
		return false
	}
	n := (max - min) / (len(s) - 1)
	for i := 0; i < len(s); i++ {
		_, ok := m[min+i*n]
		if !ok {
			return false
		}
	}
	return true
}

func TestFunction2Name(t *testing.T) {
	s := []int{1}
	t.Log(function2(s))
}

func function3(str string) bool {
	length := len(str)
	if length < 2 {
		return false
	}
	result := false
	for L := 1; L < length/2; L++ {
		if length%L != 0 {
			continue
		}
		flag := true
		repeatStr := str[:L]
		for i := 0; i < length-L; i += L {
			if str[i:i+L] != repeatStr {
				flag = false
				break
			}
		}
		if flag {
			result = true
			break
		}
	}
	return result
}

func TestFunction3(t *testing.T) {
	str := "ABCAB"
	t.Log(function3(str))
}

/*
假设您能预知某只股票未来n秒内价格（每秒1个价格），
设计一个算法来找到最大的利润，并写出实现代码。限制如下：
1) 您最多只能交易2次（买一次，卖一次，且先买后卖）
2) 同时最多能持有1股
3) 空间复杂度O(1)
4) 时间复杂度O(N)
*/

func maxProfit(prices []int) []int {
	ret := 0
	curMin := math.MaxInt
	curMinIndex := 0
	buyIndex := 0
	saleIndex := 0
	for i := 0; i < len(prices); i++ {
		if prices[i]-curMin > ret {
			ret = prices[i] - curMin
			buyIndex = curMinIndex
			saleIndex = i
		}

		if prices[i] < curMin {
			curMin = prices[i]
			curMinIndex = i
		}
	}
	return []int{buyIndex, saleIndex}
}

func maxLenSubStr(str string) int {
	ret := 0
	i := 0
	for i < len(str) {
		if ret > len(str)-i {
			break
		}
		m := make(map[string]int)
		for j := i; j < len(str); j++ {
			val, ok := m[string(str[j])]
			if !ok {
				m[string(str[j])] = j
			} else {
				if len(m) > ret {
					ret = len(m)
					i = val + 1
				}
				break
			}
			if j == len(str)-1 {
				if len(m) > ret {
					ret = len(m)
					i = len(str)
				}
				break
			}
		}
	}
	return ret
}

func TestMaxLenSubStr(t *testing.T) {
	t.Log(maxLenSubStr("abcdaf"))
}

func question1(slice []int, ability int) int {
	doubleGroupCount := 0
	singleGroupCount := 0
	sort.Ints(slice)
	j := len(slice) - 1
	for ; j > 0; j-- {
		if slice[j] < ability {
			break
		}
	}
	singleGroupCount = len(slice) - 1 - j
	first := 0
	second := j
	for first < second {
		if slice[first]+slice[second] >= ability {
			first++
			second--
			doubleGroupCount++
		} else {
			first++
		}
	}

	return doubleGroupCount + singleGroupCount
}

//func question3(slice []int) []int {
//	n := int(math.Log2(float64(len(slice))))
//	index := int(math.Pow(2, float64(n))) - 1
//	minSubNodeIndex := len(slice) - 1
//	minSubNodeVal := slice[len(slice)-1]
//	for i := index; i < len(slice); i++ {
//		if slice[i] < minSubNodeVal && slice[i] != -1 {
//			minSubNodeVal = slice[i]
//			minSubNodeIndex = i
//		}
//	}
//	ret := make([]int, 0)
//	for {
//		ret = append(ret, slice[minSubNodeIndex])
//		if minSubNodeIndex == 0 {
//			break
//		}
//		minSubNodeIndex = (minSubNodeIndex - 1) / 2
//	}
//	reverseSlice(ret)
//	return ret
//}
//
//func reverseSlice(s []int) {
//	for i, j := 0, len(s)-1; i < j; {
//		s[i], s[j] = s[j], s[i]
//		i++
//		j--
//	}
//}

//func TestQuestion1(t *testing.T) {
//	s := []int{5, 9, 8, -1, -1, 7, -1, -1, -1, -1, -1, 6}
//	t.Log(question3(s))
//}

func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for ; i <= p2 && nums[i] == 2; p2-- {
			nums[i], nums[p2] = nums[p2], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}

type Node struct {
	key  int
	val  int
	next *Node
	pre  *Node
}

type LRUCache struct {
	m    map[int]*Node
	head *Node
	tail *Node
	cap  int
}

func Constructor(capacity int) LRUCache {
	lryCache := LRUCache{
		m:    map[int]*Node{},
		head: &Node{},
		tail: &Node{},
		cap:  capacity,
	}
	lryCache.head.next = lryCache.tail
	lryCache.tail.pre = lryCache.head
	return lryCache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.m[key]
	if ok {
		// 链表中取出该节点
		node.pre.next = node.next
		node.next.pre = node.pre
		// 将当前节点放置链表首部
		node.next = this.head.next
		this.head.next.pre = node
		node.pre = this.head
		this.head.next = node
		return node.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.m[key]
	if ok {
		node.val = value
		// 链表中取出该节点
		node.pre.next = node.next
		node.next.pre = node.pre
		// 将该节点放置链表首部
		node.next = this.head.next
		this.head.next.pre = node
		node.pre = this.head
		this.head.next = node
	} else {
		// 创建新节点
		newNode := &Node{key: key, val: value}
		// 将新节点放置链表首部
		newNode.next = this.head.next
		this.head.next.pre = newNode
		newNode.pre = this.head
		this.head.next = newNode
		this.m[key] = newNode
		if len(this.m) > this.cap {
			lastNode := this.tail.pre
			lastNode.pre.next = this.tail
			this.tail.pre = lastNode.pre
			delete(this.m, lastNode.key)
		}
	}
}

func TestLRU(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(2, 1)
	lruCache.Put(1, 1)
	lruCache.Put(2, 3)
	lruCache.Put(4, 1)
	t.Log(lruCache.Get(1))
	t.Log(lruCache.Get(2))
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	ret := 0
	tem := 1
	pre := nums[0]
	for _, num := range nums {
		if num == pre+1 {
			tem++
			pre = num
		} else if num == pre {
			continue
		} else {
			if tem > ret {
				ret = tem
			}
			pre = num
			tem = 1
		}
	}
	if tem > ret {
		ret = tem
	}
	return ret
}

func longestConsecutivePlus(nums []int) int {
	ret := 1
	m := map[int]bool{}
	for _, val := range nums {
		m[val] = false
	}
	for _, val := range nums {
		if m[val] == true {
			continue
		}
		cur := leftSearch(m, val) + rightSearch(m, val) + 1
		if cur > ret {
			ret = cur
		}
	}
	return ret
}

func leftSearch(m map[int]bool, key int) int {
	ret := 0
	for i := key - 1; ; i-- {
		_, ok := m[i]
		if !ok {
			break
		}
		m[i] = true
		ret++
	}
	return ret
}

func rightSearch(m map[int]bool, key int) int {
	ret := 0
	for i := key + 1; ; i++ {
		_, ok := m[i]
		if !ok {
			break
		}
		m[i] = true
		ret++
	}
	return ret
}

func TestLongestConsecutive(t *testing.T) {
	nums := []int{0, 1, 2}
	t.Log(longestConsecutive(nums))
}
