package orderedmap_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/elliotchance/orderedmap"
)

func benchmarkMap_Set(multiplier int) func(b *testing.B) {
	return func(b *testing.B) {
		m := make(map[int]bool)
		for i := 0; i < b.N*multiplier; i++ {
			m[i] = true
		}
	}
}

func BenchmarkMap_Set(b *testing.B) {
	benchmarkMap_Set(1)(b)
}

func benchmarkOrderedMap_Set(multiplier int) func(b *testing.B) {
	return func(b *testing.B) {
		m := orderedmap.NewOrderedMap()
		for i := 0; i < b.N*multiplier; i++ {
			m.Set(i, true)
		}
	}
}

func BenchmarkOrderedMap_Set(b *testing.B) {
	benchmarkOrderedMap_Set(1)(b)
}

func benchmarkMap_Get(multiplier int) func(b *testing.B) {
	m := make(map[int]bool)
	for i := 0; i < 1000*multiplier; i++ {
		m[i] = true
	}

	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = m[i%1000*multiplier]
		}
	}
}

func BenchmarkMap_Get(b *testing.B) {
	benchmarkMap_Get(1)(b)
}

func benchmarkOrderedMap_Get(multiplier int) func(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < 1000*multiplier; i++ {
		m.Set(i, true)
	}

	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Get(i % 1000 * multiplier)
		}
	}
}

func BenchmarkOrderedMap_Get(b *testing.B) {
	benchmarkOrderedMap_Get(1)(b)
}

func benchmarkOrderedMap_GetElement(multiplier int) func(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < 1000*multiplier; i++ {
		m.Set(i, true)
	}

	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.GetElement(i % 1000 * multiplier)
		}
	}
}

func BenchmarkOrderedMap_GetElement(b *testing.B) {
	benchmarkOrderedMap_GetElement(1)(b)
}

var tempInt int

func benchmarkOrderedMap_Len(multiplier int) func(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < 1000*multiplier; i++ {
		m.Set(i, true)
	}

	return func(b *testing.B) {
		var temp int
		for i := 0; i < b.N; i++ {
			temp = m.Len()
		}

		// prevent compiler from optimising Len away.
		tempInt = temp
	}
}

func BenchmarkOrderedMap_Len(b *testing.B) {
	benchmarkOrderedMap_Len(1)(b)
}

func benchmarkMap_Delete(multiplier int) func(b *testing.B) {
	return func(b *testing.B) {
		m := make(map[int]bool)
		for i := 0; i < b.N*multiplier; i++ {
			m[i] = true
		}

		for i := 0; i < b.N; i++ {
			delete(m, i)
		}
	}
}

func BenchmarkMap_Delete(b *testing.B) {
	benchmarkMap_Delete(1)(b)
}

func benchmarkOrderedMap_Delete(multiplier int) func(b *testing.B) {
	return func(b *testing.B) {
		m := orderedmap.NewOrderedMap()
		for i := 0; i < b.N*multiplier; i++ {
			m.Set(i, true)
		}

		for i := 0; i < b.N; i++ {
			m.Delete(i)
		}
	}
}

func BenchmarkOrderedMap_Delete(b *testing.B) {
	benchmarkOrderedMap_Delete(1)(b)
}

func benchmarkMap_Iterate(multiplier int) func(b *testing.B) {
	m := make(map[int]bool)
	for i := 0; i < 1000*multiplier; i++ {
		m[i] = true
	}
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range m {
				makeNothing(v)
			}
		}
	}
}
func BenchmarkMap_Iterate(b *testing.B) {
	benchmarkMap_Iterate(1)(b)
}

func benchmarkOrderedMap_Iterate(multiplier int) func(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < 1000*multiplier; i++ {
		m.Set(i, true)
	}

	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, key := range m.Keys() {
				_, v := m.Get(key)
				makeNothing(v)
			}
		}
	}
}

func BenchmarkOrderedMap_Iterate(b *testing.B) {
	benchmarkOrderedMap_Iterate(1)(b)
}

func benchmarkOrderedMap_Keys(multiplier int) func(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < 1000*multiplier; i++ {
		m.Set(i, true)
	}

	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Keys()
		}
	}
}

func benchmarkMapString_Set(multiplier int) func(b *testing.B) {
	return func(b *testing.B) {
		m := make(map[string]bool)
		a := "12345678"
		for i := 0; i < b.N*multiplier; i++ {
			m[a+strconv.Itoa(i)] = true
		}
	}
}

func BenchmarkMapString_Set(b *testing.B) {
	benchmarkMapString_Set(1)(b)
}

