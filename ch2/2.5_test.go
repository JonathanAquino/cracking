package ch2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// sumLists adds two numbers represented by a linked list. For example, 24 + 8 = 32
// is represented by (4 -> 2) + (8) = (2 -> 3).
func sumLists(a *SinglyLinkedList, b *SinglyLinkedList) *SinglyLinkedList {
	sumList := SinglyLinkedList{}
	aCurrent := a.head
	bCurrent := b.head
	carryOne := false
	for aCurrent != nil {
		sum := aCurrent.data + bCurrent.data
		if carryOne {
			sum += 1
		}
		carryOne = sum >= 10
		if carryOne {
			sum -= 10
		}
		sumList.Add(sum)
		aCurrent = aCurrent.next
		bCurrent = bCurrent.next
	}
	if carryOne {
		sumList.Add(1)
	}
	return &sumList
}

// sumListsReverse does the same as sumLists but in a more natural (reversed) order.
// For example, 24 + 8 = 32 is represented by (2 -> 4) + (8) = (3 -> 2).
func sumListsReverse(a *SinglyLinkedList, b *SinglyLinkedList) *SinglyLinkedList {
	return reverseList(sumLists(reverseList(a), reverseList(b)))
}

// reverseList returns a new linked list with the elements reversed.
func reverseList(l *SinglyLinkedList) *SinglyLinkedList {
	result := SinglyLinkedList{}
	result.head, _ = reverseListByNode(l.head)
	return &result
}

// reverseListByNode returns the head and tail of a new set of linked list nodes with the elements reversed.
func reverseListByNode(node *SinglyLinkedListNode) (*SinglyLinkedListNode, *SinglyLinkedListNode) {
	if node == nil {
		return nil, nil
	}
	curr := SinglyLinkedListNode{data: node.data}
	head, tail := reverseListByNode(node.next)
	if head == nil {
		return &curr, &curr
	}
	tail.next = &curr
	return head, &curr
}

func sumListsReverse(a *SinglyLinkedList, b *SinglyLinkedList) *SinglyLinkedList {
	return reverseList(sumLists(reverseList(a), reverseList(b)))
}

func Test2Dot5(t *testing.T) {
	a := SinglyLinkedList{}
	a.Add(7)
	a.Add(1)
	a.Add(6)
	b := SinglyLinkedList{}
	b.Add(5)
	b.Add(9)
	b.Add(2)
	actual := sumLists(&a, &b)
	assert.Equal(t, []int{2, 1, 9}, ToArray(*actual))
}

func Test2Dot5b(t *testing.T) {
	a := SinglyLinkedList{}
	a.Add(6)
	a.Add(1)
	a.Add(7)
	b := SinglyLinkedList{}
	b.Add(2)
	b.Add(9)
	b.Add(5)
	actual := sumListsReverse(&a, &b)
	assert.Equal(t, []int{9, 1, 2}, ToArray(*actual))
}
