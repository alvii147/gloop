package gloop

import "iter"

// TransformFunc is the function signature of the transformation
// function.
type TransformFunc[V, T any] func(V) T

// Transform runs a given function on each value over an iter.Seq
// sequence and allows looping over the returned values.
func Transform[V, T any](seq iter.Seq[V], f TransformFunc[V, T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range seq {
			if !yield(f(value)) {
				return
			}
		}
	}
}
