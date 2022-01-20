package pkg

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToListNode(s []int) *ListNode {
	ln := &ListNode{}
	tmp := ln
	for _, i := range s {
		// 带头节点
		tmp.Next = &ListNode{
			Val:  i,
			Next: nil,
		}
		// 带尾节点
		//tmp.Val = i
		//tmp.Next = &ListNode{}
		tmp = tmp.Next
	}
	return ln.Next
}
