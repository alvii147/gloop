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
