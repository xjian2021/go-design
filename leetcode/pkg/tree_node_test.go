package pkg

import (
	"testing"
)

func TestS2T(t *testing.T) {
	//n := []int{1, 2}
	//n := []int{1, -1, 3, 2}
	//n := []int{6, 5, 3, 2, 4, -1, -1, 8, -1, -1, 1}
	n := []int{5, 4, 3, 2, -1, -1, -1, -1, 1}
	tNode := SliceToTreeNode(n)
	tNode.EchoTreeNode()
}
