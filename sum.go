package gloop

import (
	"iter"
)

// Sum computes summation over an [iter.Seq] sequence.
func Sum[V Summable](seq iter.Seq[V]) V {
	return Fold(seq, func(acc V, value V) V {
		return acc + value
	})
}
