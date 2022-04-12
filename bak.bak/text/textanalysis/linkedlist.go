package textanalysis

type (
	LinkedList interface {

		// First returns the 'first' element of the list.
		First() Node

		// Last returns the 'last' element of the list.
		Last() Node

		// Next returns the 'next' element of the list.
		// If there is no 'current' element, it returns
		// First()
		//
		// If Next() == nil, it returns First() if 'cyclical'
		// is true.
		Next() Node

		// Previous returns the 'previous' element of the list.
		Prev() Node

		// Cyclical returns true if the list is cyclical.
		Cyclical() bool
	}

	// LinkedList is a linked list of nodes. These nodes
	// may have additional properties and functionality
	// by using graph and edge information.
	linkedList struct {
		first    Node
		last     Node
		current  Node
		cyclical bool
	}

	ListItem interface {
		Node
		Next() Node
		Prev() Node
	}

	listItem struct {
		node // Base struct

		// a linked list funcionality is handy for bulk
		// processing. Graph.nodes may also work fine.
		//
		// The *Nodes next and previous are not required
		// and the default values are nil unless linked lists
		// are activated as an option.
		next     Node
		previous Node
	}
)
