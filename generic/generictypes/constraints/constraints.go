package constraints

/// User Defined types ... some of the discussion:

// Lesser defines a type-parameterized interface[1] with one method, Less, which returns a boolean for whether the caller, of type T, is less than some other instance of T. This is blatantly stolen from Robert Griesemer's talk at Gophercon 2020[2] about the type parameters proposal. Probably more controversially, this library also defines a wrapper called Basic over the built-in numerical types, exposing the underlying < operator through this Less method. The reasoning for this follows.
//
// 	// Interface is an interface that wraps the Less method.
// 	//
// 	// Less compares a caller of type T to some other variable of type T,
// 	// returning true if the caller is the lesser of the two values, and false
// 	// otherwise. If the two values are equal it returns false.
// 	type Interface[T any] interface {
// 		Less(other T) bool
// 	}
//
//	// Basic is a parameterized type that abstracts over the entire class of
//	// Ordered types (the set of Go built-in types which respond to the <
//	// operator), and exposes this behavior via a Less method so that they
//	// fall under the lesser.Interface constraint.
//	type Basic[N constraints.Ordered] struct{ Val N }
//
// 	// Less implements Interface[Basic[N]] for Basic[N]. Returns true if the value
// 	// of the caller is less than that of the parameter; otherwise returns
// 	// false.
// 	func (x Basic[N]) Less(y Basic[N]) bool {
// 		return x.Val < y.Val
// 	}
//
// Ian Lance Taylor made (what I think is) a great suggestion here, arguing for constraining containers to any and passing a comparison function to the constructor. This is a good idea and subjectively feels more idiomatic than what I have done here. [3]
//
// [1]: https://github.com/lelysses/lesser
//
// [2]: https://www.youtube.com/watch?v=TborQFPY2IM&ab_channel=GopherAcademy
//
// [3]: https://github.com/golang/go/issues/47632#issuecomment-897168431
type Lesser[E any] interface {
	Less(E) bool
}

// Type constraints (mostly) from the constraints package.
type (

	// comparable is an interface that is implemented by all comparable types
	// (booleans, numbers, strings, pointers, channels, arrays of comparable types,
	// structs whose fields are all comparable types).
	// The comparable interface may only be used as a type parameter constraint,
	// not as the type of a variable.
	Comparable interface{ comparable }

	// Sortable interface {
	// 	Ordered | UserOrdered
	// }

	UserOrdered[T any] interface {
		Len() int
	}

	// Ordered is a constraint that permits any ordered type: any type
	// that supports the operators < <= >= >.
	Ordered interface{ Number | ~string }

	// Number is a constraint that permits any real number type.
	Number interface{ Integer | Float }

	// Integer is a constraint that permits any integer type.
	Integer interface{ Signed | Unsigned }

	// Complex is a constraint that permits any complex numeric type.
	Complex interface{ ~complex64 | ~complex128 }

	// Float is a constraint that permits any floating-point type.
	Float interface{ ~float64 | ~float32 }

	// Signed is a constraint that permits any signed integer type.
	Signed interface {
		// rune is an alias of int32
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}

	// Unsigned is a constraint that permits any unsigned integer type.
	Unsigned interface {
		// byte is an alias of uint8
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}
)
