package gloop

import (
	"container/list"
	"iter"
)

// DeferLoopFunc is the function signature of the defer function used
// in [DeferLoop].
type DeferLoopFunc func(func())

// DeferLoop allows looping over an [iter.Seq] sequence, yielding a
// defer function that can register another function to be executed at
// the end of the currently running loop. If multiple functions are
// registered, they are executed in FIFO order.
func DeferLoop[V any](seq iter.Seq[V]) iter.Seq2[V, DeferLoopFunc] {
	return func(yield func(V, DeferLoopFunc) bool) {
		for value := range seq {
			funcs := list.New()
			setDeferFunc := func(deferFunc func()) {
				funcs.PushFront(deferFunc)
			}

			end := yield(value, setDeferFunc)

			for elem := range List(funcs) {
				f := elem.Value.(func())
				if f != nil {
					f()
				}
			}

			if !end {
				return
			}
		}
	}
}
