package concurrent

import "sync"

// MutexValue holds a value of any type T with a mutex for safe concurrent access.
// All operations on the value are protected by the mutex.
type MutexValue[T any] struct {
	mu sync.Mutex
	v  T
}

func NewMutexValue[T any](v T) *MutexValue[T] {
	return &MutexValue[T]{
		v: v,
	}
}

// Get returns a copy of the current value safely.
// It locks the mutex during the operation.
func (m *MutexValue[T]) Get() T {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.v
}

// Set updates the value safely by locking the mutex.
// It replaces the current value with the new one.
func (m *MutexValue[T]) Set(v T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.v = v
}

// WithLock executes the given function while holding the lock.
// The function receives a pointer to the value for in-place modification.
func (m *MutexValue[T]) WithLock(f func(*T)){
	if f == nil { return }
	m.mu.Lock()
	defer m.mu.Unlock()
	f(&m.v)
}
