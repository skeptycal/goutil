package maptext

var (
	blankNodeMap = make(map[id]*node, smallNodeMapSize)
	blankEdgeMap = make(map[id]*edge, smallEdgeMapSize)
)

type (
	// Tree represents a Binary Tree structure
	//
	/// TODO: rules:
	// - last == parent means size is 1 (root only)
	// - could last == nil work better?
	//
	// - name is stored in Tree and parent ...
	Tree struct {
		root *node // root is the root node of the tree
		last *node // last is the last node in the tree

		// n    int    // n is the number of nodes
		// name string // name is the name of the node set

		// nodes NodeMap // map of nodes by index id
		// edges EdgeMap // map of edges by index id
	}

	NodeMap map[id]*node // map of nodes by index id
	EdgeMap map[id]*edge // map of edges by index id
)

func (t *Tree) String() string {
	return "Tree"
}

// makeTree is a testing utility
//	 nodeMap struct {
//	 	n     int          // n is the number of nodes
//	 	name  string       // name is the name of the node set
//	 	root  *node        // root is the root node of the tree
//	 	last  *node        // last is the last node in the tree
//	 	nodes map[id]*node // map of nodes by index id
//	 	edges map[id]*edge // map of edges by index id
//	 }
func makeTree(name string) *Tree {
	root := NewRoot(name, BlankData)
	nm := Tree{
		// name: name,
		root: root,
		last: root,
	}
	// nm.makeNodeMap()
	// nm.makeEdgeMap()
	return &nm
}

// func (t *Tree) makeNodeMap() {
// 	if t.nodes == nil {
// 		// 	make(map[id]*node, smallNodeMapSize)
// 		t.nodes = blankNodeMap
// 	}
// }

// func (t *Tree) makeEdgeMap() {
// 	if t.edges == nil {
// 		// make(map[id]*edge, smallEdgeMapSize)
// 		t.edges = blankEdgeMap
// 	}
// }
