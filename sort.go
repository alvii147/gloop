package gloop

import (
	"cmp"
	"iter"
	"sort"
)

// Sort allows looping over an [iter.Seq] sequence in sorted order.
func Sort[V cmp.Ordered](seq iter.Seq[V], ascending bool) iter.Seq[V] {
	return SortByComparison(seq, func(value1 V, value2 V) bool {
		return value1 < value2
	}, ascending)
}

// SortByComparisonFunc is the function signature of the comparison
// function used in [SortByComparison].
type SortByComparisonFunc[V any] func(V, V) bool

// SortByComparison allows looping over an [iter.Seq] sequence in
// sorted order using a comparison function.
func SortByComparison[V any](
	seq iter.Seq[V],
	less SortByComparisonFunc[V],
	ascending bool,
) iter.Seq[V] {
	values := ToSlice(seq)
	sort.Slice(values, func(i int, j int) bool {
		isLess := less(values[i], values[j])
		return (ascending && isLess) || (!ascending && !isLess)
	})

	return func(yield func(V) bool) {
		for _, value := range values {
			if !yield(value) {
				return
			}
		}
	}
}

// SortByComparison2Func is the function signature of the comparison
// function used in [SortByComparison2].
type SortByComparison2Func[K, V any] func(K, V, K, V) bool

// SortByComparison2 allows looping over an [iter.Seq2] sequence in
// sorted order using a comparison function.
func SortByComparison2[K, V any](
	seq iter.Seq2[K, V],
	less SortByComparison2Func[K, V],
	ascending bool,
) iter.Seq2[K, V] {
	return KeyValue(SortByComparison(
		KeyValue2(seq),
		func(pair1 KeyValuePair[K, V], pair2 KeyValuePair[K, V]) bool {
			return less(pair1.Key, pair1.Value, pair2.Key, pair2.Value)
		},
		ascending,
	))
}

// SortByRankFunc is the function signature of the ranking function
// used in [SortByRank].
type SortByRankFunc[V any, R cmp.Ordered] func(V) R

// SortByRank allows looping over an [iter.Seq] sequence in sorted
// order using a ranking function.
func SortByRank[V any, R cmp.Ordered](
	seq iter.Seq[V],
	rank SortByRankFunc[V, R],
	ascending bool,
) iter.Seq[V] {
	return SortByComparison(seq, func(value1 V, value2 V) bool {
		return rank(value1) < rank(value2)
	}, ascending)
}

// SortByRank2Func is the function signature of the ranking function
// used in [SortByRank2].
type SortByRank2Func[K, V any, R cmp.Ordered] func(K, V) R

// SortByRank2 allows looping over an [iter.Seq2] sequence in sorted
// order using a ranking function.
func SortByRank2[K, V any, R cmp.Ordered](
	seq iter.Seq2[K, V],
	rank SortByRank2Func[K, V, R],
	ascending bool,
) iter.Seq2[K, V] {
	return SortByComparison2(seq, func(key1 K, value1 V, key2 K, value2 V) bool {
		return rank(key1, value1) < rank(key2, value2)
	}, ascending)
}
