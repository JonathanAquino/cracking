package ch2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// findLoop determines whether the linked list has a loop. If so, it returns the
// node at the start of the loop.
func findLoop(l *SinglyLinkedList) *SinglyLinkedListNode {
	// Have two pointers: a slow one which iterates over the list, and a
	// fast one which keeps starting from the beginning and goes to the slow one.
	// We keep track of prevSlow and prevFast which tells us that a loop is detected
	// when they are different when the fast one meets the slow one.
	if l.Head == nil {
		return nil
	}
	var prevSlow *SinglyLinkedListNode
	slow := l.Head
	for slow != nil {
		var prevFast *SinglyLinkedListNode
		fast := l.Head
		for fast != nil {
			if slow == fast && prevSlow != prevFast {
				// Loop detected
				return slow
			}
			if slow == fast {
				// No loop detected yet
				break
			}
			prevFast = fast
			fast = fast.Next
		}
		prevSlow = slow
		slow = slow.Next
	}
	return nil
}

func Test2Dot8a(t *testing.T) {
	l := SinglyLinkedList{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Tail().Next = l.Head.Next.Next
	assert.Equal(t, l.Head.Next.Next, findLoop(&l))
}

func Test2Dot8b(t *testing.T) {
	l := SinglyLinkedList{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	assert.Nil(t, findLoop(&l))
}
