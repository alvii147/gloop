package gloop

import (
	"cmp"
	"iter"
)

// Min computes the minimum value over an [iter.Seq] sequence.
func Min[V cmp.Ordered](seq iter.Seq[V]) V {
	return Reduce(seq, func(acc V, value V) V {
		return min(acc, value)
	})
}
