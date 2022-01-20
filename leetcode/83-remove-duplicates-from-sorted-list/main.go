package main

import (
	"fmt"

	. "github.com/xjian2021/go-design/leetcode/pkg"
)

func main() {
	node := SliceToListNode([]int{1})
	n := deleteDuplicates(node.Next)
	for i := n; i != nil; i = i.Next {
		fmt.Printf("i:%+v\n", i)
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tmpVal := head.Val
	tmpNode := head
	for i := head.Next; i != nil; i = i.Next {
		if tmpVal == i.Val {
			tmpNode.Next = i.Next
			continue
		}
		tmpVal = i.Val
		tmpNode = i
	}
	return head
}
