package ch3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TwoStackQueue is a queue implemented using two stacks
type TwoStackQueue struct {
	mainStack *Stack
	tempStack *Stack
}

func NewTwoStackQueue() *TwoStackQueue {
	return &TwoStackQueue{mainStack: &Stack{}, tempStack: &Stack{}}
}

func (q *TwoStackQueue) IsEmpty() bool {
	return q.mainStack.IsEmpty()
}

func (q *TwoStackQueue) Add(item int) {
	q.mainStack.Push(item)
}

func (q *TwoStackQueue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("can't peek into an empty queue")
	}
	err := q.transfer(q.mainStack, q.tempStack)
	if err != nil {
		return 0, err
	}
	item, err := q.tempStack.Peek()
	if err != nil {
		return 0, err
	}
	err = q.transfer(q.tempStack, q.mainStack)
	if err != nil {
		return 0, err
	}
	return item, nil
}

func (q *TwoStackQueue) Remove() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("can't remove from an empty queue")
	}
	err := q.transfer(q.mainStack, q.tempStack)
	if err != nil {
		return 0, err
	}
	item, err := q.tempStack.Pop()
	if err != nil {
		return 0, err
	}
	err = q.transfer(q.tempStack, q.mainStack)
	if err != nil {
		return 0, err
	}
	return item, nil
}

// transfer moves the items from stack a to stack b.
func (q *TwoStackQueue) transfer(a, b *Stack) error {
	for !a.IsEmpty() {
		item, err := a.Pop()
		if err != nil {
			return err
		}
		b.Push(item)
	}
	return nil
}

func Test3Dot4(t *testing.T) {
	q := NewTwoStackQueue()
	assert.True(t, q.IsEmpty())
	q.Add(1)
	assert.False(t, q.IsEmpty())
	q.Add(2)
	q.Add(3)
	item, err := q.Peek()
	assert.Equal(t, 1, item)
	item, err = q.Remove()
	assert.Equal(t, 1, item)
	item, err = q.Peek()
	assert.Equal(t, 2, item)
	item, err = q.Remove()
	assert.Equal(t, 2, item)
	q.Add(4)
	item, err = q.Peek()
	assert.Equal(t, 3, item)
	item, err = q.Remove()
	assert.Equal(t, 3, item)
	item, err = q.Peek()
	assert.Equal(t, 4, item)
	item, err = q.Remove()
	assert.Equal(t, 4, item)
	assert.True(t, q.IsEmpty())
	item, err = q.Peek()
	assert.NotNil(t, err)
	item, err = q.Remove()
	assert.NotNil(t, err)
}
