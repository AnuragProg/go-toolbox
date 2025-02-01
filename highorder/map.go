package highorder

// Map applies a transformation function to each element in the input slice.
// It returns a new slice containing the transformed elements.
func Map[T, U any](input []T, fn func(T)U) []U {
	result := make([]U, 0, len(input))
	for _, v := range input {
		result = append(result, fn(v))
	}
	return result
}