func benchmarkOrderedMapString_Set(multiplier int) func(b *testing.B) {
	return func(b *testing.B) {
		m := orderedmap.NewOrderedMap()
		a := "12345678"
		for i := 0; i < b.N*multiplier; i++ {
			m.Set(a+strconv.Itoa(i), true)
		}
	}
}

func BenchmarkOrderedMapString_Set(b *testing.B) {
	benchmarkOrderedMapString_Set(1)(b)
}

func benchmarkMapString_Get(multiplier int) func(b *testing.B) {
	m := make(map[string]bool)
	a := "12345678"
	for i := 0; i < 1000*multiplier; i++ {
		m[a+strconv.Itoa(i)] = true
	}

	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = m[a+strconv.Itoa(i%1000*multiplier)]
		}
	}
}

func BenchmarkMapString_Get(b *testing.B) {
	benchmarkMapString_Get(1)(b)
}

func benchmarkOrderedMapString_Get(multiplier int) func(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	a := "12345678"
	for i := 0; i < 1000*multiplier; i++ {
		m.Set(a+strconv.Itoa(i), true)
	}

	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Get(a + strconv.Itoa(i%1000*multiplier))
		}
	}
}

func BenchmarkOrderedMapString_Get(b *testing.B) {
	benchmarkOrderedMapString_Get(1)(b)
}

func benchmarkOrderedMapString_GetElement(multiplier int) func(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	a := "12345678"
	for i := 0; i < 1000*multiplier; i++ {
		m.Set(a+strconv.Itoa(i), true)
	}

	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.GetElement(a + strconv.Itoa(i%1000*multiplier))
		}
	}
}

func BenchmarkOrderedMapString_GetElement(b *testing.B) {
	benchmarkOrderedMapString_GetElement(1)(b)
}

func benchmarkMapString_Delete(multiplier int) func(b *testing.B) {
	return func(b *testing.B) {
		m := make(map[string]bool)
		a := "12345678"
		for i := 0; i < b.N*multiplier; i++ {
			m[a+strconv.Itoa(i)] = true
		}

		for i := 0; i < b.N; i++ {
			delete(m, a+strconv.Itoa(i))
		}
	}
}

func BenchmarkMapString_Delete(b *testing.B) {
	benchmarkMapString_Delete(1)(b)
}

func benchmarkOrderedMapString_Delete(multiplier int) func(b *testing.B) {
	return func(b *testing.B) {
		m := orderedmap.NewOrderedMap()
		a := "12345678"
		for i := 0; i < b.N*multiplier; i++ {
			m.Set(a+strconv.Itoa(i), true)
		}

		for i := 0; i < b.N; i++ {
			m.Delete(a + strconv.Itoa(i))
		}
	}
}

func BenchmarkOrderedMapString_Delete(b *testing.B) {
	benchmarkOrderedMapString_Delete(1)(b)
}

func benchmarkMapString_Iterate(multiplier int) func(b *testing.B) {
	m := make(map[string]bool)
	a := "12345678"
	for i := 0; i < 1000*multiplier; i++ {
		m[a+strconv.Itoa(i)] = true
	}
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range m {
				makeNothing(v)
			}
		}
	}
}
func BenchmarkMapString_Iterate(b *testing.B) {
	benchmarkMapString_Iterate(1)(b)
}

func benchmarkOrderedMapString_Iterate(multiplier int) func(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	a := "12345678"
	for i := 0; i < 1000*multiplier; i++ {
		m.Set(a+strconv.Itoa(i), true)
	}

	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, key := range m.Keys() {
				_, v := m.Get(key)
				makeNothing(v)
			}
		}
	}
}

func BenchmarkOrderedMapString_Iterate(b *testing.B) {
	benchmarkOrderedMapString_Iterate(1)(b)
}

func BenchmarkOrderedMap_Keys(b *testing.B) {
	benchmarkOrderedMap_Keys(1)(b)
}

func ExampleNewOrderedMap() {
	m := orderedmap.NewOrderedMap()

	m.Set("foo", "bar")
	m.Set("qux", 1.23)
	m.Set(123, true)

	m.Delete("qux")

	for _, key := range m.Keys() {
		value, _ := m.Get(key)
		fmt.Println(key, value)
	}
}

func ExampleOrderedMap_Front() {
	m := orderedmap.NewOrderedMap()
	m.Set(1, true)
	m.Set(2, true)

	for el := m.Front(); el != nil; el = el.Next() {
		fmt.Println(el)
	}
}

