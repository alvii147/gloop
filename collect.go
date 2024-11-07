package gloop

import "iter"

// Collect allows looping over a given set of values.
func Collect[V any](values ...V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, value := range values {
			if !yield(value) {
				return
			}
		}
	}
}
