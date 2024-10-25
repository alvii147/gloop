package gloop

import "iter"

// Chain allows looping over multiple iter.Seq sequences.
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
