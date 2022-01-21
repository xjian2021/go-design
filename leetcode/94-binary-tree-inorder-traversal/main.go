package main

import (
	"fmt"

	. "github.com/xjian2021/go-design/leetcode/pkg"
)

// 中序遍历二叉树
func main() {
	//after := time.After(1 * time.Second)
	//go func() {
	//	select {
	//	case <-after:
	//		fmt.Println("time out")
	//		os.Exit(10086)
	//	}
	//}()
	root := []int{3, 2, 4, 0, 0, 1}
	t := SliceToTreeNode(root)
	//t := NewTreeNode(3, NewTreeNode(2, NewTreeNode(4, nil, nil), NewTreeNode(1, nil, nil)), nil)
	t.EchoTreeNode()
	s := inorderTraversalByMorris(t)
	fmt.Println(s)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var s []int
	//s = dg(root, s)
	s = stack(root)
	return s
}

//dg 递归中序遍历
func dg(node *TreeNode, a []int) []int {
	if node == nil {
		return a
	}
	fmt.Printf("node:%+v a:%v\n", node, a)
	a = dg(node.Left, a)
	a = append(a, node.Val)
	a = dg(node.Right, a)
	return a
}

//stack 栈模拟递归中序遍历
func stack(root *TreeNode) (s []int) {
	//遍历条件：
	//栈中还有元素
	//协助节点不为nil 找左节点 没有就append 再右节点赋值给协助节点
	var mockStack []*TreeNode
	for root != nil || len(mockStack) > 0 {
		if root == nil {
			root = mockStack[len(mockStack)-1]
			mockStack = mockStack[:len(mockStack)-1]
			root.Left = nil
		}
		if root.Left == nil {
			fmt.Printf("tmp:%+v s:%v\n", root, s)
			s = append(s, root.Val)
			if root.Right != nil {
				root = root.Right
			} else {
				root = nil
			}
		} else {
			mockStack = append(mockStack, root)
			root = root.Left
		}
	}
	return s
}

//inorderTraversal1 官方栈式中序遍历
func inorderTraversal1(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

//inorderTraversalByMorris Morris中序遍历
// Morris算法关键在于让x的左子树的最右节点的右指针指回x 从而形成一个环
// 在左子树遍历完之后通过右指针来回到父节点 循环
func inorderTraversalByMorris(root *TreeNode) (res []int) {
	for root != nil {
		if root.Left != nil {
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}
			if predecessor.Right == root {
				res = append(res, root.Val)
				predecessor.Right = nil
				root = root.Right
			} else {
				predecessor.Right = root
				root = root.Left
			}
		} else {
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return res
}
