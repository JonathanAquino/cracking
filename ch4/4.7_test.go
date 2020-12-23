package ch4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// buildOrder returns the projects in an order that satisfies the given dependencies.
// Assumes that there are no cycles.
func buildOrder(projects []string, dependencies []Pair) []string {
	// Build the graph.
	graph := Graph{nodes: []*Node{}}
	projectNodes := make(map[string]*Node)
	for _, project := range projects {
		node := Node{name: project, parents: []*Node{}, children: []*Node{}}
		projectNodes[project] = &node
		graph.nodes = append(graph.nodes, &node)
	}
	for _, dependency := range dependencies {
		// b depends on a. So we make a a parent of b.
		projectNodes[dependency.b].parents = append(projectNodes[dependency.b].parents, projectNodes[dependency.a])
		projectNodes[dependency.a].children = append(projectNodes[dependency.a].children, projectNodes[dependency.b])
	}
	// Find the starting points (those that do not have any dependencies).
	startingPoints := []*Node{}
	for _, node := range graph.nodes {
		if len(node.parents) == 0 {
			startingPoints = append(startingPoints, node)
		}
	}
	// Radiate outwards using BFS.
	// Warning: startingPoints may be modified by these operations.
	queue := startingPoints
	buildOrder := []*Node{}
	for len(queue) != 0 {
		if !queue[0].visited {
			buildOrder = append(buildOrder, queue[0])
			queue[0].visited = true
			queue = append(queue, queue[0].children...)
		}
		queue = queue[1:]
	}
	result := []string{}
	for _, node := range buildOrder {
		result = append(result, node.name)
	}
	return result
}

type Pair struct {
	a, b string
}

func TestBuildOrder(t *testing.T) {
	buildOrder := buildOrder(
		[]string{"a", "b", "c", "d", "e", "f"},
		[]Pair{Pair{"a", "d"}, Pair{"f", "b"}, Pair{"b", "d"}, Pair{"f", "a"}, Pair{"d", "c"}},
	)
	assert.Equal(t, []string{"e", "f", "b", "a", "d", "c"}, buildOrder)
}
