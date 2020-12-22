package ch4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// isConnected returns whether two nodes are connected.
func isConnected(graph *Graph, a *Node, b *Node) bool {
	for _, node := range graph.nodes {
		node.visited = false
	}
	return isConnectedNode(a, b)
}

// isConnectedNode returns whether two nodes are connected.
func isConnectedNode(a *Node, b *Node) bool {
	// DFS
	a.visited = true
	if a == b {
		return true
	}
	for _, node := range a.children {
		if node.visited {
			continue
		}
		if isConnectedNode(node, b) {
			return true
		}
	}
	return false
}

func getNodeNamed(name string, graph *Graph) *Node {
	for _, node := range graph.nodes {
		if node.name == name {
			return node
		}
	}
	return nil
}

func Test4Dot1(t *testing.T) {
	graph := GenerateGraph1()
	assert.True(t, isConnected(graph, getNodeNamed("1", graph), getNodeNamed("3", graph)))
	assert.False(t, isConnected(graph, getNodeNamed("4", graph), getNodeNamed("3", graph)))
}
