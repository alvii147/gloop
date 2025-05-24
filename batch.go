package gloop

import (
	"iter"
)

// Batch allows looping over an [iter.Seq] sequence in batches of a
// given size. The batch size must be positive.
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

// Batch2 allows looping over an [iter.Seq2] sequence in batches of a
// given size. The batch size must be positive.
func Batch2[K, V any](seq iter.Seq2[K, V], size int) iter.Seq[iter.Seq2[K, V]] {
	if size <= 0 {
		panic("size must be positive")
	}

	listKeys, listValues := ToList2(seq)

	return func(yield func(iter.Seq2[K, V]) bool) {
		keyElem := listKeys.Front()
		valueElem := listValues.Front()

		for {
			if keyElem == nil || valueElem == nil {
				return
			}

			if !yield(func(yield func(K, V) bool) {
				for range size {
					if keyElem == nil || valueElem == nil {
						return
					}

					key := keyElem.Value.(K)
					value := valueElem.Value.(V)
					if !yield(key, value) {
						return
					}

					keyElem = keyElem.Next()
					valueElem = valueElem.Next()
				}
			}) {
				return
			}
		}
	}
}
