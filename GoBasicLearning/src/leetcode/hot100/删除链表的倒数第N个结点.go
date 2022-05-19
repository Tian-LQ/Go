package hot100

/*
 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点
*/

type listNode struct {
	Val  int
	Next *listNode
}

func removeNthFromEnd(head *listNode, n int) *listNode {
	emptyHead := &listNode{
		Next: head,
	}
	left, right := emptyHead, emptyHead
	for i := 0; i < n; i++ {
		if right == nil {
			return nil
		}
		right = right.Next
	}
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}
	delNode := left.Next
	left.Next = delNode.Next
	return emptyHead.Next
}
