package ch3

import "errors"

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
}

type SinglyLinkedListNode struct {
	data interface{}
	next *SinglyLinkedListNode
}

// Add adds the given value to the linked list.
func (l *SinglyLinkedList) Add(data interface{}) *SinglyLinkedList {
	if l.head == nil {
		l.head = &SinglyLinkedListNode{data: data}
		return l
	}
	l.Tail().next = &SinglyLinkedListNode{data: data}
	return l
}

// IsEmpty returns whether the list is empty.
func (l *SinglyLinkedList) IsEmpty() bool {
	return l.Length() == 0
}

// Length returns the length of the list.
func (l *SinglyLinkedList) Length() int {
	if l.head == nil {
		return 0
	}
	length := 1
	current := l.head
	for {
		if current.next == nil {
			return length
		}
		length++
		current = current.next
	}
}

// Tail returns the last element in the linked list.
func (l *SinglyLinkedList) Tail() *SinglyLinkedListNode {
	if l.head == nil {
		return nil
	}
	current := l.head
	for {
		if current.next == nil {
			return current
		}
		current = current.next
	}
}

// Shift returns the head of the list and removes it.
func (l *SinglyLinkedList) Shift() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("can't shift empty list")
	}
	data := l.head.data
	l.head = l.head.next
	return data, nil
}
