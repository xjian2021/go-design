package pkg

import (
	"testing"
)

func TestS2T(t *testing.T) {
	//n := []int{1, 2}
	//n := []int{1, 0, 3, 2}
	//n := []int{6, 5, 3, 2, 4, 0, 0, 8, 0, 0, 1}
	n := []int{5, 4, 3, 2, 0, 0, 0, 0, 1}
	tNode := SliceToTreeNode(n)
	tNode.EchoTreeNode()
}
