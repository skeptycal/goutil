package benchmark

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
	"unsafe"
)

type (
	Stringer  interface{ String() string }
	_stringer struct{ int }
)

func (s _stringer) String() string { return fmt.Sprintf("%d", s) }
func NewStringer() Stringer        { return &_stringer{rand.Intn(255)} }
func init()                        { rand.Seed(time.Now().Unix()) }
func randomKind() reflect.Kind     { return reflect.Kind(rand.Intn(26)) }
func randBool() bool               { return rand.Int63() >= halfHalfBool }

const (
	b8  = 1<<8 - 1
	b16 = 1<<16 - 1
	b32 = 1<<32 - 1
	b64 = 1<<64 - 1
	u8  = b64 >> 56
	u16 = b64 >> 48
)

func randomData() Any {
	switch k := randomKind(); k {
	case 0: // Invalid
		return k
	case 1: // Bool
		return randBool()
	case 2: // Int
		return rand.Int()
	case 3: // Int8
		return int8(rand.Intn(b8))
	case 4: // Int16
		return int16(rand.Intn(b16))
	case 5: // Int32
		return int32(rand.Uint32())
	case 6: // Int64
		return int64(rand.Uint64())
	case 7: // Uint
		return uint(rand.Uint64())
	case 8: // Uint8
		return rand.Uint64() >> 56
	case 9: // Uint16
		return rand.Uint64() >> 48
	case 10: // Uint32
		return rand.Uint32()
	case 11: // Uint64
		return rand.Uint64()
	case 12: // Uintptr
	case 13: // Float32
		return rand.Float32()
	case 14: // Float64
		return rand.Float64()
	case 15: // Complex64
		var r = rand.Float32()
		var c = rand.Float32()
		return complex(r, c)
	case 16: // Complex128
		var r = rand.Float64()
		var c = rand.Float64()
		return complex(r, c)
	case 17: // Array
		return [42]bool{}
	case 18: // Chan
		return make(chan int, rand.Intn(2))
	case 19: // Func
		return func(x interface{}) interface{} { return "stuff" }
	case 20: // Interface
		return NewStringer()
	case 21: // Map
		return make(map[string]interface{}, rand.Intn(250)+5)
	case 22: // Ptr
		var i int = rand.Intn(42)
		return &i
	case 23: // Slice
		return make([]interface{}, rand.Intn(250)+5)
	case 24: // String
		return RandomString(42)
	case 25: // Struct
		return _stringer{rand.Intn(42)}
	case 26: // UnsafePointer
		arr := []uint32{1, 2, 3}
		const size = unsafe.Sizeof(uint32(0))
		p := uintptr(unsafe.Pointer(&arr[0]))
		return (*(*uint32)(unsafe.Pointer(p)))
	}
	return nil
}
