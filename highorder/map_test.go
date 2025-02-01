package highorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	square := func(n int) int { return n * n }

	input := []int{1, 2, 3, 4}
	expected := []int{1, 4, 9, 16}
	result := Map(input, square)

	assert.Equal(t, expected, result, "Map() did not return the expected result")
}
