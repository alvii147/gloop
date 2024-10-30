package gloop

import (
	"cmp"
	"iter"
)

// Min computes the minimum value over an iter.Seq sequence.
func Min[V cmp.Ordered](seq iter.Seq[V]) V {
	first := true
	return Fold(seq, func(acc V, value V) V {
		if first {
			first = false
			return value
		}

		return min(acc, value)
	})
}
