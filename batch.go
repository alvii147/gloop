package gloop

import (
	"iter"
)

// Batch allows looping over an iter.Seq sequence in batches of a given
// size. The batch size must be positive.
func Batch[V any](seq iter.Seq[V], size int) iter.Seq[iter.Seq[V]] {
	if size <= 0 {
		panic("size must be positive")
	}

	l := ToList(seq)
	return func(yield func(iter.Seq[V]) bool) {
		elem := l.Front()
		for {
			if elem == nil {
				return
			}

			if !yield(func(yield func(V) bool) {
				for range size {
					if elem == nil {
						return
					}

					if !yield(elem.Value.(V)) {
						return
					}

					elem = elem.Next()
				}
			}) {
				return
			}
		}
	}
}
