package gloop

import "iter"

// Map allows looping over values in a map.
func Map[K comparable, V any](m map[K]V) iter.Seq[V] {
	return Values(Map2(m))
}

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
