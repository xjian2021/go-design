package main

import (
	"fmt"

	"github.com/xjian2021/go-design/leetcode/pkg"
)

//剑指 Offer 06. 从尾到头打印链表
//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

func main() {
	listNode := pkg.SliceToListNode([]int{1, 3, 2})
	fmt.Println(reversePrint(listNode))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *pkg.ListNode) []int {
	arr := []int{}
	for i := head; i != nil; i = i.Next {
		arr = append(arr, i.Val)
	}
	arrMaxIndex := len(arr)
	for i := 0; i < arrMaxIndex/2; i++ {
		arr[i], arr[arrMaxIndex-1-i] = arr[arrMaxIndex-1-i], arr[i]
	}
	return arr
}
