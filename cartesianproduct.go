package gloop

import (
	"container/list"
	"iter"
)

// cartesianProduct recursively computes and yields the Cartesian
// product for an iter.Seq sequence.
func cartesianProduct[V any](
	seq iter.Seq[V],
	size int,
	yield func(iter.Seq[V]) bool,
	visited *list.List,
) bool {
	if visited.Len() == size {
		return yield(func(yield func(V) bool) {
			for elem := range List(visited) {
				if !yield(elem.Value.(V)) {
					return
				}
			}
		})
	}

	for value := range seq {
		visited.PushBack(value)
		if !cartesianProduct(seq, size, yield, visited) {
			return false
		}

		visited.Remove(visited.Back())
	}

	return true
}

// CartesianProduct allows looping over the Cartesian product of a
// given size for an iter.Seq sequence. The size must be positive.
func CartesianProduct[V any](seq iter.Seq[V], size int) iter.Seq[iter.Seq[V]] {
	if size <= 0 {
		panic("size must be positive")
	}

	return func(yield func(iter.Seq[V]) bool) {
		cartesianProduct(seq, size, yield, list.New())
	}
}

// cartesianProduct2 recursively computes and yields the Cartesian
// product for an iter.Seq2 sequence.
func cartesianProduct2[K, V any](
	seq iter.Seq2[K, V],
	size int,
	yield func(iter.Seq2[K, V]) bool,
	visitedKeys *list.List,
	visitedValues *list.List,
) bool {
	if visitedKeys.Len() == size || visitedValues.Len() == size {
		return yield(func(yield func(K, V) bool) {
			for keyElem, valueElem := range Zip(List(visitedKeys), List(visitedValues)) {
				key := keyElem.Value.(K)
				value := valueElem.Value.(V)
				if !yield(key, value) {
					return
				}
			}
		})
	}

	for key, value := range seq {
		visitedKeys.PushBack(key)
		visitedValues.PushBack(value)
		if !cartesianProduct2(seq, size, yield, visitedKeys, visitedValues) {
			return false
		}

		visitedKeys.Remove(visitedKeys.Back())
		visitedValues.Remove(visitedValues.Back())
	}

	return true
}

// CartesianProduct2 allows looping over the Cartesian product of a
// given size for an iter.Seq2 sequence. The size must be positive.
func CartesianProduct2[K, V any](seq iter.Seq2[K, V], size int) iter.Seq[iter.Seq2[K, V]] {
	if size <= 0 {
		panic("size must be positive")
	}

	return func(yield func(iter.Seq2[K, V]) bool) {
		cartesianProduct2(seq, size, yield, list.New(), list.New())
	}
}
