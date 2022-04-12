package dict

import (
	"fmt"
	"sort"
	"strings"
)

type Lister[K comparable, V any] interface {
	Keys() []K
	Values() []V
}

type Mapper[K comparable, V any] interface {
	GetSetter[K, V]
	Stringer[K, V]
	Lookup(key K) (value V, ok bool)
}

type underlyingmap[K comparable, V any] map[K]V

type Dict[K comparable, V any] underlyingmap[K, V]

func (m underlyingmap[K, V]) String() string {
	sb := strings.Builder{}
	defer sb.Reset()
	for k, v := range m {
		sb.WriteString(fmt.Sprintf("%v: %v", k, v))
	}
	return sb.String()
}

func (m Dict[K, V]) String() string {
	sb := strings.Builder{}
	defer sb.Reset()
	sb.WriteString("Dict {\n")
	for k, v := range m {
		sb.WriteString(fmt.Sprintf("  %v: %v,\n", k, v))
	}
	sb.WriteString("}\n")
	return sb.String()
}

func (d Dict[K, V]) IsEmpty() bool {
	return d.Len() == 0
}

func (d Dict[K, V]) Len() int {
	return len(d)
}

func (d Dict[K, V]) Get(k K) (v V, ok bool) {
	if v, ok = d[k]; ok {
		return v, true
	}
	return v, false
}

func (d Dict[K, V]) Set(k K, v V) {
	d[k] = v
}

// func (d *Dict[K, V]) Enable() {
// 	d.disabled = true
// }

// func (d *Dict[K, V]) Disable() {
// 	d.disabled = false
// }

//////////////////// stuff

// func (d *Dict[K, V]) Swap(i, j K) {
// 	d.m[i], d.m[j] = d.m[j], d.m[i]
// }

// func (d *Dict[K, V]) Less(i, j K) bool {
// 	return d.m < d.m[j]
// }

// func (d *Dict[K, V]) InsertSorted(k K) {
// 	// i := sort.SearchStrings(ss, s)

// 	d.s = append(d.s, k)
// 	copy(d.s[i+1:], d.s[i:])
// 	d.s[i] = k
// }
