package ch4

import (
	"testing"

	"github.com/JonathanAquino/cracking/ch2"
	"github.com/stretchr/testify/assert"
)

// listOfDepths takes a binary tree and returns a list of linked lists, each with
// the tree nodes at a given depth.
func listOfDepths(node *BinaryTreeNode) []*ch2.SinglyLinkedList {
	if node == nil {
		return []*ch2.SinglyLinkedList{}
	}
	rootList := &ch2.SinglyLinkedList{}
	rootList.Add(node.data)
	lists := []*ch2.SinglyLinkedList{rootList}
	leftLists := listOfDepths(node.left)
	rightLists := listOfDepths(node.right)
	for i := 0; i < len(leftLists) || i < len(rightLists); i++ {
		list := &ch2.SinglyLinkedList{}
		if i < len(leftLists) {
			curr := leftLists[i].Head
			for curr != nil {
				list.Add(curr.Data)
				curr = curr.Next
			}
		}
		if i < len(rightLists) {
			curr := rightLists[i].Head
			for curr != nil {
				list.Add(curr.Data)
				curr = curr.Next
			}
		}
		lists = append(lists, list)
	}
	return lists
}

func Test4Dot3(t *testing.T) {
	// Level 3
	node1 := &BinaryTreeNode{data: 1}
	node4 := &BinaryTreeNode{data: 4}
	node7 := &BinaryTreeNode{data: 7}

	// Level 2
	node2 := &BinaryTreeNode{data: 2, left: node1}
	node5 := &BinaryTreeNode{data: 5, left: node4}
	node8 := &BinaryTreeNode{data: 8, left: node7}
	node10 := &BinaryTreeNode{data: 10}

	// Level 1
	node3 := &BinaryTreeNode{data: 3, left: node2, right: node5}
	node9 := &BinaryTreeNode{data: 9, left: node8, right: node10}

	// Level 0
	node6 := &BinaryTreeNode{data: 6, left: node3, right: node9}

	linkedLists := listOfDepths(node6)
	assert.Equal(t, 6, linkedLists[0].Head.Data)
	assert.Equal(t, 3, linkedLists[1].Head.Data)
	assert.Equal(t, 9, linkedLists[1].Head.Next.Data)
	assert.Equal(t, 2, linkedLists[2].Head.Data)
	assert.Equal(t, 5, linkedLists[2].Head.Next.Data)
	assert.Equal(t, 8, linkedLists[2].Head.Next.Next.Data)
	assert.Equal(t, 10, linkedLists[2].Head.Next.Next.Next.Data)
	assert.Equal(t, 1, linkedLists[3].Head.Data)
	assert.Equal(t, 4, linkedLists[3].Head.Next.Data)
	assert.Equal(t, 7, linkedLists[3].Head.Next.Next.Data)
}
