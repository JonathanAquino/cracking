package ch2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// removeDuplicates removes duplicates from an unsorted linked list.
func removeDuplicates(l *SinglyLinkedList) {
	slow := l.Head
	for slow != nil {
		prevFast := slow
		fast := slow.Next
		for fast != nil {
			nextFast := fast.Next
			if slow.Data == fast.Data {
				prevFast.Next = fast.Next
				fast.Next = nil
			}
			prevFast = fast
			fast = nextFast
		}
		slow = slow.Next
	}
}

func Test2Dot1(t *testing.T) {
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
	removeDuplicates(&l)
	assert.Equal(t, []int{3, 1, 4, 5, 9, 2, 6}, ToArray(l))
}
