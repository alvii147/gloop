package gloop

import "iter"

// Product computes the product of values over an iter.Seq sequence.
func Product[V Productable](seq iter.Seq[V]) V {
	return Reduce(seq, func(acc V, value V) V {
		return acc * value
	})
}
