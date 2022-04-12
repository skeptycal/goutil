// Package builderpool provides a sync.Pool of efficient
// strings.Builder workers that may be reused as needed.
//
// A Builder is used to efficiently build a string using
// Write methods. It minimizes memory copying.
//
// Go 1.10 or later is required.

package gogithub

import (
	"strings"
	"sync"
)

var global *BuilderPool

func init() {
	global = New()
}

// Get returns a strings.Builder from the global pool.
func Get() *strings.Builder {
	return global.Get()
}

// Release puts the given strings.Builder back into the global pool.
func Release(builder *strings.Builder) {
	global.Release(builder)
}

// BuilderPool is wrapper struct of sync.Pool for strings.Builder objects.
type BuilderPool struct {
	pool sync.Pool
}

// New returns a new BuilderPool instance.
func New() *BuilderPool {
	bp := BuilderPool{}
	bp.pool.New = allocBuilder
	return &bp
}

func allocBuilder() interface{} {
	return &strings.Builder{}
}

// Get returns a strings.Builder from the pool.
func (bp *BuilderPool) Get() *strings.Builder {
	return bp.pool.Get().(*strings.Builder)
}

// Release puts the given strings.Builder back into the pool
// after resetting the builder.
func (bp *BuilderPool) Release(builder *strings.Builder) {
	builder.Reset()
	bp.pool.Put(builder)
}
