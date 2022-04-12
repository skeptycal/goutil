package textanalysis

type (
	Edges interface {
		First() Edge
	}

	Edge interface {

		// Start returns the 'first' node of the edge.
		Start() Node

		// Start returns the 'second' node of the edge.
		End() Node

		// Forward returns the true if the edge is
		// connected in the forward direction.
		Forward() bool

		// Backward returns the true if the edge is
		// connected in the backward direction.
		Backward() bool
	}

	EdgeList []edge

	edge struct {

		// Start is the 'first' node. This may or may
		// not matter, depending on use case.
		start Node

		// Start is the 'second' node. This may or may
		// not matter, depending on use case.
		end Node

		// directionality specifies the direction
		// of the edge between Start and End:
		// forward, backward, neither, or both.
		forwarddirection bool
		backdirection    bool
	}
)
