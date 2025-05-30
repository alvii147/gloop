package gloop

import "iter"

// Chain allows looping over multiple [iter.Seq] sequences.
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

// Chain2 allows looping over multiple [iter.Seq2] sequences.
func Chain2[K, V any](seqs ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, seq := range seqs {
			for i, value := range seq {
				if !yield(i, value) {
					return
				}
			}
		}
	}
}