func makeNothing(v interface{}) {
	if v != nil {
		v = false
	}
}

func benchmarkBigMap_Set() func(b *testing.B) {
	return func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			m := make(map[int]bool)
			for i := 0; i < 10000000; i++ {
				m[i] = true
			}
		}
	}
}

func BenchmarkBigMap_Set(b *testing.B) {
	benchmarkBigMap_Set()(b)
}

func benchmarkBigOrderedMap_Set() func(b *testing.B) {
	return func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			m := orderedmap.NewOrderedMap()
			for i := 0; i < 10000000; i++ {
				m.Set(i, true)
			}
		}
	}
}

func BenchmarkBigOrderedMap_Set(b *testing.B) {
	benchmarkBigOrderedMap_Set()(b)
}

func benchmarkBigMap_Get() func(b *testing.B) {
	m := make(map[int]bool)
	for i := 0; i < 10000000; i++ {
		m[i] = true
	}

	return func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			for i := 0; i < 10000000; i++ {
				_ = m[i]
			}
		}
	}
}

func BenchmarkBigMap_Get(b *testing.B) {
	benchmarkBigMap_Get()(b)
}

func benchmarkBigOrderedMap_Get() func(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < 10000000; i++ {
		m.Set(i, true)
	}

	return func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			for i := 0; i < 10000000; i++ {
				m.Get(i)
			}
		}
	}
}

func BenchmarkBigOrderedMap_Get(b *testing.B) {
	benchmarkBigOrderedMap_Get()(b)
}

func benchmarkBigOrderedMap_GetElement() func(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < 10000000; i++ {
		m.Set(i, true)
	}

	return func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			for i := 0; i < 10000000; i++ {
				m.GetElement(i)
			}
		}
	}
}

func BenchmarkBigOrderedMap_GetElement(b *testing.B) {
	benchmarkBigOrderedMap_GetElement()(b)
}

func benchmarkBigMap_Iterate() func(b *testing.B) {
	m := make(map[int]bool)
	for i := 0; i < 10000000; i++ {
		m[i] = true
	}
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range m {
				makeNothing(v)
			}
		}
	}
}
func BenchmarkBigMap_Iterate(b *testing.B) {
	benchmarkBigMap_Iterate()(b)
}

func benchmarkBigOrderedMap_Iterate() func(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < 10000000; i++ {
		m.Set(i, true)
	}

	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, key := range m.Keys() {
				_, v := m.Get(key)
				makeNothing(v)
			}
		}
	}
}

func BenchmarkBigOrderedMap_Iterate(b *testing.B) {
	benchmarkBigOrderedMap_Iterate()(b)
}

func benchmarkBigMapString_Set() func(b *testing.B) {
	return func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			m := make(map[string]bool)
			a := "1234567"
			for i := 0; i < 10000000; i++ {
				m[a+strconv.Itoa(i)] = true
			}
		}
	}
}

func BenchmarkBigMapString_Set(b *testing.B) {
	benchmarkBigMapString_Set()(b)
}

func benchmarkBigOrderedMapString_Set() func(b *testing.B) {
	return func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			m := orderedmap.NewOrderedMap()
			a := "1234567"
			for i := 0; i < 10000000; i++ {
				m.Set(a+strconv.Itoa(i), true)
			}
		}
	}
}

func BenchmarkBigOrderedMapString_Set(b *testing.B) {
	benchmarkBigOrderedMapString_Set()(b)
}

func benchmarkBigMapString_Get() func(b *testing.B) {
	m := make(map[string]bool)
	a := "1234567"
	for i := 0; i < 10000000; i++ {
		m[a+strconv.Itoa(i)] = true
	}

	return func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			for i := 0; i < 10000000; i++ {
				_ = m[a+strconv.Itoa(i)]
			}
		}
	}
}

func BenchmarkBigMapString_Get(b *testing.B) {
	benchmarkBigMapString_Get()(b)
}

func benchmarkBigOrderedMapString_Get() func(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	a := "1234567"
	for i := 0; i < 10000000; i++ {
		m.Set(a+strconv.Itoa(i), true)
	}

	return func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			for i := 0; i < 10000000; i++ {
				m.Get(a + strconv.Itoa(i))
			}
		}
	}
}

func BenchmarkBigOrderedMapString_Get(b *testing.B) {
	benchmarkBigOrderedMapString_Get()(b)
}

func benchmarkBigOrderedMapString_GetElement() func(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	a := "1234567"
	for i := 0; i < 10000000; i++ {
		m.Set(a+strconv.Itoa(i), true)
	}

	return func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			for i := 0; i < 10000000; i++ {
				m.GetElement(a + strconv.Itoa(i))
			}
		}
	}
}

