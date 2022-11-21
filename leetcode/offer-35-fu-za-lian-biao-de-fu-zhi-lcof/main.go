package main

import (
	"fmt"
	"math"
)

const (
	null = math.MinInt
)

//剑指 Offer 35. 复杂链表的复制
//请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

func main() {
	//[[3,null],[5,17],[4,null],[-9,6],[-10,3],[5,15],[0,11],[6,null],[-6,16],[3,16],[-6,11],[9,12],[-2,1],[-3,11],[-1,10],[2,11],[-3,null],[-9,7],[-2,4],[-8,null],[5,null]]
	node := SliceToNode([]int{3, 5, 4, -9, -10, 5, 0, 6, -6, 3, -6, 9, -2, -3, -1, 2, -3, -9, -2, -8, 5}, []int{null, 17, null, 6, 3, 15, 11, null, 16, 16, 11, 12, 1, 11, 10, 11, null, 7, 4, null, null})
	fmt.Println("输入：")
	node.EchoNode()
	fmt.Println("输出：")
	newNode := copyRandomList1(node)
	newNode.EchoNode()
}

type Node struct {
	Val    int
	Next   *Node // 指向下一个节点
	Random *Node // 指向任意节点或nil
}

func SliceToNode(s []int, r []int) *Node {
	ln := &Node{}
	tmp := ln
	for _, i := range s {
		// 带头节点
		tmp.Next = &Node{
			Val:  i,
			Next: nil,
		}
		// 带尾节点
		//tmp.Val = i
		//tmp.Next = &ListNode{}
		tmp = tmp.Next
	}

	for n, v := range r {
		var index int
		var nNode *Node
		var iNode *Node
		for i := ln.Next; i != nil; i = i.Next {
			if index == n {
				nNode = i
			}
			if index == v {
				iNode = i
			}
			index++
		}
		if nNode != nil && iNode != nil {
			nNode.Random = iNode
		}
	}

	return ln.Next
}

func (t *Node) EchoNode() {
	var index int
	for i := t; i != nil; i = i.Next {
		fmt.Printf("%d node[%[2]p]:%+[2]v\n", index, i)
		index++
	}
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	rondomSet := []int{}
	nodeSet := []*Node{}
	for i := head; i != nil; i = i.Next {
		nodeSet = append(nodeSet, &Node{
			Val:    i.Val,
			Next:   nil,
			Random: nil,
		})
		r := null
		if i.Random != nil {
			r = getNodeIndex(head, i.Random)
		}
		rondomSet = append(rondomSet, r)
	}
	MaxNode := len(nodeSet)
	for i, node := range nodeSet {
		if rondomSet[i] != null {
			node.Random = nodeSet[rondomSet[i]]
		}
		if i+1 < MaxNode {
			nodeSet[i].Next = nodeSet[i+1]
		}
	}
	return nodeSet[0]
}

func getNodeIndex(head *Node, target *Node) int {
	var index int
	for i := head; i != nil; i = i.Next {
		if i == target {
			return index
		}
		index++
	}
	return null
}

var cacheNode map[*Node]*Node

func copyRandomList1(head *Node) *Node {
	cacheNode = make(map[*Node]*Node)
	return deepCopy(head)
}

//递归获取节点的深度拷贝节点
//f(node): 获取node为根节点的深度拷贝节点
//f(node)=生成node的拷贝节点(包含node的val+node的next的拷贝节点+node的random的拷贝节点)，并把拷贝节点缓存到全局map中
func deepCopy(node *Node) *Node {
	if node == nil {
		return node
	}
	if n, has := cacheNode[node]; has {
		return n
	}
	newNode := &Node{Val: node.Val}
	cacheNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}
