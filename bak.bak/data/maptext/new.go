package maptext

func NewTree(name string) *Tree {
	root := NewRoot(name, BlankData)
	return &Tree{
		// n:     1,
		// name:  name,
		root: root,
		last: root, // TODO: or nil initially??
		// nodes: blankNodeMap,
		// edges: blankEdgeMap,
	}
}

func NewNode(name string, data []byte, root, parent *node) *node {
	n := node{
		name:     name,
		data:     data,
		checksum: CheckSum(data),
		root:     root,
		parent:   parent,
		left:     nil,
		right:    nil,
	}
	return &n
}

func NewRoot(name string, data []byte) *node {
	n := node{
		name:     name,
		data:     data,
		checksum: CheckSum(data),
		root:     nil,
		parent:   nil,
		left:     nil,
		right:    nil,
	}
	n.root = &n
	return &n
}
