package gloop

import "iter"

// FilterFunc is the function signature of the filtering function.
type FilterFunc[V any] func(V) bool

// Filter runs a given function on each value from an iter.Seq sequence
// and allows looping over values for which the function returns true.
func Filter[V any](seq iter.Seq[V], f FilterFunc[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		i := 0
		for value := range seq {
			if !f(value) {
				continue
			}

			if !yield(value) {
				return
			}

			i++
		}
	}
}
