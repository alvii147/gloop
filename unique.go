package gloop

import (
	"iter"
)

// Unique allows looping over unique values in an [iter.Seq] sequence.
func Unique[V comparable](seq iter.Seq[V]) iter.Seq[V] {
	visited := map[V]struct{}{}

	return func(yield func(V) bool) {
		for value := range seq {
			_, ok := visited[value]
			if ok {
				continue
			}

			visited[value] = struct{}{}

			if !yield(value) {
				return
			}
		}
	}
}

// Unique2 allows looping over unique key value pairs in an [iter.Seq2]
// sequence.
func Unique2[K, V comparable](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	visited := map[KeyValuePair[K, V]]struct{}{}

	return func(yield func(K, V) bool) {
		for key, value := range seq {
			pair := KeyValuePair[K, V]{
				Key:   key,
				Value: value,
			}

			_, ok := visited[pair]
			if ok {
				continue
			}

			visited[pair] = struct{}{}

			if !yield(key, value) {
				return
			}
		}
	}
}
