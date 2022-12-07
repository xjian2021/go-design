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
	root := SliceToTreeNode([]int{3, 9, 8, NULL, NULL, NULL, 20, 15, NULL, NULL, 7})
	root.EchoTreeNode()
	//fmt.Println(levelOrder(root))
	fmt.Println(levelOrder3(root))
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
	// 利用队列的先入先出的特性把每层的节点从左往右插入数组中
	q := []*TreeNode{root}
	for len(q) > 0 {
		a = append(a, q[0].Val)
		if q[0].Left != nil {
			q = append(q, q[0].Left)
		}
		if q[0].Right != nil {
			q = append(q, q[0].Right)
		}
		q = q[1:]
	}
	return a
}

//剑指 Offer 32 - II. 从上到下打印二叉树 II
//从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
/*
例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder2(root *TreeNode) [][]int {
	a := [][]int{}
	if root == nil {
		return a
	}
	q := []*TreeNode{root}
	var i int
	for len(q) > 0 {
		a = append(a, []int{})
		var nextQ []*TreeNode
		for _, treeNode := range q {
			a[i] = append(a[i], treeNode.Val)
			if treeNode.Left != nil {
				nextQ = append(nextQ, treeNode.Left)
			}
			if treeNode.Right != nil {
				nextQ = append(nextQ, treeNode.Right)
			}
		}
		q = nextQ
		i++
	}
	return a
}

//剑指 Offer 32 - III. 从上到下打印二叉树 III
//请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
/*
例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [20,9],
  [15,7]
]
*/
func levelOrder3(root *TreeNode) [][]int {
	a := [][]int{}
	if root == nil {
		return a
	}
	q := []*TreeNode{root}
	var i int
	for len(q) > 0 {
		b := make([]int, 0, len(q))
		var nextQ []*TreeNode
		for _, treeNode := range q {
			if i%2 == 0 {
				b = append(b, treeNode.Val)
			} else {
				b = append([]int{treeNode.Val}, b...)
			}
			if treeNode.Left != nil {
				nextQ = append(nextQ, treeNode.Left)
			}
			if treeNode.Right != nil {
				nextQ = append(nextQ, treeNode.Right)
			}
		}
		a = append(a, b)
		q = nextQ
		i++
	}
	return a
}
