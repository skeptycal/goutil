package testes

import "reflect"

// Contains returns true if the underlying iterable
// sequence (haystack) contains the search term
// (needle) in at least one position.
func Contains(needle Any, haystack []Any) bool {
	for _, x := range haystack {
		if reflect.DeepEqual(needle, x) {
			return true
		}
	}
	return false
}
