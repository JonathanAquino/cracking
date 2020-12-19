package ch3

import (
	"errors"
)

type Stack struct {
	top *stackNode
}

type stackNode struct {
	data int
	next *stackNode
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("can't peek into an empty stack")
	}
	return s.top.data, nil
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("can't peek into an empty stack")
	}
	oldTop := s.top
	s.top = oldTop.next
	return oldTop.data, nil
}

func (s *Stack) Push(data int) {
	newTop := stackNode{data: data, next: s.top}
	s.top = &newTop
}
