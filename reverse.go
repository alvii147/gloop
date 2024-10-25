package gloop

import (
	"iter"
)

// Reverse allows looping over an iter.Seq sequence in order of
// descending index.
func Reverse[V any](seq iter.Seq[V]) iter.Seq[V] {
	l := ToList(seq)
	return func(yield func(V) bool) {
		for l.Len() > 0 {
			value := l.Remove(l.Back()).(V)
			if !yield(value) {
				return
			}
		}
	}
}
