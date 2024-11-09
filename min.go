package gloop

import (
	"cmp"
	"iter"
)

// Min computes the minimum value over an [iter.Seq] sequence.
func Min[V cmp.Ordered](seq iter.Seq[V]) V {
	return MinByComparison(seq, func(acc V, value V) bool {
		return acc < value
	})
}

// MinByComparisonFunc is the function signature of the comparison
// function used in [MinByComparison].
type MinByComparisonFunc[V any] func(V, V) bool

// MinByComparison computes the minimum value over an [iter.Seq]
// sequence using a comparison function.
func MinByComparison[V any](
	seq iter.Seq[V],
	less SortByComparisonFunc[V],
) V {
	return Reduce(seq, func(acc V, value V) V {
		if less(acc, value) {
			return acc
		}

		return value
	})
}

// MinByComparison2Func is the function signature of the comparison
// function used in [MinByComparison2].
type MinByComparison2Func[K, V any] func(K, V, K, V) bool

// MinByComparison2 computes the minimum key and value over an
// [iter.Seq2] sequence using a comparison function.
func MinByComparison2[K, V any](
	seq iter.Seq2[K, V],
	less MinByComparison2Func[K, V],
) (K, V) {
	return Reduce2(seq, func(accKey K, accValue V, key K, value V) (K, V) {
		if less(accKey, accValue, key, value) {
			return accKey, accValue
		}

		return key, value
	})
}
