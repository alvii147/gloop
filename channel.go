package gloop

import "iter"

// Channel allows looping over values from a given channel. The values
// are consumed from the channel.
func Channel[V any](ch <-chan V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for value := range ch {
			if !yield(value) {
				return
			}
		}
	}
}

// Channel2 allows looping over values from a given channel with an
// index. The values are consumed from the channel.
func Channel2[V any](ch <-chan V) iter.Seq2[int, V] {
	return Enumerate(Channel(ch))
}
