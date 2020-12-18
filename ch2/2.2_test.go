package ch2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func kthToLast(l *SinglyLinkedList, k int) *SinglyLinkedListNode {
	slow := l.head
	fast := l.head
	for i := 0; i < k-1; i++ {
		fast = fast.next
		if fast == nil {
			return nil
		}
	}
	for {
		fast = fast.next
		if fast == nil {
			break
		}
		slow = slow.next
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
	assert.Equal(t, 5, kthToLast(&l, 1).data)
	assert.Equal(t, 6, kthToLast(&l, 2).data)
	assert.Equal(t, 2, kthToLast(&l, 3).data)
	assert.Nil(t, kthToLast(&l, 100))
}
