package gloop

import (
	"iter"
)

// meanAccumulator helps accumulate both summation and length.
type meanAccumulator[V Number] struct {
	// sum represents the current summation.
	sum V
	// len represents the current length.
	len int
}

// Mean computes the mean value over an iter.Seq sequence.
func Mean[V Number](seq iter.Seq[V]) float64 {
	meanAcc := Reduce(seq, func(acc meanAccumulator[V], value V) meanAccumulator[V] {
		return meanAccumulator[V]{
			sum: acc.sum + value,
			len: acc.len + 1,
		}
	})

	mean := float64(meanAcc.sum) / float64(meanAcc.len)

	return mean
}
