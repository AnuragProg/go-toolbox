package highorder

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }

	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}
	result := Filter(input, isEven)

	assert.Equal(t, expected, result, "Filter() did not return the expected result")
}
