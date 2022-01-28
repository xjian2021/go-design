package pkg

import "testing"

func TestSliceToTwoWayListNode(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	node := SliceToTwoWayListNode(s)
	node.EchoTwoWayListNode()
}

func TestSliceToListNode(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	node := SliceToListNode(s)
	node.EchoListNode()
}
