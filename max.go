package gloop

import (
	"cmp"
	"iter"
)

// Max computes the maximum value over an [iter.Seq] sequence.
func Max[V cmp.Ordered](seq iter.Seq[V]) V {
	return MaxByComparison(seq, func(acc V, value V) bool {
		return acc < value
	})
}

// MaxByComparisonFunc is the function signature of the comparison
// function used in MaxByComparison].
type MaxByComparisonFunc[V any] func(V, V) bool

// MaxByComparison computes the maximum value over an [iter.Seq]
// sequence using a comparison function.
func MaxByComparison[V any](
	seq iter.Seq[V],
	less SortByComparisonFunc[V],
) V {
	return Reduce(seq, func(acc V, value V) V {
		if less(acc, value) {
			return value
		}

		return acc
	})
}

// MaxByComparison2Func is the function signature of the comparison
// function used in [MaxByComparison2].
type MaxByComparison2Func[K, V any] func(K, V, K, V) bool

// MaxByComparison2 computes the maximum key and value over an
// [iter.Seq2] sequence using a comparison function.
func MaxByComparison2[K, V any](
	seq iter.Seq2[K, V],
	less MinByComparison2Func[K, V],
) (K, V) {
	return Reduce2(seq, func(accKey K, accValue V, key K, value V) (K, V) {
		if less(accKey, accValue, key, value) {
			return key, value
		}

		return accKey, accValue
	})
}

// MaxByRankFunc is the function signature of the ranking function used
// in [MaxByRank].
type MaxByRankFunc[V any, R cmp.Ordered] func(V) R

// MaxByRank computes the maximum value over an [iter.Seq] sequence
// using a ranking function.
func MaxByRank[V any, R cmp.Ordered](
	seq iter.Seq[V],
	rank MinByRankFunc[V, R],
) V {
	return MaxByComparison(seq, func(acc V, value V) bool {
		return rank(acc) < rank(value)
	})
}

// MaxByRank2Func is the function signature of the ranking function
// used in [MaxByRank2].
type MaxByRank2Func[K, V any, R cmp.Ordered] func(K, V) R

// MaxByRank2 computes the maximum value over an [iter.Seq2] sequence
// using a ranking function.
func MaxByRank2[K, V any, R cmp.Ordered](
	seq iter.Seq2[K, V],
	rank MinByRank2Func[K, V, R],
) (K, V) {
	return MaxByComparison2(seq, func(accKey K, accValue V, key K, value V) bool {
		return rank(accKey, accValue) < rank(key, value)
	})
}
