package gloop

import "iter"

// Values allows looping over an [iter.Seq2] and converting it to an
// [iter.Seq] sequence by discarding the key.
func Values[K, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, value := range seq {
			if !yield(value) {
				return
			}
		}
	}
}
