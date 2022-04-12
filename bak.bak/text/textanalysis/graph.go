package textanalysis

type (
	Grapher interface {
		First() Node
		Nodes() []Node
		Edges() EdgeList
	}

	Graph struct {
		first Node
		// nodes is a numbered list of nodes that closely
		// corresponds to a SQL table with a primary key.
		nodes map[int]Node
	}
)
