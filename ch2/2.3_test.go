package ch2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeNode(node *SinglyLinkedListNode) {
	current := node
	for {
		current.data = current.next.data
		if current.next.next == nil {
			current.next = nil
			break
		}
		current = current.next
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
