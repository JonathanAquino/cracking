package ch2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// removeNode removes a middle (not first or last) node from a singly linked list,
// given only that node.
func removeNode(node *SinglyLinkedListNode) {
	current := node
	for {
		current.Data = current.Next.Data
		if current.Next.Next == nil {
			current.Next = nil
			break
		}
		current = current.Next
	}
}

func Test2Dot3(t *testing.T) {
	l := SinglyLinkedList{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	thirdNode := l.Tail()
	l.Add(4)
	l.Add(5)
	l.Add(6)
	removeNode(thirdNode)
	assert.Equal(t, []int{1, 2, 4, 5, 6}, ToArray(l))
}
