package ch3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// SortStack sorts the elements in the given stack using a temporary stack as the
// only additional data structure.
func SortStack(s *Stack) {
	tempStack := Stack{}
	sortedItemCount := 0
	len := 0
	for !s.IsEmpty() {
		len++
		transferN(s, &tempStack, 1)
	}
	// Rewind.
	transfer(&tempStack, s)
	for i := 0; i < len; i++ {
		// Advance past the sorted items.
		transferN(s, &tempStack, sortedItemCount)
		// Find the next minimum.
		min, _ := s.Peek()
		for !s.IsEmpty() {
			item, _ := s.Peek()
			if item < min {
				min = item
			}
			transferN(s, &tempStack, 1)
		}
		// Rewind.
		transfer(&tempStack, s)
		// Advance past the sorted items.
		transferN(s, &tempStack, sortedItemCount)
		// Find the min and remove it.
		for {
			item, _ := s.Pop()
			if item == min {
				break
			}
			tempStack.Push(item)
		}
		// Rewind.
		transfer(&tempStack, s)
		// Advance past the sorted items.
		transferN(s, &tempStack, sortedItemCount)
		// Add the min.
		s.Push(min)
		sortedItemCount++
		// Rewind.
		transfer(&tempStack, s)
	}
}

// transfer moves n items from stack a to stack b.
func transferN(a, b *Stack, n int) error {
	for i := 0; i < n; i++ {
		item, err := a.Pop()
		if err != nil {
			return err
		}
		b.Push(item)
	}
	return nil
}

func Test3Dot5(t *testing.T) {
	q := Stack{}
	q.Push(3)
	q.Push(1)
	q.Push(4)
	q.Push(1)
	q.Push(5)
	SortStack(&q)
	item, _ := q.Pop()
	assert.Equal(t, 1, item)
	item, _ = q.Pop()
	assert.Equal(t, 1, item)
	item, _ = q.Pop()
	assert.Equal(t, 3, item)
	item, _ = q.Pop()
	assert.Equal(t, 4, item)
	item, _ = q.Pop()
	assert.Equal(t, 5, item)
}
