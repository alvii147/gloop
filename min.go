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

// MinByRankFunc is the function signature of the ranking function used
// in [MinByRank].
type MinByRankFunc[V any, R cmp.Ordered] func(V) R

// MinByRank computes the minimum value over an [iter.Seq] sequence
// using a ranking function.
func MinByRank[V any, R cmp.Ordered](
	seq iter.Seq[V],
	rank MinByRankFunc[V, R],
) V {
	return MinByComparison(seq, func(acc V, value V) bool {
		return rank(acc) < rank(value)
	})
}

// MinByRank2Func is the function signature of the ranking function
// used in [MinByRank2].
type MinByRank2Func[K, V any, R cmp.Ordered] func(K, V) R

// MinByRank2 computes the minimum value over an [iter.Seq2] sequence
// using a ranking function.
func MinByRank2[K, V any, R cmp.Ordered](
	seq iter.Seq2[K, V],
	rank MinByRank2Func[K, V, R],
) (K, V) {
	return MinByComparison2(seq, func(accKey K, accValue V, key K, value V) bool {
		return rank(accKey, accValue) < rank(key, value)
	})
}
