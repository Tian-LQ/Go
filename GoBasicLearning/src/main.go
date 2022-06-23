package main

import (
	"fmt"
	"math"
)

//func main() {
//	size := 0
//	slice := make([]int, 0)
//	ability := 0
//
//	fmt.Scan(&size)
//	for i := 0; i < size; i++ {
//		val := 0
//		fmt.Scan(&val)
//		slice = append(slice, val)
//	}
//	fmt.Scan(&ability)
//
//	doubleGroupCount := 0
//	singleGroupCount := 0
//	sort.Ints(slice)
//	j := len(slice) - 1
//	for ; j > 0; j-- {
//		if slice[j] < ability {
//			break
//		}
//	}
//	singleGroupCount = len(slice) - 1 - j
//	first := 0
//	second := j
//	for first < second {
//		if slice[first]+slice[second] >= ability {
//			first++
//			second--
//			doubleGroupCount++
//		} else {
//			first++
//		}
//	}
//	fmt.Println(doubleGroupCount + singleGroupCount)
//}

func ScanLine() []int {
	var c int
	var err error
	var b []int
	for err == nil {
		_, err = fmt.Scanf("%d", &c)
		b = append(b, c)
	}
	return b[:len(b)-1]
}

func question3(slice []int) []int {
	n := int(math.Log2(float64(len(slice))))
	index := int(math.Pow(2, float64(n))) - 1
	minSubNodeIndex := len(slice) - 1
	minSubNodeVal := slice[len(slice)-1]
	for i := index; i < len(slice); i++ {
		if slice[i] < minSubNodeVal && slice[i] != -1 {
			minSubNodeVal = slice[i]
			minSubNodeIndex = i
		}
	}
	ret := make([]int, 0)
	for {
		ret = append(ret, slice[minSubNodeIndex])
		if minSubNodeIndex == 0 {
			break
		}
		minSubNodeIndex = (minSubNodeIndex - 1) / 2
	}
	reverseSlice(ret)
	return ret
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

//func main() {
//	s := ScanLine()
//	ret := question3(s)
//	for i := 0; i < len(ret); i++ {
//		fmt.Printf("%d ", ret[i])
//	}
//	fmt.Println()
//}

/*
给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0开头。

例如：
输入： l1 = [2, 4, 3], l2 = [5, 6, 4]
输出：[7, 0, 8]

public class ListNode {
        int val;
        ListNode next;

        ListNode(int x) {
            val = x;
        }
}

public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
}


*/

type ListNode struct {
	val  int
	next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	headNode := &ListNode{}
	preNode := headNode
	flag := 0

	for l1 != nil || l2 != nil {
		var val1 int
		var val2 int
		if l1 != nil {
			val1 = l1.val
			l1 = l1.next
		} else {
			val1 = 0
		}
		if l2 != nil {
			val2 = l2.val
			l2 = l2.next
		} else {
			val2 = 0
		}
		cur := (val1 + val2 + flag) % 10
		flag = (val1 + val2 + flag) / 10

		curNode := &ListNode{
			val:  cur,
			next: nil,
		}
		preNode.next = curNode
		preNode = curNode
	}
	if flag == 1 {
		curNode := &ListNode{
			val:  1,
			next: nil,
		}
		preNode.next = curNode
	}
	return headNode.next
}

func main() {
	l1 := &ListNode{
		val: 9,
		next: &ListNode{
			val:  9,
			next: nil,
		},
	}
	l2 := &ListNode{
		val: 1,
		next: &ListNode{
			val: 1,
			next: &ListNode{
				val:  9,
				next: nil,
			},
		},
	}
	l3 := addTwoNumbers(l1, l2)
	ListNodePrint(l3)
}

func ListNodePrint(l1 *ListNode) {
	for l1 != nil {
		fmt.Printf("%d ", l1.val)
		l1 = l1.next
	}
	fmt.Println()
}
