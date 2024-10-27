package gloop

import "iter"

// Map2 allows looping over keys and values in a map.
func Map2[K comparable, V any](m map[K]V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for key, value := range m {
			if !yield(key, value) {
				return
			}
		}
	}
}
