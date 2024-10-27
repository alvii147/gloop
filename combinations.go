package gloop

import (
	"container/list"
	"iter"
)

// combinations recursively computes and yields the combinations of an
// iter.Seq sequence.
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
// given size for an iter.Seq sequence. The size must be positive.
func Combinations[V any](seq iter.Seq[V], size int) iter.Seq[iter.Seq[V]] {
	if size <= 0 {
		panic("size must be positive")
	}

	return func(yield func(iter.Seq[V]) bool) {
		combinations(size, yield, list.New(), ToList(seq).Front())
	}
}

// combinations2 recursively computes and yields the combinations of an
// iter.Seq2 sequence.
func combinations2[K, V any](
	size int,
	yield func(iter.Seq2[K, V]) bool,
	visitedKeys *list.List,
	visitedValues *list.List,
	unvisitedKeyElem *list.Element,
	unvisitedValueElem *list.Element,
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

	keyElem := unvisitedKeyElem
	valueElem := unvisitedValueElem
	for keyElem != nil && valueElem != nil {
		visitedKeys.PushBack(keyElem.Value)
		visitedValues.PushBack(valueElem.Value)

		if !combinations2(size, yield, visitedKeys, visitedValues, keyElem.Next(), valueElem.Next()) {
			return false
		}

		visitedKeys.Remove(visitedKeys.Back())
		visitedValues.Remove(visitedValues.Back())

		keyElem = keyElem.Next()
		valueElem = valueElem.Next()
	}

	return true
}

// Combinations2 allows a for loop to range over all combinations of a
// given size for an iter.Seq2 sequence. The size must be positive.
func Combinations2[K, V any](seq iter.Seq2[K, V], size int) iter.Seq[iter.Seq2[K, V]] {
	if size <= 0 {
		panic("size must be positive")
	}

	listKeys, listValues := ToList2(seq)
	return func(yield func(iter.Seq2[K, V]) bool) {
		combinations2(size, yield, list.New(), list.New(), listKeys.Front(), listValues.Front())
	}
}
