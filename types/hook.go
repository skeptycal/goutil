package types

type (
	Node interface{}

	hook = node
	node struct {
		// id     *node
		left   *node
		right  *node
		parent *node

		name     string
		path     string
		lineno   int
		fn       Any
		funcName string
		args     []*nodeArg
	}

	nodeArg struct {
		parentID *node
		name     string
	}
)

// BuildNodes returns the parent node of a tree structure.
func BuildNodes() Node {
	return &node{name: "root", parent: nil}
}

func (n *node) AddLeftChild(c *node) {
	n.left = c
	c.parent = n
}

func (n *node) AddRightChild(c *node) {
	if c == nil {
		c = &node{parent: n}
	}
	n.right = c
	c.parent = n
}
