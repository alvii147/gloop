package gloop

import "iter"

// Enumerate allows a for loop to iterate over a sequence with an
// index.
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
