package dict

import "sort"

type GetSetter[K comparable, V any] interface {
	Get(key K) (value V)
	Set(key K, value V)
}

type Stringer[K comparable, V any] interface {
	String() string
}

type Sorter[K comparable, V any] interface {
	sort.Interface
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
