package generic

import (
	"fmt"
	"sort"
)

// Reference: https://codilime.com/blog/generics-in-go-definition-history-and-examples-of-use/
type FloatStringer interface {
	~float32 | ~float64 // union of types
	String() string
}

type floatString[T FloatType] struct {
	f T
}

func (f *floatString[T]) String() string {
	return fmt.Sprintf("%v", f.f)
}

type FloatLister interface {
	sort.Interface
	String() string
}

func (f *floatString[T]) Len() int { return len(f.String()) }
func (f *floatString[T]) Less(i, j int)
