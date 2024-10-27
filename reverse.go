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

// Reverse2 allows looping over an iter.Seq2 sequence in order of
// descending index.
func Reverse2[K, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	listKeys, listValues := ToList2(seq)
	return func(yield func(K, V) bool) {
		for listKeys.Len() > 0 && listValues.Len() > 0 {
			key := listKeys.Remove(listKeys.Back()).(K)
			value := listValues.Remove(listValues.Back()).(V)
			if !yield(key, value) {
				return
			}
		}
	}
}
