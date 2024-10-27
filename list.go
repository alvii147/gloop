package gloop

import (
	"container/list"
	"iter"
)

// List allows looping over a given linked list from container/list.
func List(l *list.List) iter.Seq[*list.Element] {
	return func(yield func(*list.Element) bool) {
		elem := l.Front()
		for elem != nil {
			if !yield(elem) {
				return
			}

			elem = elem.Next()
		}
	}
}

// ToList converts an iter.Seq sequence to a linked list from
// container/list.
func ToList[V any](seq iter.Seq[V]) *list.List {
	l := list.New()
	for value := range seq {
		l.PushBack(value)
	}

	return l
}

// ToList2 converts an iter.Seq2 sequence to linked lists
// from container/list of keys and values.
func ToList2[K, V any](seq iter.Seq2[K, V]) (*list.List, *list.List) {
	listKeys := list.New()
	listValues := list.New()
	for key, value := range seq {
		listKeys.PushBack(key)
		listValues.PushBack(value)
	}

	return listKeys, listValues
}
