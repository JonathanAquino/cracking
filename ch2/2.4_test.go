package ch2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// partition partitions a linked list into values < threshold and values >= threshold.
func partition(l *SinglyLinkedList, threshold int) *SinglyLinkedList {
	// Less than list
	ltList := SinglyLinkedList{}
	// Greater than or equal list
	gteList := SinglyLinkedList{}
	current := l.Head
	for current != nil {
		if current.Data < threshold {
			ltList.Add(current.Data)
		} else {
			gteList.Add(current.Data)
		}
		current = current.Next
	}
	if ltList.Head == nil {
		return &gteList
	}
	ltList.Tail().Next = gteList.Head
	return &ltList
}

func Test2Dot4(t *testing.T) {
	l := SinglyLinkedList{}
	l.Add(3)
	l.Add(5)
	l.Add(8)
	l.Add(5)
	l.Add(10)
	l.Add(2)
	l.Add(1)
	actual := partition(&l, 5)
	assert.Equal(t, []int{3, 2, 1, 5, 8, 5, 10}, ToArray(*actual))
}
