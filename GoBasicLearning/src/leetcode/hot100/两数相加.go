package hot100

/*
 给你两个非空的链表，表示两个非负的整数。
 它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
 请你将两个数相加，并以相同形式返回一个表示和的链表。
 你可以假设除了数字 0 之外，这两个数都不会以 0开头。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	headNode := ListNode{}
	preNode := &headNode
	val1, val2 := 0, 0
	flag := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		} else {
			val1 = 0
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		} else {
			val2 = 0
		}
		curVal := (val1 + val2 + flag) % 10
		flag = (val1 + val2 + flag) / 10
		curNode := ListNode{
			Val:  curVal,
			Next: nil,
		}
		preNode.Next = &curNode
		preNode = &curNode
	}
	if flag == 1 {
		lastNode := ListNode{
			Val:  1,
			Next: nil,
		}
		preNode.Next = &lastNode
	}
	return headNode.Next
}
