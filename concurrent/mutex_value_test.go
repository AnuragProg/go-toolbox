package concurrent

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMutexValue_GetSet(t *testing.T) {
	m := NewMutexValue(42)
	assert.Equal(t, 42, m.Get(), "initial value should be 42")

	m.Set(100)
	assert.Equal(t, 100, m.Get(), "value should be updated to 100")
}

func TestMutexValue_WithLock(t *testing.T) {
	m := NewMutexValue(10)
	m.WithLock(func(val *int) {
		*val += 5
	})
	assert.Equal(t, 15, m.Get(), "value should be modified to 15")
}

func TestMutexValue_ConcurrentAccess(t *testing.T) {
	m := NewMutexValue(0)
	var wg sync.WaitGroup
	iterations := 100

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			m.WithLock(func(val *int) {
				*val++
			})
			wg.Done()
		}()
	}

	wg.Wait()
	assert.Equal(t, iterations, m.Get(), "value should equal number of increments")
}
