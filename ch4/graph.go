package ch4

// A Graph is a general graph structure.
type Graph struct {
	nodes []*Node
}

// A Node is a node in a Graph.
type Node struct {
	name     string
	parents  []*Node
	children []*Node
	visited  bool
}

// AddChild adds the child node to the given node and graph;
func (node *Node) AddChild(child *Node, graph *Graph) {
	node.children = append(node.children, child)
	graph.nodes = append(graph.nodes, child)
}

// GenerateGraph1 produces sample graph 1 with two clusters.
func GenerateGraph1() *Graph {
	graph := &Graph{nodes: []*Node{}}

	// First cluster
	node0 := &Node{name: "0", children: []*Node{}}
	node1 := &Node{name: "1", children: []*Node{}}
	node2 := &Node{name: "2", children: []*Node{}}
	node3 := &Node{name: "3", children: []*Node{}}

	// Second cluster
	node4 := &Node{name: "4", children: []*Node{}}
	node5 := &Node{name: "5", children: []*Node{}}
	node6 := &Node{name: "6", children: []*Node{}}

	// First cluster
	node0.AddChild(node1, graph)
	node1.AddChild(node2, graph)
	node2.AddChild(node0, graph)
	node2.AddChild(node0, graph)
	node2.AddChild(node3, graph)
	node3.AddChild(node2, graph)

	// Second cluster
	node5.AddChild(node4, graph)
	node4.AddChild(node6, graph)
	node6.AddChild(node5, graph)
	return graph
}
