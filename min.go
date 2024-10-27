package gloop

import (
	"cmp"
	"iter"
)

// Min computes the minumum value over an iter.Seq sequence.
func Min[V cmp.Ordered](seq iter.Seq[V]) V {
	first := true
	return Reduce(seq, func(acc V, value V) V {
		if first {
			first = false
			return value
		}

		return min(acc, value)
	})
}

// Min2 computes the minumum value over an iter.Seq2 sequence.
func Min2[K any, V cmp.Ordered](seq iter.Seq2[K, V]) V {
	return Min(Values(seq))
}
