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

// permutations2 recursively computes and yields the permutations of an
// iter.Seq2 sequence.
func permutations2[K, V any](
	size int,
	yield func(iter.Seq2[K, V]) bool,
	listKeys *list.List,
	listValues *list.List,
	visitedKeys *list.List,
	visitedValues *list.List,
	visitedIdx []bool,
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

	i := 0
	keyElem := listKeys.Front()
	valueElem := listValues.Front()
	for keyElem != nil && valueElem != nil {
		if visitedIdx[i] {
			i++
			keyElem = keyElem.Next()
			valueElem = valueElem.Next()
			continue
		}

		visitedKeys.PushBack(keyElem.Value)
		visitedValues.PushBack(valueElem.Value)

		visitedIdx[i] = true

		if !permutations2(size, yield, listKeys, listValues, visitedKeys, visitedValues, visitedIdx) {
			return false
		}

		visitedKeys.Remove(visitedKeys.Back())
		visitedValues.Remove(visitedValues.Back())

		visitedIdx[i] = false

		i++
		keyElem = keyElem.Next()
		valueElem = valueElem.Next()
	}

	return true
}

// Permutations2 allows looping over all permutations of a given size
// for an iter.Seq2 sequence. The size must be positive.
func Permutations2[K, V any](seq iter.Seq2[K, V], size int) iter.Seq[iter.Seq2[K, V]] {
	if size <= 0 {
		panic("size must be positive")
	}

	listKeys, listValues := ToList2(seq)
	return func(yield func(iter.Seq2[K, V]) bool) {
		permutations2(size, yield, listKeys, listValues, list.New(), list.New(), make([]bool, listKeys.Len()))
	}
}
