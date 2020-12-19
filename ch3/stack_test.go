package ch3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	result, err := s.Peek()
	assert.Equal(t, 3, result)
	result, err = s.Pop()
	assert.Equal(t, 3, result)
	result, err = s.Peek()
	assert.False(t, s.IsEmpty())
	result, err = s.Pop()
	result, err = s.Pop()
	assert.True(t, s.IsEmpty())
	result, err = s.Pop()
	assert.NotNil(t, err)
}
