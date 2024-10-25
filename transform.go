package gloop

import "iter"

// TransformFunc is the function signature of the transformation
// function.
type TransformFunc[V, T any] func(V) T

// Transform runs a given function on each value from a given sequence
// and allows a for loop to range over the returned values.
func Transform[V, T any](seq iter.Seq[V], f TransformFunc[V, T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range seq {
			if !yield(f(value)) {
				return
			}
		}
	}
}
