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
