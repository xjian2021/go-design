package main

import (
	"fmt"
	. "github.com/xjian2021/go-design/leetcode/pkg"
)

//剑指 Offer 26. 树的子结构
//输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
//
//B是A的子结构， 即 A中有出现和B相同的结构和节点值。
/*
例如:
给定的树 A:
     3
    / \
   4   5
  / \
 1   2
给定的树 B：
   4
  /
 1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：
输入：A = [1,2,3], B = [3,1]
输出：false

示例 2：
输入：A = [3,4,5,1,2], B = [4,1]
输出：true
*/

func main() {
	//A := SliceToTreeNode([]int{3, 4, 1, NULL, NULL, 2, NULL, NULL, 5})
	//A := SliceToTreeNode([]int{1, 2, NULL, NULL, 3})
	A := SliceToTreeNode([]int{4, 2, 4, 8, NULL, NULL, 9, NULL, NULL, 5, NULL, NULL, 3, 6, NULL, NULL, 7})
	//A.EchoTreeNode()
	//B := SliceToTreeNode([]int{4, 1})
	//B := SliceToTreeNode([]int{3, 1})
	B := SliceToTreeNode([]int{4, 8, NULL, NULL, 9})
	//B.EchoTreeNode()
	fmt.Println(isSubStructure2(A, B))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// isSubStructure 队列+递归
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}
	aq := []*TreeNode{A}
	for len(aq) > 0 {
		if aq[0].Val == B.Val {
			//fmt.Println(B.Val)
			//aq[0].EchoTreeNode()
			//B.EchoTreeNode()
			//fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~")
			if contrastTree(aq[0].Left, B.Left) && contrastTree(aq[0].Right, B.Right) {
				return true
			}
		}
		if aq[0].Left != nil {
			aq = append(aq, aq[0].Left)
		}
		if aq[0].Right != nil {
			aq = append(aq, aq[0].Right)
		}
		aq = aq[1:]
	}
	return false
}

// isSubStructure2 递归+递归
func isSubStructure2(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}
	return contrastTree(A, B) || isSubStructure2(A.Left, B) || isSubStructure2(A.Right, B)
}

func contrastTree(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return contrastTree(A.Left, B.Left) && contrastTree(A.Right, B.Right)
}
