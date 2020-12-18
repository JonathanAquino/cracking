package ch2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func intersects(a *SinglyLinkedList, b *SinglyLinkedList) bool {
	aCurrent := a.head
	for aCurrent != nil {
		bCurrent := b.head
		for bCurrent != nil {
			if aCurrent == bCurrent {
				return true
			}
			bCurrent = bCurrent.next
		}
		aCurrent = aCurrent.next
	}
	return false
}

func Test2Dot7a(t *testing.T) {
	a := SinglyLinkedList{}
	a.Add(7)
	a.Add(1)
	a.Add(6)
	b := SinglyLinkedList{}
	b.Add(7)
	b.Add(1)
	b.Add(6)
	assert.False(t, intersects(&a, &b))
}

func Test2Dot7b(t *testing.T) {
	a := SinglyLinkedList{}
	a.Add(7)
	a.Add(1)
	a.Add(6)
	b := SinglyLinkedList{}
	b.Add(7)
	b.Add(1)
	b.Tail().next = a.Tail()
	assert.True(t, intersects(&a, &b))
}
