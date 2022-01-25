package pkg

import (
	"testing"
)

func TestS2T(t *testing.T) {
	//n := []int{1, 2}
	//n := []int{1, NULL, 3, 2}
	//n := []int{6, 5, 3, 2, 4, NULL, NULL, 8, NULL, NULL, 1}
	n := []int{5, 4, 3, 2, NULL, NULL, NULL, NULL, 1}
	tNode := SliceToTreeNode(n)
	tNode.EchoTreeNode()
}
