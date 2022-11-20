package main

import (
	"fmt"

	"github.com/xjian2021/go-design/leetcode/pkg"
)

//剑指 Offer 24. 反转链表
//定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

func main() {
	listNode := pkg.SliceToListNode([]int{1, 2, 3, 4, 5})
	fmt.Println("输入：")
	listNode.EchoListNode()
	fmt.Println("输出：")
	listNode = reverseList(listNode)
	listNode.EchoListNode()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *pkg.ListNode) *pkg.ListNode {
	tmp := head
	var out *pkg.ListNode
	for tmp != nil {
		next := tmp.Next
		tmp.Next = out // 这里才是真正改变链表顺序的操作，把当前节点的下一个节点
		out = tmp
		tmp = next
	}
	return out
}
