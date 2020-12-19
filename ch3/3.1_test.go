package ch3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// NewMultiStack creates a MultiStack, which is a group of n stacks implemented
// using a single array.
func NewMultiStack(n int) *MultiStack {
	m := MultiStack{n: n}
	m.topIndices = make([]int, n)
	for i := 0; i < n; i++ {
		m.topIndices[i] = -1
	}
	return &m
}

type MultiStack struct {
	n          int
	data       []string
	topIndices []int
}

func (m *MultiStack) IsEmpty(index int) bool {
	return m.topIndices[index] == -1
}

func (m *MultiStack) Peek(index int) (string, error) {
	if m.topIndices[index] == -1 {
		return "", fmt.Errorf("stack %d is empty and cannot be peeked", index)
	}
	return m.data[m.topIndices[index]*m.n+index], nil
}

func (m *MultiStack) Pop(index int) (string, error) {
	if m.topIndices[index] == -1 {
		return "", fmt.Errorf("stack %d is empty and cannot be popped", index)
	}
	result := m.data[m.topIndices[index]*m.n+index]
	m.topIndices[index]--
	return result, nil
}

func (m *MultiStack) Push(index int, data string) {
	m.topIndices[index]++
	location := m.topIndices[index]*m.n + index
	for i := len(m.data) - 1; i < location; i++ {
		m.data = append(m.data, "")
	}
	m.data[location] = data
}

func Test3Dot1(t *testing.T) {
	m := NewMultiStack(3)
	m.Push(0, "a")
	m.Push(0, "b")
	m.Push(0, "c")
	m.Push(1, "d")
	m.Push(1, "e")
	m.Push(2, "f")
	result, err := m.Peek(0)
	assert.Equal(t, "c", result)
	result, err = m.Peek(1)
	assert.Equal(t, "e", result)
	result, err = m.Peek(2)
	assert.Equal(t, "f", result)
	result, err = m.Pop(0)
	assert.Equal(t, "c", result)
	result, err = m.Pop(1)
	assert.Equal(t, "e", result)
	result, err = m.Pop(2)
	assert.Equal(t, "f", result)
	result, err = m.Peek(0)
	assert.Equal(t, "b", result)
	result, err = m.Peek(1)
	assert.Equal(t, "d", result)
	result, err = m.Peek(2)
	assert.NotNil(t, err)
	assert.False(t, m.IsEmpty(0))
	assert.False(t, m.IsEmpty(1))
	assert.True(t, m.IsEmpty(2))
}
