package pkg

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

//EchoTreeNode 以json格式化打印树
func (t *TreeNode) EchoTreeNode() {
	b, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		fmt.Printf("json.MarshalIndent fail err:%s\n", err.Error())
		return
	}
	fmt.Printf("tree:\n%s\n", string(b))
}

//SliceToTreeNode 前序遍历方式生成树
func SliceToTreeNode(s []int) *TreeNode {
	t := &TreeNode{}
	tmp, ms, skip := t, len(s), 0
	if ms == 0 {
		return nil
	}
	treeNodes := make([]*TreeNode, ms)
	for i := 0; i < ms; i++ {
		if s[i] > 0 {
			treeNodes[i] = &TreeNode{Val: s[i]}
			switch skip {
			case 0:
				tmp.Left = treeNodes[i]
			case 1:
				tmp.Right = treeNodes[i]
				treeNodes[i-2] = nil
			default:
				j := i - 4
				for ; treeNodes[j] == nil; j-- {
				}
				if j < 0 {
					panic("数组不构成树")
				}
				tmp = treeNodes[j]
				tmp.Right = treeNodes[i]
				treeNodes[j] = nil
			}
			tmp = treeNodes[i]
			skip = 0
		} else {
			skip++
			if skip > 1 {
				treeNodes[i-skip*2+2] = nil
			}
		}
	}

	return t.Left
}
