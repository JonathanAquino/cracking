package ch3

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// StackWithMin is a Stack with an O(1) Min operation
type StackWithMin struct {
	mainStack *Stack
	minStack  *Stack
}

func NewStackWithMin() StackWithMin {
	return StackWithMin{mainStack: &Stack{}, minStack: &Stack{}}
}

func (s *StackWithMin) Peek() (int, error) {
	return s.mainStack.Peek()
}

func (s *StackWithMin) Min() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("can't do min on an empty stack")
	}
	return s.minStack.Peek()
}

func (s *StackWithMin) IsEmpty() bool {
	return s.mainStack.IsEmpty()
}

func (s *StackWithMin) Pop() (int, error) {
	result, err := s.mainStack.Pop()
	if err != nil {
		return result, err
	}
	min, err := s.minStack.Peek()
	if err != nil {
		return 0, err
	}
	if result == min {
		s.minStack.Pop()
	}
	return result, nil
}

func (s *StackWithMin) Push(data int) {
	if s.IsEmpty() {
		s.minStack.Push(data)
	}
	min, _ := s.minStack.Peek()
	if data <= min {
		s.minStack.Push(data)
	}
	s.mainStack.Push(data)
}

func Test3Dot2(t *testing.T) {
	s := NewStackWithMin()
	s.Push(4)
	result, err := s.Min()
	assert.Equal(t, 4, result)
	s.Push(2)
	result, err = s.Min()
	assert.Equal(t, 2, result)
	s.Push(3)
	result, err = s.Min()
	assert.Equal(t, 2, result)
	s.Push(1)
	result, err = s.Min()
	assert.Equal(t, 1, result)

	result, err = s.Pop()
	assert.Equal(t, 1, result)
	result, err = s.Peek()
	assert.Equal(t, 3, result)
	result, err = s.Min()
	assert.Equal(t, 2, result)
	assert.False(t, s.IsEmpty())

	result, err = s.Pop()
	assert.Equal(t, 3, result)
	result, err = s.Peek()
	assert.Equal(t, 2, result)
	result, err = s.Min()
	assert.Equal(t, 2, result)
	assert.False(t, s.IsEmpty())

	result, err = s.Pop()
	assert.Equal(t, 2, result)
	result, err = s.Peek()
	assert.Equal(t, 4, result)
	result, err = s.Min()
	assert.Equal(t, 4, result)
	assert.False(t, s.IsEmpty())

	result, err = s.Pop()
	assert.Equal(t, 4, result)
	result, err = s.Peek()
	assert.NotNil(t, err)
	result, err = s.Min()
	assert.NotNil(t, err)
	assert.True(t, s.IsEmpty())
}
