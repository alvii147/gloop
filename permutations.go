package gloop

import (
	"container/list"
	"iter"
)

// permutations recursively computes and yields the permutations of an
// iter.Seq sequence.
func permutations[V any](
	size int,
	yield func(iter.Seq[V]) bool,
	l *list.List,
	visited *list.List,
	visitedIdx []bool,
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

	for i, elem := range Enumerate(List(l)) {
		if visitedIdx[i] {
			continue
		}

		visited.PushBack(elem.Value)
		visitedIdx[i] = true

		if !permutations(size, yield, l, visited, visitedIdx) {
			return false
		}

		visited.Remove(visited.Back())
		visitedIdx[i] = false
	}

	return true
}

// Permutations allows looping over all permutations of a given size
// for an iter.Seq sequence. The size must be positive.
func Permutations[V any](seq iter.Seq[V], size int) iter.Seq[iter.Seq[V]] {
	if size <= 0 {
		panic("size must be positive")
	}

	return func(yield func(iter.Seq[V]) bool) {
		l := ToList(seq)
		permutations(size, yield, l, list.New(), make([]bool, l.Len()))
	}
}
