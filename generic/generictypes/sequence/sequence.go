package sequence

import "golang.org/x/exp/constraints"

// "golang.org/x/exp/constraints"

// func New[O Ordered]() []O {
// 	return Sequence[string]
// }

type (
	Ordered interface {
		constraints.Ordered
	}

	Sequence[O Ordered] interface {
		~[]O
	}

	Slice[E any] interface {
		~[]E
	}

	sequence struct {
	}
)

// First returns the first element in a sequence.
func First[T any](s []T) T {
	return s[0]
}

// Last returns the last element in a sequence.
func Last[T any](s []T) T {
	return s[len(s)-1]
}
