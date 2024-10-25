package gloop

import (
	"container/list"
	"iter"
)

// combinations recursively computes and yields the combinations of a
// given sequence.
func combinations[V any](
	size int,
	yield func(iter.Seq[V]) bool,
	visited *list.List,
	unvisitedElem *list.Element,
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

	elem := unvisitedElem
	for elem != nil {
		visited.PushBack(elem.Value)
		if !combinations(size, yield, visited, elem.Next()) {
			return false
		}

		visited.Remove(visited.Back())
		elem = elem.Next()
	}

	return true
}

// Combinations allows a for loop to range over all combinations of a
// given size for a given sequence. The size must be positive.
func Combinations[V any](seq iter.Seq[V], size int) iter.Seq[iter.Seq[V]] {
	if size <= 0 {
		panic("size must be positive")
	}

	return func(yield func(iter.Seq[V]) bool) {
		combinations(size, yield, list.New(), ToList(seq).Front())
	}
}
