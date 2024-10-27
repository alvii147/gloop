package gloop

import (
	"cmp"
	"iter"
)

// Max computes the maximum value over an iter.Seq sequence.
func Max[V cmp.Ordered](seq iter.Seq[V]) V {
	first := true
	return Reduce(seq, func(acc V, value V) V {
		if first {
			first = false
			return value
		}

		return max(acc, value)
	})
}

// Max2 computes the maximum value over an iter.Seq2 sequence.
func Max2[K any, V cmp.Ordered](seq iter.Seq2[K, V]) V {
	return Max(Values(seq))
}
