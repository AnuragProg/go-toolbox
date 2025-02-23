package concurrent

import "sync"

// RWMutexValue is a generic wrapper that pairs a value of type T with an RWMutex
// to provide thread-safe read and write access.
type RWMutexValue[T any] struct {
	mu sync.RWMutex
	v  T
}

func NewRWMutexValue[T any](v T) *RWMutexValue[T] {
	return &RWMutexValue[T]{
		v: v,
	}
}

// Get returns a copy of the underlying value v.
// It acquires a read lock to ensure safe concurrent access.
func (m *RWMutexValue[T]) Get() T {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.v
}

// Set updates the underlying value v to the new value provided.
// It acquires a write lock to ensure exclusive access during the update.
func (m *RWMutexValue[T]) Set(v T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.v = v
}

// WithRLock executes the provided function f while holding a read lock.
// f is given a copy of the underlying value v, ensuring that any modifications
// made within f do not affect the actual stored value.
// This method is useful for performing compound read operations without incurring
// the cost of multiple Get calls.
func (m* RWMutexValue[T]) WithRLock(f func(T)){
	if f == nil { return }
	m.mu.RLock()
	defer m.mu.RUnlock()
	f(m.v)
}

// WithLock executes the provided function f while holding a write lock.
// f is given a pointer to the underlying value v, allowing safe modifications.
// The lock is held for the duration of f to ensure exclusive access.
func (m* RWMutexValue[T]) WithLock(f func(*T)){
	if f == nil { return }
	m.mu.Lock()
	defer m.mu.Unlock()
	f(&m.v)
}
