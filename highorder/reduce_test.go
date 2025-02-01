package highorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	sum := func(n int, acc int) int { return acc + n }

	input := []int{1, 2, 3, 4}
	expected := 10
	result := Reduce(input, 0, sum)

	assert.Equal(t, expected, result, "Reduce() did not return the expected result")
}
