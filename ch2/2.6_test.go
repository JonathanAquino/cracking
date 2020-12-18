package ch2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPalindrome(l *SinglyLinkedList) bool {
	len := l.Length()
	// Integer division rounds down.
	for i := 0; i < len/2; i++ {
		if l.At(i).data != l.At(len-1-i).data {
			return false
		}
	}
	return true
}

func Test2Dot6a(t *testing.T) {
	l := SinglyLinkedList{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(2)
	l.Add(1)
	assert.True(t, isPalindrome(&l))
}

func Test2Dot6b(t *testing.T) {
	l := SinglyLinkedList{}
	l.Add(1)
	l.Add(2)
	l.Add(2)
	l.Add(1)
	assert.True(t, isPalindrome(&l))
}

func Test2Dot6c(t *testing.T) {
	l := SinglyLinkedList{}
	l.Add(1)
	l.Add(2)
	l.Add(2)
	l.Add(2)
	assert.False(t, isPalindrome(&l))
}
