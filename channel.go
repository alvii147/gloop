package gloop

import "iter"

// Channel allows a for loop to receive and range over values from a
// given channel.
func Channel[V any](ch <-chan V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for value := range ch {
			if !yield(value) {
				return
			}
		}
	}
}
