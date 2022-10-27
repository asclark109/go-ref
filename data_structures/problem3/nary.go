package problem3

import (
	"strings"
)

// You cannot modify the name of this struct
type Node struct {
	// You may add as many fields to this struct as you wish
	children []*Node // NEW e.g. [&node1 &node2]
	nodeStr  string  // NEW e.g. "1"
	parent   *Node   // NEW e.g. "nil" for root Node
}

// addChild method adds a Node as a child of this Node
func (n *Node) addChild(node *Node) {
	n.children = append(n.children, node)
}

// numChildren returns the number of children of this Node
func (n *Node) numChildren() int {
	return len(n.children)
}

// newNode constructs a new Node struct and returns a pointer to it
func newNode(nodeStr string) *Node {
	return &Node{make([]*Node, 0), nodeStr, nil}
}

// You cannot modify this data structure at all (including its types). However
// You are not required to use all of its fields.
type Tree struct {
	root  *Node
	nAry  int
	nodes map[string]*Node
}

// contains returns true if this tree has a Node corresponding to the string name provided.
// All Nodes have a nodeStr, which is essentially a name. e.g. "1" -> true if tree
// has a Node with name "1".
func (t *Tree) contains(nodeStr string) bool {
	_, ok := t.nodes[nodeStr]
	return ok
}

// insertEdge accepts an edge as two strings (name of parent, name of child), and
// adds the edge to the tree. The first edge that gets added to tree has the following
// side effect: the parent becomes the root of the tree.
func (t *Tree) insertEdge(parentStr, childStr string) {
	// HIGH LEVEL ALGORITHM:
	// when adding edge (u,v) where u,v are strings:
	// (1) if u not in Tree
	//        create and insert new Node U into T
	//        if this is the first Node inserted, set this Node to be root of T
	// (2) if V not in Tree
	//        create and insert new Node V into T
	// (3) update parent of node V to be U
	//     add node V as a child of node U

	// (1) + (2)
	if !t.contains(parentStr) {
		newNode := newNode(parentStr)
		if len(t.nodes) == 0 { // set node to be root
			t.root = newNode
		}
		t.nodes[parentStr] = newNode
	}
	if !t.contains(childStr) {
		newNode := newNode(childStr)
		t.nodes[childStr] = newNode
	}
	// (3)
	parent := t.nodes[parentStr]
	child := t.nodes[childStr]
	parent.addChild(child)
	child.parent = parent
}

// numNodesExceedingArity returns the number of Nodes in this tree exceeding tree arity.
func (t *Tree) numNodesExceedingArity() int {
	count := 0
	for _, node := range t.nodes {
		if node.numChildren() > t.nAry {
			count++
		}
	}
	return count
}

// newTree constructs a new Tree struct and returns a pointer to it
func newTree(nAry int) *Tree {
	tree := Tree{root: nil, nAry: nAry, nodes: make(map[string]*Node, 0)}
	return &tree
}

// SearchTree takes in a tree and returns the number of internal nodes that have children that are greater
// than the arty of the tree, which is given by the second argument to the function.
// Note: Please make sure to read the assignment description for more details.
func SearchTree(edges []string, nAry int) int {
	// HIGH LEVEL ALGORITHM
	// (1) create new Tree struct. Add all edges to Tree.
	// (2) ask Tree to internally count the Nodes that exceed its arity

	// (1)
	tree := newTree(nAry)
	for _, edge := range edges {
		parentStr, childStr := interpretEdgeEntry(edge)
		tree.insertEdge(parentStr, childStr)
	}
	// (2)
	return tree.numNodesExceedingArity()
}

// interpretEdgeEntry accepts a string of format "**<parent>**<child>**",
// where ** represents any amount of whitespace, <parent>, <child> are any strings,
// and returns the strings <parent>, <child>
func interpretEdgeEntry(edgeStr string) (string, string) {
	strElems := strings.Fields(edgeStr)
	parentStr := strings.TrimSpace(strElems[0]) // e.g. "1"
	childStr := strings.TrimSpace(strElems[1])  // e.g. "2"
	return parentStr, childStr
}
