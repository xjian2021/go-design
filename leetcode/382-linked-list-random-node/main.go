package main

import (
	"math/rand"

	. "github.com/xjian2021/go-design/leetcode/pkg"
)

func main() {
	node := SliceToListNode([]int{1, 2, 3})
	constructor := Constructor(node.Next)
	println(constructor.GetRandom())
	println(constructor.GetRandom())
	println(constructor.GetRandom())
	println(constructor.GetRandom())
	println(constructor.GetRandom())
	println(constructor.GetRandom())
	println(constructor.GetRandom())
	println(constructor.GetRandom())
	//fmt.Printf("listnode:%+v\n", node.Next)
	//for i := node; i != nil; i = i.Next {
	//	fmt.Printf("i:%+v\n", i)
	//}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	Len  int
	Head *ListNode
}

func Constructor(head *ListNode) Solution {
	var l int
	for i := head; i != nil; i = i.Next {
		l++
	}
	return Solution{Head: head, Len: l}
}

func (this *Solution) GetRandom() int {
	num := rand.Intn(this.Len)
	for i := this.Head; i != nil; i = i.Next {
		if num == 0 {
			return i.Val
		}
		num--
	}
	return this.Head.Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
