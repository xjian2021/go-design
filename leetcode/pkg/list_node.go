package pkg

import (
	"fmt"
)

//ListNode 单向链表
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

//EchoListNode 打印单向链表
func (t *ListNode) EchoListNode() {
	for i := t; i != nil; i = i.Next {
		fmt.Printf("node[%p]:%+[1]v\n", i)
	}
}

//TwoWayListNode 双向链表
type TwoWayListNode struct {
	Val  int
	Pre  *TwoWayListNode
	Next *TwoWayListNode
}

func SliceToTwoWayListNode(s []int) *TwoWayListNode {
	ln := &TwoWayListNode{}
	tmp := ln
	for _, i := range s {
		// 带头节点
		tmp.Next = &TwoWayListNode{
			Val:  i,
			Pre:  tmp,
			Next: nil,
		}
		tmp = tmp.Next
	}
	ln.Next.Pre = nil
	return ln.Next
}

//EchoTwoWayListNode 打印双向链表
func (t *TwoWayListNode) EchoTwoWayListNode() {
	for i := t; i != nil; i = i.Next {
		fmt.Printf("node[%p]:%+[1]v\n", i)
	}
}
