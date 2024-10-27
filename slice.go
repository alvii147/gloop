package gloop

import (
	"iter"
)

// Slice allows looping over a given slice.
func Slice[V any](values []V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, value := range values {
			if !yield(value) {
				return
			}
		}
	}
}

// ToSlice converts an iter.Seq sequence to a slice.
func ToSlice[V any](seq iter.Seq[V]) []V {
	l := ToList(seq)
	values := make([]V, l.Len())
	for i := 0; i < len(values); i++ {
		values[i] = l.Remove(l.Front()).(V)
	}

	return values
}

// ToSlice2 converts an iter.Seq2 sequence to slices of keys and
// values.
func ToSlice2[K, V any](seq iter.Seq2[K, V]) ([]K, []V) {
	listKeys, listValues := ToList2(seq)
	n := listKeys.Len()

	keys := make([]K, n)
	values := make([]V, n)
	for i := 0; i < n; i++ {
		keys[i] = listKeys.Remove(listKeys.Front()).(K)
		values[i] = listValues.Remove(listValues.Front()).(V)
	}

	return keys, values
}
