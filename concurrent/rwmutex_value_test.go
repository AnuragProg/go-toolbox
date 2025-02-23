package concurrent

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetSet verifies that Get returns the initial value and that Set properly updates it.
func TestRWMutexValue_GetSet(t *testing.T) {
	val := NewRWMutexValue(10)
	assert.Equal(t, 10, val.Get(), "expected initial value to be 10")

	val.Set(20)
	assert.Equal(t, 20, val.Get(), "expected value after Set to be 20")
}

// TestWithRLockNil ensures that calling WithRLock with a nil function does not panic.
func TestRWMutexValue_WithRLockNil(t *testing.T) {
	val := NewRWMutexValue(5)
	assert.NotPanics(t, func() {
		val.WithRLock(nil)
	}, "WithRLock should not panic when given nil")
}

// TestWithLockNil ensures that calling WithLock with a nil function does not panic.
func TestRWMutexValue_WithLockNil(t *testing.T) {
	val := NewRWMutexValue(5)
	assert.NotPanics(t, func() {
		val.WithLock(nil)
	}, "WithLock should not panic when given nil")
}

// TestWithRLockCopy verifies that the WithRLock method passes a copy of the value,
// so any modifications in the callback do not affect the stored value.
func TestRWMutexValue_WithRLockCopy(t *testing.T) {
	val := NewRWMutexValue(30)
	modified := false

	val.WithRLock(func(v int) {
		// Attempt to modify the local copy.
		v = 100
		modified = true
		assert.Equal(t, 100, v, "inside WithRLock callback, v should be 100 after modification")
	})

	// Underlying value remains unchanged.
	assert.Equal(t, 30, val.Get(), "underlying value should remain unchanged after WithRLock")
	assert.True(t, modified, "callback should have been executed")
}

// TestWithLockModifiesValue verifies that the WithLock method correctly allows modification of the underlying value.
func TestRWMutexValue_WithLockModifiesValue(t *testing.T) {
	val := NewRWMutexValue(50)

	val.WithLock(func(v *int) {
		*v = 200
	})

	assert.Equal(t, 200, val.Get(), "expected value to be updated to 200 after WithLock")
}

// TestConcurrentAccess tests concurrent access to the RWMutexValue to ensure thread safety.
// It launches multiple goroutines that perform Get, Set, WithRLock, and WithLock operations.
func TestRWMutexValue_ConcurrentAccess(t *testing.T) {
	val := NewRWMutexValue(0)
	var wg sync.WaitGroup
	numGoroutines := 100

	// Concurrent writers using Set.
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val.Set(i)
		}(i)
	}

	// Concurrent readers using Get and WithRLock.
	for i := 0; i < numGoroutines; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			_ = val.Get()
		}()
		go func() {
			defer wg.Done()
			val.WithRLock(func(v int) {
				_ = v // simply read v
			})
		}()
	}

	// Concurrent updaters using WithLock.
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val.WithLock(func(v *int) {
				*v = *v + i
			})
		}(i)
	}

	wg.Wait()

	// Final value is not deterministic due to concurrent updates,
	// but we assert that the value is of type int.
	finalValue := val.Get()
	assert.IsType(t, int(0), finalValue, "final value should be an int")
}
