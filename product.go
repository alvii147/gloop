package gloop

import "iter"

// Product executes the product of values a given sequence.
func Product[V Productable](seq iter.Seq[V]) V {
	return Reduce(seq, func(acc V, value V) V {
		return acc * value
	}, WithReduceInitialValue[V](1))
}
