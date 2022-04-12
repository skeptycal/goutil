package maptext

import "fmt"

const (
	smallNodeMapSize int = 8
	smallEdgeMapSize int = 16
)

type (
	node struct {
		name     string
		data     []byte
		checksum [16]byte

		// Connecting nodes
		root   *node
		parent *node
		left   *node
		right  *node
	}
)

func (n *node) String() string {
	return fmt.Sprintf("%s: %s", n.name, n.data)
}

func (n *node) Len() int {
	if n.parent == nil {
		return len(n.data)
	}
	return n.root.Len()
}
func (n *node) Swap(i, j int)          { n.data, n.name = n.parent.data, n.parent.name }
func (n *node) Less(i, j int) bool     { return len(n.data) < len(n.parent.data) }
func (n *node) Equal(other *node) bool { return string(n.data) == string(other.data) }
func (n *node) CheckSum() [16]byte     { return CheckSum(n.data) }
func (n *node) AddLeft(other *node) {
	n.left = other
	other.parent = n
}
func (n *node) AddRight(other *node) {
	n.right = other
	other.parent = n
}

func (n *node) Delete() {}
