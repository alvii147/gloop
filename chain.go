package gloop

import "iter"

// Chain allows a for loop to range over multiple sequences.
func Chain[V any](seqs ...iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, seq := range seqs {
			for value := range seq {
				if !yield(value) {
					return
				}
			}
		}
	}
}
