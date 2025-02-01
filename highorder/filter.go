package highorder


// Filter filters elements from the input slice based on the given predicate function.
// It returns a new slice containing only elements that satisfy the predicate.
func Filter[T any](input []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range input {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
