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
	//listNode = reverseList(listNode)
	listNode = reverseList2(listNode)
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

//递归法
// 1->2->3->nil
// 3->2->1->nil
//f(node):逆转node为头节点的链表
//f(node.next):逆转node.next为头节点的链表
//f(node) = f(node.next) + (node.next.next=node + node.next=nil):逆转当前节点和把下下节点指向当前节点(逆转链表操作)
// 递归写法的技巧在于把每一层与下一层的关系找出来，然后只需要处理好当前层需要做的操作以及下一层的递归调用即可
// 还需要把终止条件找出来!!!
func reverseList2(head *pkg.ListNode) *pkg.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newNode := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil // 这里之所以没事，是因为在把2的next改为nil前已经把1的next改为2了，建立了联系之后，往上一层就是3的next是2，所以不会断点
	return newNode
}
