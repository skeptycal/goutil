package generic

type (
	Comparable interface {
		comparable
	}

	// All Ordered types
	// string
	// int, int16, int32, int64, int8
	// uint, uint16, uint32, uint64, uint8
	// float32, float64
	// uintptr
	AllOrdered interface {
		Number | ~string | uintptr
	}

	Ordered interface {
		Number | ~string
	}

	Stringable interface {
		IntType | UintType
	}

	Number interface {
		IntType | UintType | FloatType
	}

	IntType interface {
		int | int8 | int16 | int32 | int64
	}

	UintType interface {
		uint | uint8 | uint16 | uint32 | uint64
	}

	FloatType interface {
		float32 | float64
	}
)

// var IsComparable = types.IsComparable
