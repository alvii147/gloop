package gloop

import (
	"iter"
)

// Window allows looping over an iter.Seq sequence in sliding windows
// of a given size.
func Window[V any](seq iter.Seq[V], size int) iter.Seq[iter.Seq[V]] {
	if size <= 0 {
		panic("size must be positive")
	}

	l := ToList(seq)
	return func(yield func(iter.Seq[V]) bool) {
		firstElem := l.Front()
		for range l.Len() - size + 1 {
			if !yield(func(yield func(V) bool) {
				elem := firstElem
				for range size {
					if !yield(elem.Value.(V)) {
						return
					}

					elem = elem.Next()
				}
			}) {
				return
			}

			firstElem = firstElem.Next()
		}
	}
}

// Windo2w allows looping over an iter.Seq2 sequence in sliding windows
// of a given size.
func Window2[K, V any](seq iter.Seq2[K, V], size int) iter.Seq[iter.Seq2[K, V]] {
	if size <= 0 {
		panic("size must be positive")
	}

	listKeys, listValues := ToList2(seq)
	return func(yield func(iter.Seq2[K, V]) bool) {
		firstKeyElem := listKeys.Front()
		firstValueElem := listValues.Front()
		for range listKeys.Len() - size + 1 {
			if !yield(func(yield func(K, V) bool) {
				keyElem := firstKeyElem
				valueElem := firstValueElem
				for range size {
					if !yield(keyElem.Value.(K), valueElem.Value.(V)) {
						return
					}

					keyElem = keyElem.Next()
					valueElem = valueElem.Next()
				}
			}) {
				return
			}

			firstKeyElem = firstKeyElem.Next()
			firstValueElem = firstValueElem.Next()
		}
	}
}
