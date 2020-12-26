package ch4

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

// BST is a binary search tree. Assumes that no two nodes have the same value.
type BST struct {
	root *BinaryTreeNode
}

func (b *BST) Delete(data int) *BinaryTreeNode {
	node := b.Find(data)
	if node == nil {
		return nil
	}
	if node.left == nil && node.right == nil {
		// Case 1: It is a leaf node.
		if node.parent.left == node {
			node.parent.left = nil
		}
		if node.parent.right == node {
			node.parent.right = nil
		}
		return node
	}
	if node.left != nil && node.right == nil {
		// Case 2a: It is a parent of one child.
		if node.parent.left == node {
			node.parent.left = node.left
		}
		if node.parent.right == node {
			node.parent.right = node.left
		}
		return node
	}
	if node.left == nil && node.right != nil {
		// Case 2b: It is a parent of one child.
		if node.parent.left == node {
			node.parent.left = node.right
		}
		if node.parent.right == node {
			node.parent.right = node.right
		}
		return node
	}
	// It is a parent of two children. Find successor, delete it, and
	// relabel node as successor. The successor is the leftmost node of the
	// right subtree.
	curr := node.right
	for true {
		if curr.left == nil {
			break
		}
		curr = curr.left
	}
	// curr is either case 1 (a leaf node) or case 2 (a parent of one child).
	// Delete it then relabel node using its data.
	b.Delete(curr.data)
	node.data = curr.data
	return node
}

func (b *BST) Find(data int) *BinaryTreeNode {
	if b.root == nil {
		return nil
	}
	curr := b.root
	for true {
		if data == curr.data {
			return curr
		}
		if data < curr.data {
			if curr.left == nil {
				return nil
			}
			curr = curr.left
			continue
		}
		if curr.right == nil {
			return nil
		}
		curr = curr.right
	}
	panic("Shouldn't get here")
}

func (b *BST) Insert(data int) {
	if b.root == nil {
		b.root = &BinaryTreeNode{data: data}
		return
	}
	curr := b.root
	for true {
		if data == curr.data {
			panic("did not expect data to already exist in the tree")
		}
		if data < curr.data {
			if curr.left == nil {
				curr.left = &BinaryTreeNode{data: data}
				curr.left.parent = curr
				return
			}
			curr = curr.left
			continue
		}
		if curr.right == nil {
			curr.right = &BinaryTreeNode{data: data}
			curr.right.parent = curr
			return
		}
		curr = curr.right
	}
}

func (b *BST) GetRandomNode() *BinaryTreeNode {
	if b.root == nil {
		return nil
	}
	// Use BFS to get all nodes
	nodes := []*BinaryTreeNode{}
	queue := []*BinaryTreeNode{b.root}
	for len(queue) > 0 {
		head := queue[0]
		nodes = append(nodes, head)
		if head.left != nil {
			queue = append(queue, head.left)
		}
		if head.right != nil {
			queue = append(queue, head.right)
		}
		queue = queue[1:]
	}
	// Pick a random element from the list.
	i := rand.Intn(len(nodes))
	return nodes[i]
}

func createBST() *BST {
	bst := BST{}
	bst.Insert(4)
	bst.Insert(2)
	bst.Insert(6)
	bst.Insert(1)
	bst.Insert(7)
	bst.Insert(3)
	return &bst
}

func TestBSTFind(t *testing.T) {
	bst := createBST()
	assert.Equal(t, 4, bst.Find(4).data)
	assert.Equal(t, 2, bst.Find(2).data)
	assert.Equal(t, 6, bst.Find(6).data)
	assert.Equal(t, 1, bst.Find(1).data)
	assert.Nil(t, bst.Find(8))
	assert.NotNil(t, bst.GetRandomNode())
}

func TestBSTGetRandomNode(t *testing.T) {
	bst := createBST()
	assert.NotNil(t, bst.GetRandomNode())
}

func TestBSTDeleteLeaf(t *testing.T) {
	bst := createBST()
	bst.Delete(1)
	assert.Equal(t, 4, bst.root.data)
	assert.Equal(t, 2, bst.root.left.data)
	assert.Equal(t, 6, bst.root.right.data)
	assert.Nil(t, bst.root.left.left)
	assert.Equal(t, 3, bst.root.left.right.data)
	assert.Equal(t, 7, bst.root.right.right.data)
}

func TestBSTDeleteParentOfOneChild(t *testing.T) {
	bst := createBST()
	bst.Delete(6)
	assert.Equal(t, 4, bst.root.data)
	assert.Equal(t, 2, bst.root.left.data)
	assert.Equal(t, 7, bst.root.right.data)
	assert.Equal(t, 1, bst.root.left.left.data)
	assert.Equal(t, 3, bst.root.left.right.data)
}

func TestBSTDeleteParentOfTwoChildren(t *testing.T) {
	bst := createBST()
	bst.Delete(4)
	assert.Equal(t, 2, bst.root.left.data)
	assert.Equal(t, 6, bst.root.data)
	assert.Equal(t, 1, bst.root.left.left.data)
	assert.Equal(t, 3, bst.root.left.right.data)
	assert.Equal(t, 7, bst.root.right.data)
}
