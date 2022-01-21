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
	tmp := t
	ms := len(s)
	if ms == 0 {
		return nil
	}
	var treeNodes []*TreeNode
	for i := 0; i < ms; i++ {
		var node *TreeNode
		if s[i] > 0 {
			node = &TreeNode{Val: s[i]}
		}
		treeNodes = append(treeNodes, node)
	}
	if treeNodes == nil {
		panic("数组不构成树")
	}
	var skip int
	for i := 0; i < ms; i++ {
		no := treeNodes[i]
		if s[i] > 0 {
			switch skip {
			case 0:
				tmp.Left = no
			case 1:
				tmp.Right = no
				treeNodes[i-2] = nil
			default: // skip > 2
				j := i - 4
				for ; treeNodes[j] == nil; j-- {
				}
				if j < 0 {
					panic("数组不构成树")
				}
				tmp = treeNodes[j]
				tmp.Right = no
				treeNodes[j] = nil
			}
			tmp = no
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
