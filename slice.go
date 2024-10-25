package gloop

import (
	"iter"
)

// Slice allows a for loop to range over a given slice.
func Slice[V any](values []V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, value := range values {
			if !yield(value) {
				return
			}
		}
	}
}

// ToSlice converts a sequence to a slice.
func ToSlice[V any](seq iter.Seq[V]) []V {
	l := ToList(seq)
	values := make([]V, l.Len())
	for i := 0; i < len(values); i++ {
		values[i] = l.Remove(l.Front()).(V)
	}

	return values
}
