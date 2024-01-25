package ring

import (
	"container/ring"
	"sync"
)

// Ring wraps container/ring and adds thread-safety and helper functions.
// Use New to create a new instance with a given size.
type Ring[T any] struct {
	bufferMu sync.RWMutex
	buffer   *ring.Ring
}

// New returns a pointer to a new instance of Ring.
func New[T any](size int) *Ring[T] {
	if size == 0 {
		panic("ring size cannot be zero")
	}

	return &Ring[T]{
		buffer: ring.New(size),
	}
}

// Add an item
func (r *Ring[T]) Add(v T) {
	r.bufferMu.Lock()
	defer r.bufferMu.Unlock()

	r.buffer.Value = v
	r.buffer = r.buffer.Next()
}

func (r *Ring[T]) Slice() []T {
	r.bufferMu.RLock()
	defer r.bufferMu.RUnlock()

	values := []T{}
	r.buffer.Do(func(elem any) {
		if v, ok := elem.(T); ok {
			values = append(values, v)
		}
	})

	return values
}
