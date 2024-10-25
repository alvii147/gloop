package gloop

import (
	"iter"
)

// Sum executes summation over an iter.Seq sequence.
func Sum[V Summable](seq iter.Seq[V]) V {
	return Reduce(seq, func(acc V, value V) V {
		return acc + value
	})
}
