package ch2

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
}

type SinglyLinkedListNode struct {
	data int
	next *SinglyLinkedListNode
}

// Add adds the given value to the linked list.
func (l *SinglyLinkedList) Add(data int) *SinglyLinkedList {
	if l.head == nil {
		l.head = &SinglyLinkedListNode{data: data}
		return l
	}
	l.Tail().next = &SinglyLinkedListNode{data: data}
	return l
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

// At returns the node at the given index.
func (l *SinglyLinkedList) At(i int) *SinglyLinkedListNode {
	// TODO: Return an error if i > len - 1.
	j := 0
	current := l.head
	for {
		if i == j {
			return current
		}
		j++
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

// ToArray converts the linked list to an array.
func ToArray(l SinglyLinkedList) []int {
	a := []int{}
	if l.head != nil {
		current := l.head
		for current != nil {
			a = append(a, current.data)
			current = current.next
		}
	}
	return a
}
