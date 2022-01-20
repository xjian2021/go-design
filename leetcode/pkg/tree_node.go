package pkg

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

// TODO
//SliceToTreeNode 前序遍历
func SliceToTreeNode(s []int) *TreeNode {
	var lr int
	t := &TreeNode{}
	tmp := t
	var mockStack []*TreeNode
	for _, i := range s {
		if i > 0 {
			if lr%2 == 0 {

			}
			mockStack = append(mockStack, tmp)
			tmp.Val = i
			tmp.Left = &TreeNode{}
			tmp = tmp.Left
		} else {
			lr++
			if lr >= 2 {
				lr -= 2
				mockStack = mockStack[:len(mockStack)-1]
				tmp = mockStack[len(mockStack)-1]
			}
		}

		//if i > 0 {
		//	if lr == 0 {
		//		tmp.Left = &TreeNode{Val: i}
		//		tmp = tmp.Left
		//	} else {
		//		lr--
		//		if lr == 0 {
		//			tmp.Right = &TreeNode{Val: i}
		//			tmp = tmp.Right
		//		} else {
		//
		//		}
		//	}
		//} else {
		//	lr++
		//}
	}
	return t.Left
}

func s2t(s []int) *TreeNode {
	var mockStack []*TreeNode
	t := &TreeNode{}
	tmp := t
	ms := len(s)
	for i := 0; i < ms; i++ {
		mockStack = append(mockStack, tmp)
		tmp.Val = s[i]
		i++
		if i < ms && s[i] > 0 {
			tmp.Left = &TreeNode{Val: s[i]}
		}
		i++
		if i < ms && s[i] > 0 {
			tmp.Right = &TreeNode{Val: s[i]}
		}
		tmp = nil
	}
	return t
}
