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
