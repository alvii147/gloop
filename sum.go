package gloop

import (
	"iter"
)

// Sum executes summation over a given sequence.
func Sum[V Summable](seq iter.Seq[V]) V {
	return Reduce(seq, func(acc V, value V) V {
		return acc + value
	})
}
