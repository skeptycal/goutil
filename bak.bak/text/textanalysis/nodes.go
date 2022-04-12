package textanalysis

type (
	Nodes []Node
	Node  interface {
		Data() Any      // returns the data contained in the Node
		Graph() Grapher // returns the graph that this node is a member of.
		// Edges() EdgeList
	}

	node struct {

		// Data represents the information stored in this node
		Data interface{}

		// Edges is a slice of edges containing this Node
		// If edges are not enabled in this node set, this
		// slice should be empty.
		//
		// This is an inefficient way to store information
		// that can be produced by filtering the main list
		// of edges.
		//
		// ... but
		//
		// It greatly reduces the amount of time it takes
		// to process Node relationships. (It is basically
		// a cache of edges and should only be produced upon
		// request ... JIT generation.)
		//
		// Depending on the use case, you may want to use
		// the method Node.Edges() instead as this will
		// return the slice for processing but will not
		// store the information.
		Edges EdgeList
	}
)
