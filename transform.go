package gloop

import "iter"

// TransformFunc is the function signature of the transformation
// function in [Transform].
type TransformFunc[V, T any] func(V) T

// Transform runs a given function on each value over an [iter.Seq]
// sequence and allows looping over the returned values.
func Transform[V, T any](seq iter.Seq[V], f TransformFunc[V, T]) iter.Seq[T] {
	return Transform2(Enumerate(seq), func(_ int, value V) T {
		return f(value)
	})
}

// TransformFunc is the function signature of the transformation
// function in [Transform2].
type Transform2Func[K, V, T any] func(K, V) T

// Transform2 runs a given function on each key and value over an
// [iter.Seq2] sequence and allows looping over the returned values.
func Transform2[K, V, T any](seq iter.Seq2[K, V], f Transform2Func[K, V, T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for key, value := range seq {
			if !yield(f(key, value)) {
				return
			}
		}
	}
}
