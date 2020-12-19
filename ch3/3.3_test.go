package ch3

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// SetOfStacks is a group of stacks with a size threshold and a PopAt(n) function
// for popping a specific stack.
type SetOfStacks struct {
	threshold int
	stacks    *ArrayStack
}

func NewSetOfStacks(threshold int) *SetOfStacks {
	return &SetOfStacks{threshold: threshold, stacks: &ArrayStack{}}
}

func (s *SetOfStacks) Push(item interface{}) {
	if s.stacks.IsEmpty() {
		s.stacks.Push(&ArrayStack{})
		lastStack, _ := s.stacks.Peek()
		lastStack.(*ArrayStack).Push(item)
		return
	}
	lastStack, _ := s.stacks.Peek()
	if lastStack.(*ArrayStack).Len() == s.threshold {
		s.stacks.Push(&ArrayStack{})
		lastStack, _ := s.stacks.Peek()
		lastStack.(*ArrayStack).Push(item)
		return
	}
	lastStack.(*ArrayStack).Push(item)
}

func (s *SetOfStacks) PopAt(index int) (interface{}, error) {
	stack, err := s.stacks.Get(index)
	if err != nil {
		return nil, err
	}
	return stack.(*ArrayStack).Pop()
}

func (s *SetOfStacks) Pop() (interface{}, error) {
	s.trim()
	if s.stacks.IsEmpty() {
		return 0, errors.New("can't pop an empty stack")
	}
	lastStack, _ := s.stacks.Peek()
	return lastStack.(*ArrayStack).Pop()
}

// Removes empty stacks from the top of the SetOfStacks
func (s *SetOfStacks) trim() {
	for {
		if s.stacks.IsEmpty() {
			return
		}
		top, _ := s.stacks.Peek()
		if !top.(*ArrayStack).IsEmpty() {
			return
		}
		s.stacks.Pop()
	}
}

func Test3Dot3(t *testing.T) {
	s := NewSetOfStacks(3)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	item, _ := s.PopAt(0)
	assert.Equal(t, 3, item)
	item, _ = s.PopAt(0)
	assert.Equal(t, 2, item)
	item, _ = s.Pop()
	assert.Equal(t, 5, item)
	item, _ = s.Pop()
	assert.Equal(t, 4, item)
	item, _ = s.Pop()
	assert.Equal(t, 1, item)
}
