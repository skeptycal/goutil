package sort

import generic "github.com/skeptycal/goutil/generic/generictypes"

// An implementation of the standard library sort.Interface
// that can be sorted by the routines in this package.
type Sorter[K generic.Ordered, V any] interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i
	// must sort before the element with index j.
	// More details can be found in the standard
	// library sort package documentation ...
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
