package ch2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// kthToLast finds the kth to last element of a linked list.
func kthToLast(l *SinglyLinkedList, k int) *SinglyLinkedListNode {
	slow := l.Head
	fast := l.Head
	for i := 0; i < k-1; i++ {
		fast = fast.Next
		if fast == nil {
			return nil
		}
	}
	for {
		fast = fast.Next
		if fast == nil {
			break
		}
		slow = slow.Next
	}
	return slow
}

func Test2Dot2(t *testing.T) {
	l := SinglyLinkedList{}
	l.Add(3)
	l.Add(1)
	l.Add(4)
	l.Add(1)
	l.Add(5)
	l.Add(9)
	l.Add(2)
	l.Add(6)
	l.Add(5)
	assert.Equal(t, 5, kthToLast(&l, 1).Data)
	assert.Equal(t, 6, kthToLast(&l, 2).Data)
	assert.Equal(t, 2, kthToLast(&l, 3).Data)
	assert.Nil(t, kthToLast(&l, 100))
}