func BenchmarkBigOrderedMapString_GetElement(b *testing.B) {
	benchmarkBigOrderedMapString_GetElement()(b)
}

func benchmarkBigMapString_Iterate() func(b *testing.B) {
	m := make(map[string]bool)
	a := "12345678"
	for i := 0; i < 10000000; i++ {
		m[a+strconv.Itoa(i)] = true
	}
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range m {
				makeNothing(v)
			}
		}
	}
}
func BenchmarkBigMapString_Iterate(b *testing.B) {
	benchmarkBigMapString_Iterate()(b)
}

func benchmarkBigOrderedMapString_Iterate() func(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	a := "12345678"
	for i := 0; i < 10000000; i++ {
		m.Set(a+strconv.Itoa(i), true)
	}

	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, key := range m.Keys() {
				_, v := m.Get(key)
				makeNothing(v)
			}
		}
	}
}

func BenchmarkBigOrderedMapString_Iterate(b *testing.B) {
	benchmarkBigOrderedMapString_Iterate()(b)
}

func BenchmarkAll(b *testing.B) {
	b.Run("BenchmarkOrderedMap_Keys", BenchmarkOrderedMap_Keys)

	b.Run("BenchmarkOrderedMap_Set", BenchmarkOrderedMap_Set)
	b.Run("BenchmarkMap_Set", BenchmarkMap_Set)
	b.Run("BenchmarkOrderedMap_Get", BenchmarkOrderedMap_Get)
	b.Run("BenchmarkMap_Get", BenchmarkMap_Get)
	b.Run("BenchmarkOrderedMap_GetElement", BenchmarkOrderedMap_GetElement)
	b.Run("BenchmarkOrderedMap_Delete", BenchmarkOrderedMap_Delete)
	b.Run("BenchmarkMap_Delete", BenchmarkMap_Delete)
	b.Run("BenchmarkOrderedMap_Iterate", BenchmarkOrderedMap_Iterate)
	b.Run("BenchmarkMap_Iterate", BenchmarkMap_Iterate)

	b.Run("BenchmarkBigMap_Set", BenchmarkBigMap_Set)
	b.Run("BenchmarkBigOrderedMap_Set", BenchmarkBigOrderedMap_Set)
	b.Run("BenchmarkBigMap_Get", BenchmarkBigMap_Get)
	b.Run("BenchmarkBigOrderedMap_Get", BenchmarkBigOrderedMap_Get)
	b.Run("BenchmarkBigOrderedMap_GetElement", BenchmarkBigOrderedMap_GetElement)
	b.Run("BenchmarkBigOrderedMap_Iterate", BenchmarkBigOrderedMap_Iterate)
	b.Run("BenchmarkBigMap_Iterate", BenchmarkBigMap_Iterate)

	b.Run("BenchmarkOrderedMapString_Set", BenchmarkOrderedMapString_Set)
	b.Run("BenchmarkMapString_Set", BenchmarkMapString_Set)
	b.Run("BenchmarkOrderedMapString_Get", BenchmarkOrderedMapString_Get)
	b.Run("BenchmarkMapString_Get", BenchmarkMapString_Get)
	b.Run("BenchmarkOrderedMapString_GetElement", BenchmarkOrderedMapString_GetElement)
	b.Run("BenchmarkOrderedMapString_Delete", BenchmarkOrderedMapString_Delete)
	b.Run("BenchmarkMapString_Delete", BenchmarkMapString_Delete)
	b.Run("BenchmarkOrderedMapString_Iterate", BenchmarkOrderedMapString_Iterate)
	b.Run("BenchmarkMapString_Iterate", BenchmarkMapString_Iterate)

	b.Run("BenchmarkBigMapString_Set", BenchmarkBigMapString_Set)
	b.Run("BenchmarkBigOrderedMapString_Set", BenchmarkBigOrderedMapString_Set)
	b.Run("BenchmarkBigMapString_Get", BenchmarkBigMapString_Get)
	b.Run("BenchmarkBigOrderedMapString_Get", BenchmarkBigOrderedMapString_Get)
	b.Run("BenchmarkBigOrderedMapString_GetElement", BenchmarkBigOrderedMapString_GetElement)
	b.Run("BenchmarkBigOrderedMapString_Iterate", BenchmarkBigOrderedMapString_Iterate)
	b.Run("BenchmarkBigMapString_Iterate", BenchmarkBigMapString_Iterate)
}
