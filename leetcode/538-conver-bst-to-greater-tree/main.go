package main

import (
	. "github.com/xjian2021/go-design/leetcode/pkg"
)

//538. 把二叉搜索树转换为累加树
//给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
//
//提醒一下，二叉搜索树满足下列约束条件：
//
//节点的左子树仅包含键 小于 节点键的节点。
//节点的右子树仅包含键 大于 节点键的节点。
//左右子树也必须是二叉搜索树。
func main() {
	n := []int{4, 1, 0, -1, -1, 2, -1, 3, -1, -1, 6, 5, -1, -1, 7, -1, 8}
	root := SliceToTreeNode(n)
	t := convertBST(root)
	t.EchoTreeNode()
}

/**
 * 后续遍历二叉树
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	num := 0
	dg(root, &num)
	return root
}

func dg(root *TreeNode, val *int) {
	if root == nil {
		return
	}
	dg(root.Right, val)
	root.Val += *val
	*val = root.Val
	dg(root.Left, val)
}
