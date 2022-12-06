package main

import (
	"fmt"
	. "github.com/xjian2021/go-design/leetcode/pkg"
)

//剑指 Offer 32 - I. 从上到下打印二叉树
//从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
/*
例如:
	给定二叉树: [3,9,20,null,null,15,7],

		3
	   / \
	  9  20
		/  \
	   15   7
	返回：

	[3,9,20,15,7]
*/
// 打印前缀树
func main() {
	root := SliceToTreeNode([]int{3, 9, NULL, NULL, 20, 15, 7})
	root.EchoTreeNode()
	fmt.Println(levelOrder(root))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
	a := []int{}
	if root == nil {
		return a
	}
	q := []*TreeNode{root}
	a = append(a, root.Val)
	for len(q) > 0 {
		tmp := q[0]
		// TODO 把元素加进去???
		if tmp.Left != nil {
			a = append(a, tmp.Left.Val)
			if tmp.Left.Left != nil {
				q = append(q, tmp.Left.Left)
			}
			if tmp.Left.Right != nil {
				q = append(q, tmp.Left.Right)
			}
		}
		if tmp.Right != nil {
			a = append(a, tmp.Right.Val)
			if tmp.Right.Left != nil {
				q = append(q, tmp.Right.Left)
			}
			if tmp.Right.Right != nil {
				q = append(q, tmp.Right.Right)
			}
		}
		q = q[1:]
	}
	return a
}
