// Package sort provides a generic implementation  of the
// standard library sort package primitives for sorting
// slices and user-defined collections.
//
// It should be noted that even with generics, some
// limitations remain on operations such as sorting
// based on data type, e.g.
// While any Comparable type can be uniquely identified,
// used as map keys, or compared with operators, only Ordered
// types can be sorted.
//
// Of the Go built-in types, the only Ordered types are
// ints, uints, uintptr, floats, and strings. Including
// user defined types in the sorting constraint infers
// that they must be ordered in some way. A user-defined
// interface is provided that must be implemented by any
// user-defined type in order for it to implement the
// algorithms in the sort package.
//
// The basic interface implemented is:
//  type Sorter interface {
// 		// Len is the number of elements in the collection.
// 		Len() int
//
// 		// Less reports whether the element with index i
// 		// must sort before the element with index j.
// 		// More details can be found in the standard
// 		// library sort package documentation ...
// 		Less(i, j int) bool
//
// 		// Swap swaps the elements with indexes i and j.
// 		Swap(i, j int)
// 	}
package sort
