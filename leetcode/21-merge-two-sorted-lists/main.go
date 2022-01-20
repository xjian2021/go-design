package main

import (
	"fmt"

	. "github.com/xjian2021/go-design/leetcode/pkg"
)

//21. 合并两个有序链表
//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func main() {
	l1 := []int{1, 2, 4}
	l2 := []int{1, 3, 4}
	n1 := SliceToListNode(l1)
	n2 := SliceToListNode(l2)
	n3 := mergeTwoLists(n1, n2)
	for ; n3 != nil; n3 = n3.Next {
		fmt.Println(n3)
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l := &ListNode{}
	lTemp := l
	for {
		if l1.Val > l2.Val {
			lTemp.Next = &ListNode{l2.Val, nil}
			l2 = l2.Next
		} else {
			lTemp.Next = &ListNode{l1.Val, nil}
			l1 = l1.Next
		}
		lTemp = lTemp.Next
		if l1 == nil {
			if l2 != nil {
				lTemp.Next = l2
			}
			break
		}
		if l2 == nil {
			lTemp.Next = l1
			break
		}
	}

	return l.Next
}
