package ch3

import (
	"errors"
	"fmt"
)

// An ArrayStack is a stack implemented as an array. You can do Get(n) to get the
// nth item.
type ArrayStack struct {
	data []interface{}
}

func (s *ArrayStack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return 0, errors.New("can't peek into an empty stack")
	}
	return s.data[len(s.data)-1], nil
}

func (s *ArrayStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *ArrayStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return 0, errors.New("can't pop an empty stack")
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top, nil
}

func (s *ArrayStack) Push(item interface{}) {
	s.data = append(s.data, item)
}

func (s *ArrayStack) Get(index int) (interface{}, error) {
	if index < 0 {
		return nil, fmt.Errorf("Array index out of bounds: %d", index)
	}
	if index >= len(s.data) {
		return nil, fmt.Errorf("Array index out of bounds: %d", index)
	}
	return s.data[index], nil
}

func (s *ArrayStack) Len() int {
	return len(s.data)
}
