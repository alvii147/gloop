package gloop

import "iter"

// Keys allows looping over an iter.Seq2, converting it to an iter.Seq
// sequence by discarding the value.
func Keys[K, V any](seq iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for key := range seq {
			if !yield(key) {
				return
			}
		}
	}
}
