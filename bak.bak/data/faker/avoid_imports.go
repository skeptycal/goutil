package faker

/* The following types, constants, and functions are placed here to avoid importing from other utilities.

Because this is one of several low level utility packages, they often rely on each other. This can create circular import issues.
*/
import "github.com/skeptycal/goutil/types"

type (
	Any = interface{}

	AnyValue = types.AnyValue
)

var (

	// NewAnyValue returns a new AnyValue interface, e.g. v := NewAnyValue(uint(42))
	//
	// AnyValue is a wrapper around the Any interface, or interface{}, which may contain any value. The original interface{} value is returned by v.Interface()
	//
	// The extra features of this wrapper allow value, type, and kind information, as well as whether the type is comparable, ordered, and/or iterable.
	NewAnyValue = types.NewAnyValue

	// Contains returns true if the underlying iterable
	// sequence (haystack) contains the search term
	// (needle) in at least one position.
	Contains = types.Contains
)

// // Contains returns true if the underlying iterable
// // sequence (haystack) contains the search term
// // (needle) in at least one position.
// func Contains(needle Any, haystack []Any) bool {
// 	for _, x := range haystack {
// 		if reflect.DeepEqual(needle, x) {
// 			return true
// 		}
// 	}
// 	return false
// }
