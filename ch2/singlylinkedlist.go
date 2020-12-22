package ch2

type SinglyLinkedList struct {
	Head *SinglyLinkedListNode
}

type SinglyLinkedListNode struct {
	Data int
	Next *SinglyLinkedListNode
}

// Add adds the given value to the linked list.
func (l *SinglyLinkedList) Add(data int) *SinglyLinkedList {
	if l.Head == nil {
		l.Head = &SinglyLinkedListNode{Data: data}
		return l
	}
	l.Tail().Next = &SinglyLinkedListNode{Data: data}
	return l
}

// Length returns the length of the list.
func (l *SinglyLinkedList) Length() int {
	if l.Head == nil {
		return 0
	}
	length := 1
	current := l.Head
	for {
		if current.Next == nil {
			return length
		}
		length++
		current = current.Next
	}
}

// At returns the node at the given index.
func (l *SinglyLinkedList) At(i int) *SinglyLinkedListNode {
	// TODO: Return an error if i > len - 1.
	j := 0
	current := l.Head
	for {
		if i == j {
			return current
		}
		j++
		current = current.Next
	}
}

// Tail returns the last element in the linked list.
func (l *SinglyLinkedList) Tail() *SinglyLinkedListNode {
	if l.Head == nil {
		return nil
	}
	current := l.Head
	for {
		if current.Next == nil {
			return current
		}
		current = current.Next
	}
}

// ToArray converts the linked list to an array.
func ToArray(l SinglyLinkedList) []int {
	a := []int{}
	if l.Head != nil {
		current := l.Head
		for current != nil {
			a = append(a, current.Data)
			current = current.Next
		}
	}
	return a
}
