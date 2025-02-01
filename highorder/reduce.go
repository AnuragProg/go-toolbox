package highorder

// Reduce applies an accumulator function to each element of the input slice,
// carrying forward the accumulated result.
// It returns the final accumulated value.
func Reduce[T, U any](input []T, acc U, accfn func(T, U) U) U {
	for _, v := range input {
		acc = accfn(v, acc)
	}
	return acc
}
