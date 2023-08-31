package list

import (
	"testing"
)

func TestLinkList_Add(t *testing.T) {

}

func TestLinkList_Get(t *testing.T) {
	linked := NewLinkedList[int]()
	linked.Add(1)
	t.Log()
	linked.Add(2)

}
