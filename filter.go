package gloop

import "iter"

// FilterFunc is the function signature of the filtering function used
// in [Filter].
type FilterFunc[V any] func(V) bool

// Filter runs a given function on each value from an [iter.Seq]
// sequence and allows looping over values for which the function
// returns true.
func Filter[V any](seq iter.Seq[V], f FilterFunc[V]) iter.Seq[V] {
	return Values(Filter2(Enumerate(seq), func(_ int, value V) bool {
		return f(value)
	}))
}

// Filter2Func is the function signature of the filtering function used
// in [Filter2].
type Filter2Func[K, V any] func(K, V) bool

// Filter2 runs a given function on each value from an [iter.Seq2]
// sequence and allows looping over values for which the function
// returns true.
func Filter2[K, V any](seq iter.Seq2[K, V], f Filter2Func[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for i, value := range seq {
			if !f(i, value) {
				continue
			}

			if !yield(i, value) {
				return
			}
		}
	}
}
