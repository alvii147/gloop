package gloop

import "iter"

// Product executes the product of values over an iter.Seq sequence.
func Product[V Productable](seq iter.Seq[V]) V {
	return Reduce(seq, func(acc V, value V) V {
		return acc * value
	}, WithReduceInitialValue[V](1))
}

// Product2 executes the product of values over an iter.Seq2 sequence.
func Product2[K any, V Productable](seq iter.Seq2[K, V]) V {
	return Product(Values(seq))
}
