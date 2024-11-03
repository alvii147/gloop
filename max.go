package gloop

import (
	"cmp"
	"iter"
)

// Max computes the maximum value over an [iter.Seq] sequence.
func Max[V cmp.Ordered](seq iter.Seq[V]) V {
	return Reduce(seq, func(acc V, value V) V {
		return max(acc, value)
	})
}
