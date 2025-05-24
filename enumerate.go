package gloop

import "iter"

// Enumerate allows looping over an [iter.Seq] sequence with an index,
// converting it to an [iter.Seq2] sequence.
func Enumerate[V any](seq iter.Seq[V]) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		i := 0
		for value := range seq {
			if !yield(i, value) {
				return
			}

			i++
		}
	}
}
