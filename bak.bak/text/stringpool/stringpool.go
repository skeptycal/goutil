// Package stringpool provides a sync.Pool of efficient
// strings.Builder workers that may be reused as needed,
// reducing the need to instantiate and allocate new
// builders in text heavy applications.
//
// From the Go standard library:
//
// A Builder is used to efficiently build a string using
// Write methods. It minimizes memory copying.
//
// A Pool is used to cache allocated but unused items for
// later reuse, relieving pressure on the sgarbage collector.
// That is, it makes it easy to build efficient, thread-safe
// free lists.
//
// Go 1.10 or later is required.
package stringpool

import (
	"strings"
	"sync"
)

// StringPool is a sync.Pool for strings.Builder objects.
//
// Reference (Go standard library):
// A Pool is a set of temporary objects that may be individually saved and retrieved.
//
// Any item stored in the Pool may be removed automatically at any time without
// notification. If the Pool holds the only reference when this happens, the
// item might be deallocated.
//
// A Pool is safe for use by multiple goroutines simultaneously.
//
// Pool's purpose is to cache allocated but unused items for later reuse,
// relieving pressure on the garbage collector. That is, it makes it easy to
// build efficient, thread-safe free lists. However, it is not suitable for all
// free lists.
//
// An appropriate use of a Pool is to manage a group of temporary items
// silently shared among and potentially reused by concurrent independent
// clients of a package. Pool provides a way to amortize allocation overhead
// across many clients.
//
// An example of good use of a Pool is in the fmt package, which maintains a
// dynamically-sized store of temporary output buffers. The store scales under
// load (when many goroutines are actively printing) and shrinks when
// quiescent.
//
// On the other hand, a free list maintained as part of a short-lived object is
// not a suitable use for a Pool, since the overhead does not amortize well in
// that scenario. It is more efficient to have such objects implement their own
// free list.
//
// A Pool must not be copied after first use.
type StringPool struct {
	pool sync.Pool
}

// global is the global StringPool used to allocate and
// release strings.Builder objects as needed. A Pool must
// not be copied after first use.
var global *StringPool

func init() {
	global = New()
}

// newBuilder implements the sync.Pool interface
// by providing the New method:
//
//  New Func() interface{}
//
// that specifically returns a strings.Builder
//
// A Builder is used to efficiently build a string using Write methods. It minimizes memory copying. The zero value is ready to use. Do not copy a non-zero Builder.
func newBuilder() interface{} {
	return &strings.Builder{}
}

// New returns a new StringPool instance. A StringPool is
// used to allocate and release strings.Builder objects
// as needed.
//
// A Pool must not be copied after first use. A Pool
// is safe for use by multiple goroutines simultaneously.
func New() *StringPool {
	bp := StringPool{}
	bp.pool.New = newBuilder
	return &bp
}

// Get returns an empty strings.Builder from
// the global pool.
//
// A Builder is used to efficiently build a
// string using Write methods. It minimizes memory
// copying. The zero value is ready to use. Do
// not copy a non-zero Builder.
func Get() *strings.Builder {
	return global.Get()
}

// Release puts the given strings.Builder back into
// the global pool after resetting the Builder.
// It will no longer be accesible after this operation,ss
// but its resources will still be available to be
// reallocated in a new Get() call.
//
// Builders stored in the Pool may be removed
// automatically at any time without notification.
// If the Pool holds the only reference when this
// happens, the item might be deallocated.
func Release(b *strings.Builder) {
	global.Release(b)
}

// Get returns an empty strings.Builder from
// the pool.
//
// A Builder is used to efficiently build a
// string using Write methods. It minimizes memory
// copying. The zero value is ready to use. Do
// not copy a non-zero Builder.
func (bp StringPool) Get() *strings.Builder {
	return bp.pool.Get().(*strings.Builder)
}

// Release puts the given strings.Builder back into
// the pool after resetting the Builder.
// It will no longer be accesible after this operation,ss
// but its resources will still be available to be
// reallocated in a new Get() call.
//
// Builders stored in the Pool may be removed
// automatically at any time without notification.
// If the Pool holds the only reference when this
// happens, the item might be deallocated.
func (bp StringPool) Release(b *strings.Builder) {
	b.Reset()
	bp.pool.Put(b)
}
